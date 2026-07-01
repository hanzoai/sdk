# gen/ — spec-driven SDK generation (the openapi-generator half of hanzoai/sdk).
# Replaces Stainless: one unified spec → an idiomatic client per language.
SPEC_REPO ?= hanzoai/openapi
SPEC_PATH ?= hanzo.yaml
LANGS     ?= typescript go python rust
ALL_LANGS = typescript go python rust java kotlin ruby php csharp swift

.PHONY: all every spec clean list $(ALL_LANGS)
all: spec ; @LANGS="$(LANGS)" ./scripts/generate.sh
every: spec ; @LANGS="$(ALL_LANGS)" ./scripts/generate.sh
# Private-repo aware: with GH_TOKEN → GitHub contents API; else public raw.
spec:
	@mkdir -p spec
	@if [ -n "$$GH_TOKEN" ]; then \
	  curl -fsSL -H "Authorization: Bearer $$GH_TOKEN" -H "Accept: application/vnd.github.raw" \
	    "https://api.github.com/repos/$(SPEC_REPO)/contents/$(SPEC_PATH)?ref=main" -o spec/hanzo.yaml; \
	else curl -fsSL "https://raw.githubusercontent.com/$(SPEC_REPO)/main/$(SPEC_PATH)" -o spec/hanzo.yaml; fi
	@echo "fetched spec ($$(grep -c '^  /' spec/hanzo.yaml) paths)"
$(ALL_LANGS): spec ; @LANGS="$@" ./scripts/generate.sh
list: ; @echo "supported: $(ALL_LANGS)"; echo "default: $(LANGS)"
clean: ; @rm -rf clients
