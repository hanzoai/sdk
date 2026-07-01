# gen/ — spec-driven SDK generation (the openapi-generator half of hanzoai/sdk).
# Replaces Stainless: one unified spec → an idiomatic client per language.
SPEC_URL ?= https://raw.githubusercontent.com/hanzoai/openapi/main/hanzo.yaml
LANGS    ?= typescript go python rust
ALL_LANGS = typescript go python rust java kotlin ruby php csharp swift

.PHONY: all every spec clean list $(ALL_LANGS)
all: spec ; @LANGS="$(LANGS)" ./scripts/generate.sh
every: spec ; @LANGS="$(ALL_LANGS)" ./scripts/generate.sh
spec: ; @mkdir -p spec && curl -fsSL "$(SPEC_URL)" -o spec/hanzo.yaml && echo "fetched spec ($$(grep -c '^  /' spec/hanzo.yaml) paths)"
$(ALL_LANGS): spec ; @LANGS="$@" ./scripts/generate.sh
list: ; @echo "supported: $(ALL_LANGS)"; @echo "default:   $(LANGS)"
clean: ; @rm -rf clients
