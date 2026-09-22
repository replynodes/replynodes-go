#!/usr/bin/env sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
image="openapitools/openapi-generator-cli:v$(tr -d '\n' < "$root/sdk-generation/VERSION")"
mode="${1:-}"
stage="$root/tmp/generated-stage"
generated="$root/internal/generated"

case "$mode" in
  ""|--check)
    ;;
  --offline)
    ;;
  *)
    echo "usage: $0 [--check|--offline]" >&2
    exit 2
    ;;
esac

rm -rf "$stage"
mkdir -p "$stage"
trap 'rm -rf "$stage"' EXIT INT TERM

if [ "$mode" = "--offline" ]; then
  docker image inspect "$image" >/dev/null
fi

docker run --rm \
  --user "$(id -u):$(id -g)" \
  -v "$root:/local" \
  "$image" generate \
  -i /local/api-docs/replynodes-fetcher.openapi.json \
  -g go \
  -o /local/tmp/generated-stage \
  -c /local/sdk-generation/config/go.json \
  --ignore-file-override /local/sdk-generation/ignore/go.openapi-generator-ignore \
  --skip-overwrite

find "$stage" -maxdepth 1 -type f -name '*.go' -exec gofmt -w {} +

if ! find "$stage" -maxdepth 1 -type f -name '*.go' -print -quit | grep -q .; then
  echo "generator produced no Go files" >&2
  exit 1
fi

if [ "$mode" = "--check" ]; then
  drift=0
  for file in "$stage"/*.go; do
    name=$(basename "$file")
    if [ ! -f "$generated/$name" ] || ! diff -u "$generated/$name" "$file"; then
      drift=1
    fi
  done
  for file in "$generated"/*.go; do
    [ -e "$file" ] || continue
    name=$(basename "$file")
    if [ ! -f "$stage/$name" ]; then
      echo "generated file is missing from canonical output: $name" >&2
      drift=1
    fi
  done
  if [ "$drift" -ne 0 ]; then
    echo "generated Go output is out of date" >&2
    exit 1
  fi
  echo "generated Go output is up to date"
  exit 0
fi

mkdir -p "$generated"
find "$generated" -maxdepth 1 -type f -name '*.go' -delete
for file in "$stage"/*.go; do
  cp "$file" "$generated/$(basename "$file")"
done
echo "generated Go client in internal/generated"
