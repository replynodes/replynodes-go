#!/usr/bin/env python3
"""Verify that the public Go surface matches the canonical GET operations."""

from __future__ import annotations

import json
import re
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parent.parent
SPEC = ROOT / "api-docs" / "replynodes-fetcher.openapi.json"
RESOURCES = ROOT / "resources.go"


def fail(message: str) -> int:
    print(f"surface coverage check failed: {message}", file=sys.stderr)
    return 1


def parameter_metadata(parameters: list[dict]) -> list[tuple[str, str, bool]]:
    metadata = []
    for parameter in parameters:
        if "$ref" in parameter:
            raise ValueError(f"unresolved parameter reference: {parameter['$ref']}")
        location = parameter.get("in")
        name = parameter.get("name")
        if not isinstance(location, str) or not isinstance(name, str):
            raise ValueError(f"parameter is missing location or name: {parameter!r}")
        metadata.append((location, name, bool(parameter.get("required", False))))
    return metadata


def canonical_operations(spec: dict) -> dict[str, tuple[str, list[tuple[str, str, bool]]]]:
    operations = {}
    for path, path_item in spec.get("paths", {}).items():
        path_parameters = path_item.get("parameters", [])
        for method, operation in path_item.items():
            if method.lower() != "get" or not operation.get("operationId"):
                continue
            operation_id = operation["operationId"]
            if operation_id in operations:
                raise ValueError(f"canonical GET operation ID is not unique: {operation_id}")
            parameters = path_parameters + operation.get("parameters", [])
            operations[operation_id] = (path, parameter_metadata(parameters))
    return operations


def parse_registry(source: str) -> list[tuple[str, str, str]]:
    registry_start = source.find("var PublicOperationRegistry")
    if registry_start < 0:
        raise ValueError("PublicOperationRegistry declaration not found")

    entries = []
    current_resource: str | None = None
    for line in source[registry_start:].splitlines():
        group = re.match(r'^\s*"([^"\\]+)":\s*\{\s*$', line)
        if group:
            current_resource = group.group(1)
            continue
        entry = re.match(r'^\s+"([^"\\]+)":\s*"([^"\\]+)",\s*$', line)
        if entry and current_resource is not None:
            entries.append((current_resource, entry.group(1), entry.group(2)))
    return entries


def parse_parameter_structs(source: str) -> dict[str, list[tuple[str, str, bool]]]:
    structs = {}
    struct_pattern = re.compile(
        r"^type (?P<name>[A-Za-z0-9]+Params) struct \{(?P<body>.*?)^\}",
        re.MULTILINE | re.DOTALL,
    )
    field_pattern = re.compile(r'^\s*[A-Za-z0-9]+\s+[^\s]+\s+`(?P<tag>[^`]+)`\s*$')
    for match in struct_pattern.finditer(source):
        metadata = []
        for line in match.group("body").splitlines():
            field = field_pattern.match(line)
            if not field:
                continue
            tags = field.group("tag")
            for location in ("path", "query"):
                tag = re.search(rf'{location}:"([^"]+)"', tags)
                if not tag or tag.group(1) == "-":
                    continue
                parts = tag.group(1).split(",")
                metadata.append((location, parts[0], "required" in parts[1:]))
                break
        structs[match.group("name")] = metadata
    return structs


def parse_methods(source: str) -> list[dict[str, str]]:
    method_pattern = re.compile(
        r"^func \(s \*(?P<service>[A-Za-z0-9]+)Service\) "
        r"(?P<method>[A-Za-z0-9]+)\(ctx context\.Context, params "
        r"(?P<params>[A-Za-z0-9]+Params)\) \(\*Response, error\) \{\s*"
        r"return s\.client\.do\(ctx, \"(?P<operation>[^\"]+)\", "
        r"\"(?P<path>[^\"]+)\", params\)",
        re.MULTILINE,
    )
    return [match.groupdict() for match in method_pattern.finditer(source)]


