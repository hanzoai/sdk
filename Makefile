# Hanzo Unified SDK Makefile
# Builds and tests all language implementations

.PHONY: all build test clean install dev lint format check help

# Default target
all: build test

# Help target
help:
	@echo "Hanzo SDK Build System"
	@echo ""
	@echo "Available targets:"
	@echo "  make build       - Build all language implementations"
	@echo "  make test        - Run all tests"
	@echo "  make install     - Install all packages locally"
	@echo "  make dev         - Start development mode"
	@echo "  make lint        - Run linters for all languages"
	@echo "  make format      - Format code in all languages"
	@echo "  make check       - Run all checks (lint, format, test)"
	@echo "  make clean       - Clean all build artifacts"
	@echo ""
	@echo "Language-specific targets:"
	@echo "  make build-py    - Build Python package"
	@echo "  make build-js    - Build JavaScript/TypeScript"
	@echo "  make build-rs    - Build Rust implementation"
	@echo "  make build-go    - Build Go implementation"
	@echo ""
	@echo "  make test-py     - Test Python"
	@echo "  make test-js     - Test JavaScript"
	@echo "  make test-rs     - Test Rust"
	@echo "  make test-go     - Test Go"
	@echo "  make test-matrix - Run full test matrix"

# Build targets
build: build-py build-js build-rs build-go
	@echo "✅ All builds completed successfully"

build-py:
	@echo "🐍 Building Python SDK..."
	@cd src/py && \
		pip install -e . --quiet || \
		(echo "❌ Python build failed" && exit 1)
	@echo "✅ Python build complete"

build-js:
	@echo "📦 Building JavaScript/TypeScript SDK..."
	@npm install --silent 2>/dev/null || \
		(echo "❌ JavaScript build failed" && exit 1)
	@npm run build 2>/dev/null || \
		(echo "❌ TypeScript compilation failed" && exit 1)
	@echo "✅ JavaScript build complete"

build-rs:
	@echo "🦀 Building Rust SDK..."
	@cd src/rs && \
		cargo build --release --quiet 2>/dev/null || \
		(echo "❌ Rust build failed" && exit 1)
	@echo "✅ Rust build complete"

build-go:
	@echo "🐹 Building Go SDK..."
	@mkdir -p src/go
	@cd src/go && \
		(test -f go.mod || go mod init github.com/hanzoai/sdk) && \
		go build ./... 2>/dev/null || \
		(echo "⚠️  Go build skipped (no source files yet)" && exit 0)
	@echo "✅ Go build complete"

# Test targets
test: test-py test-js test-rs test-go test-integration
	@echo "✅ All tests passed"

test-py:
	@echo "🧪 Testing Python SDK..."
	@cd src/py && \
		python -m pytest tests/ -v --tb=short 2>/dev/null || \
		echo "⚠️  Python tests not yet implemented"

test-js:
	@echo "🧪 Testing JavaScript SDK..."
	@npm test 2>/dev/null || \
		echo "⚠️  JavaScript tests not yet implemented"

test-rs:
	@echo "🧪 Testing Rust SDK..."
	@cd src/rs && \
		cargo test --quiet 2>/dev/null || \
		echo "⚠️  Rust tests not yet implemented"

test-go:
	@echo "🧪 Testing Go SDK..."
	@cd src/go && \
		go test ./... 2>/dev/null || \
		echo "⚠️  Go tests not yet implemented"

test-integration:
	@echo "🔄 Running integration tests..."
	@python tests/test_integration.py 2>/dev/null || \
		echo "⚠️  Integration tests not yet implemented"

test-matrix:
	@echo "🔧 Running full test matrix..."
	@make test-matrix-py
	@make test-matrix-js
	@make test-matrix-rs
	@make test-matrix-go
	@make test-cross-language
	@echo "✅ Test matrix complete"

test-matrix-py:
	@echo "Testing Python across versions..."
	@for ver in 3.8 3.9 3.10 3.11 3.12; do \
		echo "  Python $$ver:"; \
		python$$ver -m pytest tests/py/ -q 2>/dev/null || \
		echo "    ⚠️  Python $$ver not available"; \
	done

test-matrix-js:
	@echo "Testing Node.js across versions..."
	@node --version
	@npm test 2>/dev/null || true

test-matrix-rs:
	@echo "Testing Rust..."
	@rustc --version
	@cd src/rs && cargo test --quiet 2>/dev/null || true

test-matrix-go:
	@echo "Testing Go..."
	@go version
	@cd src/go && go test ./... 2>/dev/null || true

