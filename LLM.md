# @hanzo/cli

The unified Hanzo **CLI** (`hanzo node|agent|mcp|net|dev …`). Built from
`src/js` only (`dist/`). This is the CLI, a separate concern from the client
SDKs.

## This repo is NOT the SDK generator (decided 2026-07)

The client SDKs are one repo per language, each generated from the ONE spec
`hanzoai/openapi/hanzo.yaml`:

- Python/Go/JS → `hanzoai/{python,go,js}-sdk` via Stainless project `hanzo-ai`.
- Rust → `hanzoai/rust-sdk` (hand-written).
- C++ / Dart → `hanzoai/{cpp,dart}-sdk` via openapi-generator on ARC.

The earlier `gen/` (openapi-generator over 10 langs, "replaces Stainless") and
the `src/{py,go,rs}` client copies were a SECOND way that duplicated those
repos. Retired — removed. There is one interface (`hanzo.yaml`) and one repo per
language; generator backend is an orthogonal per-language choice. This repo owns
only the CLI (`src/js`).
