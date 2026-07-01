#!/usr/bin/env bash
# Generate every Hanzo API-client SDK from the ONE unified OpenAPI spec, using
# open-source openapi-generator. This is the spec-driven half of hanzoai/sdk
# (the hand-written `hanzo` CLI lives at the repo root); it replaces Stainless.
#
#   ./scripts/generate.sh                 # default languages
#   LANGS="typescript go" ./scripts/generate.sh
#   SPEC=/path/to/hanzo.yaml ./scripts/generate.sh
set -euo pipefail
cd "$(dirname "$0")/.."

SPEC="${SPEC:-spec/hanzo.yaml}"
OUT="${OUT:-clients}"
VERSION="${VERSION:-$(cat VERSION 2>/dev/null || echo 0.1.0)}"

# lang:generator — the canonical matrix. Add a line to add a language.
MATRIX="
typescript:typescript-axios
go:go
python:python
rust:rust
java:java
kotlin:kotlin
ruby:ruby
php:php
csharp:csharp
swift:swift5
"
LANGS="${LANGS:-typescript go python rust}"

[ -f "$SPEC" ] || { echo "spec not found: $SPEC — run 'make spec' first" >&2; exit 1; }

gen() {
  if command -v openapi-generator-cli >/dev/null 2>&1; then openapi-generator-cli "$@"
  elif command -v openapi-generator >/dev/null 2>&1; then openapi-generator "$@"
  elif command -v docker >/dev/null 2>&1 && docker info >/dev/null 2>&1; then
    docker run --rm -v "$PWD:/local" -w /local openapitools/openapi-generator-cli:v7.14.0 "$@"
  else npx -y @openapitools/openapi-generator-cli@2.20.0 "$@"; fi
}

pkg_props() {
  case "$1" in
    typescript-axios) echo "npmName=@hanzo/sdk,npmVersion=$VERSION,supportsES6=true" ;;
    go)               echo "packageName=hanzo,packageVersion=$VERSION,withGoMod=true" ;;
    python)           echo "packageName=hanzo,projectName=hanzo,packageVersion=$VERSION,library=urllib3" ;;
    rust)             echo "packageName=hanzo,packageVersion=$VERSION,supportAsync=true,library=reqwest" ;;
    java)             echo "groupId=ai.hanzo,artifactId=hanzo-sdk,artifactVersion=$VERSION,library=native" ;;
    kotlin)           echo "groupId=ai.hanzo,artifactId=hanzo-sdk,artifactVersion=$VERSION,library=jvm-okhttp4" ;;
    ruby)             echo "gemName=hanzo,moduleName=Hanzo,gemVersion=$VERSION" ;;
    php)              echo "invokerPackage=Hanzo,packageName=hanzo,artifactVersion=$VERSION" ;;
    csharp)           echo "packageName=Hanzo,packageVersion=$VERSION,targetFramework=net8.0,library=httpclient" ;;
    swift5)           echo "projectName=Hanzo,podVersion=$VERSION,responseAs=AsyncAwait" ;;
    *)                echo "packageVersion=$VERSION" ;;
  esac
}

generator_for() { echo "$MATRIX" | awk -F: -v l="$1" '$1==l{print $2}'; }

rc=0
for lang in $LANGS; do
  generator="$(generator_for "$lang")"
  [ -z "$generator" ] && { echo "!! unknown language: $lang"; rc=1; continue; }
  dest="$OUT/$lang"; cfg="config/$lang.yaml"; extra=""
  [ -f "$cfg" ] && extra="-c $cfg"
  echo "==> $lang ($generator) → $dest"
  rm -rf "$dest"; mkdir -p "$dest"
  if gen generate -i "$SPEC" -g "$generator" -o "$dest" \
        --additional-properties="$(pkg_props "$generator")" \
        --git-user-id hanzoai --git-repo-id sdk \
        $extra >/dev/null 2>"$dest/.gen.log"; then
    echo "    ok"
  else
    echo "    FAILED (see $dest/.gen.log)"; tail -3 "$dest/.gen.log" 2>/dev/null | sed 's/^/    /'; rc=1
  fi
done
echo; echo "done — SDKs in $OUT/ for: $LANGS"
exit $rc