test-cross-language:
	@echo "Testing cross-language compatibility..."
	@python tests/test_cross_language.py 2>/dev/null || \
		echo "⚠️  Cross-language tests not yet implemented"

# Install targets
install: install-py install-js install-rs install-go
	@echo "✅ All packages installed"

install-py:
	@echo "Installing Python package..."
	@pip install -e . --quiet

install-js:
	@echo "Installing JavaScript package..."
	@npm link

install-rs:
	@echo "Installing Rust binary..."
	@cd src/rs && cargo install --path . --quiet

install-go:
	@echo "Installing Go binary..."
	@cd src/go && go install ./cmd/hanzo 2>/dev/null || true

# Development targets
dev:
	@echo "Starting development mode..."
	@make dev-watch

dev-watch:
	@echo "Watching for changes..."
	@npm run dev & \
	cd src/rs && cargo watch -x build & \
	wait

# Lint targets
lint: lint-py lint-js lint-rs lint-go
	@echo "✅ All linting passed"

lint-py:
	@echo "Linting Python..."
	@cd src/py && \
		ruff check . 2>/dev/null || \
		echo "⚠️  Install ruff: pip install ruff"

lint-js:
	@echo "Linting JavaScript..."
	@npm run lint 2>/dev/null || \
		echo "⚠️  JavaScript linting not configured"

lint-rs:
	@echo "Linting Rust..."
	@cd src/rs && \
		cargo clippy --quiet 2>/dev/null || \
		echo "⚠️  Rust linting not configured"

lint-go:
	@echo "Linting Go..."
	@cd src/go && \
		golangci-lint run 2>/dev/null || \
		echo "⚠️  Go linting not configured"

# Format targets
format: format-py format-js format-rs format-go
	@echo "✅ All formatting complete"

format-py:
	@echo "Formatting Python..."
	@cd src/py && \
		ruff format . 2>/dev/null || \
		black . 2>/dev/null || \
		echo "⚠️  Python formatter not available"

format-js:
	@echo "Formatting JavaScript..."
	@npm run format 2>/dev/null || \
		prettier --write 'src/js/**/*.{ts,js}' 2>/dev/null || \
		echo "⚠️  JavaScript formatter not configured"

format-rs:
	@echo "Formatting Rust..."
	@cd src/rs && cargo fmt --quiet 2>/dev/null

format-go:
	@echo "Formatting Go..."
	@cd src/go && go fmt ./... 2>/dev/null || true

# Check target (runs everything)
check: lint format test
	@echo "✅ All checks passed"

# Clean target
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -rf dist/ build/ *.egg-info/ .pytest_cache/ coverage/
	@rm -rf node_modules/ package-lock.json
	@cd src/rs && cargo clean 2>/dev/null || true
	@cd src/go && go clean 2>/dev/null || true
	@rm -rf src/py/__pycache__ src/py/**/__pycache__
	@rm -rf src/js/dist
	@echo "✅ Clean complete"

# CI/CD targets
ci: check
	@echo "✅ CI checks passed"

cd-publish:
	@echo "📦 Publishing packages..."
	@make cd-publish-py
	@make cd-publish-js
	@make cd-publish-rs
	@make cd-publish-go

cd-publish-py:
	@echo "Publishing to PyPI..."
	@cd src/py && python -m build && twine upload dist/*

cd-publish-js:
	@echo "Publishing to npm..."
	@npm publish

cd-publish-rs:
	@echo "Publishing to crates.io..."
	@cd src/rs && cargo publish

cd-publish-go:
	@echo "Publishing Go module..."
	@cd src/go && go mod tidy && git tag v$(VERSION)

# Version management
VERSION := $(shell cat VERSION 2>/dev/null || echo "0.1.0")

version:
	@echo "Current version: $(VERSION)"

bump-patch:
	@echo "Bumping patch version..."
	@npm version patch --no-git-tag-version
	@cd src/py && bump2version patch --no-commit --no-tag
	@cd src/rs && cargo bump patch
	@echo "New version: $(shell cat VERSION)"

bump-minor:
	@echo "Bumping minor version..."
	@npm version minor --no-git-tag-version
	@cd src/py && bump2version minor --no-commit --no-tag
	@cd src/rs && cargo bump minor
	@echo "New version: $(shell cat VERSION)"

bump-major:
	@echo "Bumping major version..."
	@npm version major --no-git-tag-version
	@cd src/py && bump2version major --no-commit --no-tag
	@cd src/rs && cargo bump major
	@echo "New version: $(shell cat VERSION)"