#!/usr/bin/env python3
"""Verify the handwritten public registry covers every canonical GET operation."""

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


def main() -> int:
    try:
        spec = json.loads(SPEC.read_text(encoding="utf-8"))
        source = RESOURCES.read_text(encoding="utf-8")
    except (OSError, json.JSONDecodeError) as exc:
        return fail(str(exc))

    canonical = [
        operation["operationId"]
        for path_item in spec.get("paths", {}).values()
        for method, operation in path_item.items()
        if method.lower() == "get" and operation.get("operationId")
    ]
    duplicates = sorted({item for item in canonical if canonical.count(item) > 1})
    if duplicates:
        return fail("canonical GET operation IDs are not unique: " + ", ".join(duplicates))

    registry_start = source.find("var PublicOperationRegistry")
    if registry_start < 0:
        return fail("PublicOperationRegistry declaration not found")
    registry_source = source[registry_start:]
    entries: list[tuple[str, str, str]] = []
    current_resource: str | None = None
    for line in registry_source.splitlines():
        group = re.match(r'^\s*"([^"\\]+)":\s*\{\s*$', line)
        if group:
            current_resource = group.group(1)
            continue
        entry = re.match(r'^\s+"([^"\\]+)":\s*"([^"\\]+)",\s*$', line)
        if entry and current_resource is not None:
            entries.append((current_resource, entry.group(1), entry.group(2)))

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

    if not any(resource == "google" and method == "search" and operation_id == "googleSearch" for resource, method, operation_id in entries):
        return fail("google.search must map to googleSearch")
    if not any(resource == "web" and method == "search" and operation_id == "googleSearch" for resource, method, operation_id in entries):
        return fail("web.search must explicitly map to the googleSearch compatibility alias")

    method_pattern = re.compile(
        r"^func \(s \*(?P<service>[A-Za-z0-9]+)Service\) (?P<method>[A-Za-z0-9]+)"
        r"\(ctx context\.Context, params [A-Za-z0-9]+Params\) \(\*Response, error\) \{",
        re.MULTILINE,
    )
    methods = {(match.group("service"), match.group("method")) for match in method_pattern.finditer(source)}
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
    missing_methods = []
    for resource, method, _ in entries:
        service = resource_services.get(resource)
        if service is None:
            missing_methods.append(f"{resource}.{method} (unknown resource)")
            continue
        expected_method = method[:1].upper() + method[1:]
        if (service, expected_method) not in methods:
            missing_methods.append(f"{resource}.{method}")
    if missing_methods:
        return fail("registry entries without typed methods: " + ", ".join(missing_methods))

    print(f"Surface coverage OK: {len(canonical)}/{len(canonical)} canonical GET operations; {len(entries)} registry entries.")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