def main() -> int:
    try:
        spec = json.loads(SPEC.read_text(encoding="utf-8"))
        source = RESOURCES.read_text(encoding="utf-8")
        canonical = canonical_operations(spec)
        entries = parse_registry(source)
        parameter_structs = parse_parameter_structs(source)
        methods = parse_methods(source)
    except (OSError, json.JSONDecodeError, ValueError) as exc:
        return fail(str(exc))

    registry_ids = [operation_id for _, _, operation_id in entries]
    missing = sorted(set(canonical) - set(registry_ids))
    unexpected = sorted(set(registry_ids) - set(canonical))
    if missing or unexpected:
        details = []
        if missing:
            details.append("missing: " + ", ".join(missing))
        if unexpected:
            details.append("unexpected: " + ", ".join(unexpected))
        return fail("canonical and public operation IDs differ (" + "; ".join(details) + ")")

    duplicate_registry_ids = sorted({item for item in registry_ids if registry_ids.count(item) > 1})
    if duplicate_registry_ids != ["googleSearch"] or registry_ids.count("googleSearch") != 2:
        return fail("only the intentional googleSearch alias may be duplicated")

    resource_services = {
        "app_store": "AppStore",
        "brand": "Brand",
        "fomo": "Fomo",
        "google": "Google",
        "google_maps": "GoogleMaps",
        "google_play": "GooglePlay",
        "google_shopping": "GoogleShopping",
        "hacker_news": "HackerNews",
        "instagram": "Instagram",
        "reddit": "Reddit",
        "tiktok": "Tiktok",
        "web": "Web",
        "youtube": "Youtube",
    }
    methods_by_surface = {(method["service"], method["method"]): method for method in methods}
    if len(methods) != len(methods_by_surface):
        return fail("typed method surface contains duplicate service/method declarations")

    expected_surfaces = set()
    for resource, method_name, operation_id in entries:
        service = resource_services.get(resource)
        if service is None:
            return fail(f"registry entry has unknown resource: {resource}.{method_name}")
        public_method = method_name[:1].upper() + method_name[1:]
        surface = (service, public_method)
        expected_surfaces.add(surface)
        method = methods_by_surface.get(surface)
        if method is None:
            return fail(f"registry entry without typed method: {resource}.{method_name}")
        if method["operation"] != operation_id:
            return fail(
                f"{resource}.{method_name} maps to {method['operation']}, want {operation_id}"
            )
        if operation_id not in canonical:
            return fail(f"{resource}.{method_name} references unknown operation {operation_id}")
        canonical_path, canonical_parameters = canonical[operation_id]
        if method["path"] != canonical_path:
            return fail(
                f"{resource}.{method_name} path is {method['path']}, want canonical {canonical_path}"
            )
        public_parameters = parameter_structs.get(method["params"])
        if public_parameters is None:
            return fail(f"{resource}.{method_name} uses missing parameter type {method['params']}")
        if public_parameters != canonical_parameters:
            return fail(
                f"{resource}.{method_name} parameter metadata is {public_parameters}, "
                f"want canonical {canonical_parameters}"
            )

    unexpected_surfaces = sorted(set(methods_by_surface) - expected_surfaces)
    if unexpected_surfaces:
        formatted = ", ".join(f"{service}.{method}" for service, method in unexpected_surfaces)
        return fail("typed methods missing from PublicOperationRegistry: " + formatted)

    for resource in ("google", "web"):
        service = resource_services[resource]
        alias = methods_by_surface.get((service, "Search"))
        if alias is None:
            return fail(f"{resource}.search alias is missing")
        if alias["operation"] != "googleSearch":
            return fail(f"{resource}.search must map to googleSearch")
        canonical_path, canonical_parameters = canonical["googleSearch"]
        if alias["path"] != canonical_path:
            return fail(f"{resource}.search alias path drifted from {canonical_path}")
        if alias["params"] != "GoogleSearchParams":
            return fail(f"{resource}.search must use GoogleSearchParams")
        if parameter_structs[alias["params"]] != canonical_parameters:
            return fail(f"{resource}.search alias parameter mapping drifted")

    print(
        f"Surface coverage OK: {len(canonical)}/{len(canonical)} canonical GET operations; "
        f"{len(entries)} registry entries; {len(methods)} typed methods."
    )
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
