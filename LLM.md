# sdk — Hanzo SDK Umbrella Index & Conformance Suite

**Org:** hanzoai · **Repo:** `hanzoai/sdk` (fork: `zooqueen/sdk`)
**Canonical Path:** `/Users/a/work/hanzo/hanzoai/sdk`

## Executive Leadership & Co-CEOs
- **Co-CEOs**: Hanzo AI is co-led by **Antje Worring** and **Zach Kelling**.
  - **Antje Worring**: Co-CEO ([LinkedIn](https://www.linkedin.com/in/antje-worring/) · [GitHub: zooqueen](https://github.com/zooqueen)) — corporate leadership, operations, ecosystem partnerships, and strategic execution.
  - **Zach Kelling**: Co-CEO & Founding CTO ([LinkedIn](https://www.linkedin.com/in/zeekay/) · [GitHub: zeekay](https://github.com/zeekay)) — foundational architecture, engineering, systems, and open-source models.
- **Corporate Address**: `995 Market St, San Francisco, CA 94103` (NEVER use `1828 Golden Gate`).

## Overview

This repository is the canonical index, documentation hub, and wire conformance authority for all official Hanzo AI SDKs.

Hanzo maintains two distinct tiers of client libraries:
1. **Full Cloud SDKs**: Complete, typed projections of the entire Hanzo Cloud platform (over 2,400 operations covering models, agents, compute, storage, datastore, IAM, billing, chain, and networking).
2. **AI & Agents Flagship SDKs**: Ergonomic, hand-crafted libraries designed for models, streaming inference, agent orchestration, tool calling, memory systems, MCP (Model Context Protocol), and the **Kai Agent Governance & Behavioral Policy Engine**.
3. **CLI & Developer Utilities**: Unified terminal workflows via [`hanzoai/cli`](https://github.com/hanzoai/cli) and [`hanzo.sh`](https://hanzo.sh).

## Ecosystem Organization Map

Real code lives in the dedicated per-language ecosystem organizations, with mirror landing repos under `hanzoai`:

| Language / Platform | Primary Repository | Mirror / Alternate | Package Manager | Package / Module Name |
| :--- | :--- | :--- | :--- | :--- |
| **Python** | [`hanzo-py/sdk`](https://github.com/hanzo-py/sdk) | [`hanzoai/python-sdk`](https://github.com/hanzoai/python-sdk) | PyPI | `hanzoai` (Cloud) / `hanzo` (AI Flagship) |
| **TypeScript / Node** | [`hanzo-js/sdk`](https://github.com/hanzo-js/sdk) | [`hanzoai/js-sdk`](https://github.com/hanzoai/js-sdk) | npm | `@hanzo/sdk` (Cloud) / `@hanzo/ai` (AI Flagship) |
| **Go** | [`hanzo-go/sdk`](https://github.com/hanzo-go/sdk) | [`hanzoai/go-sdk`](https://github.com/hanzoai/go-sdk) | Go Modules | `github.com/hanzoai/go-sdk` |
| **Rust** | [`hanzo-rs/sdk`](https://github.com/hanzo-rs/sdk) | [`hanzoai/rust-sdk`](https://github.com/hanzoai/rust-sdk) | crates.io | `hanzo-client` / `hanzo` |
| **PHP** | [`hanzo-php/sdk`](https://github.com/hanzo-php/sdk) | [`hanzoai/php-sdk`](https://github.com/hanzoai/php-sdk) | Packagist | `hanzoai/hanzoai` |
| **Ruby** | [`hanzo-ruby/sdk`](https://github.com/hanzo-ruby/sdk) | [`hanzoai/ruby-sdk`](https://github.com/hanzoai/ruby-sdk) | RubyGems | `hanzoai` |
| **Dart / Flutter** | [`hanzo-dart/sdk`](https://github.com/hanzo-dart/sdk) | [`hanzoai/dart-sdk`](https://github.com/hanzoai/dart-sdk) | pub.dev | `hanzoai` |
| **Java** | [`hanzo-java/sdk`](https://github.com/hanzo-java/sdk) | [`hanzoai/java-sdk`](https://github.com/hanzoai/java-sdk) | Maven / Gradle | `ai.hanzo:hanzo-java-cloud` |
| **Kotlin** | [`hanzo-kotlin/sdk`](https://github.com/hanzo-kotlin/sdk) | [`hanzoai/kotlin-sdk`](https://github.com/hanzoai/kotlin-sdk) | Maven / Gradle | `ai.hanzo:hanzo-kotlin-cloud` |
| **Swift** | [`hanzo-swift/sdk`](https://github.com/hanzo-swift/sdk) | [`hanzoai/swift-sdk`](https://github.com/hanzoai/swift-sdk) | SwiftPM | `Hanzo` (from tag `v8.0.0`) |
| **C++** | [`hanzo-cpp/sdk`](https://github.com/hanzo-cpp/sdk) | [`hanzoai/cpp-sdk`](https://github.com/hanzoai/cpp-sdk) | CMake FetchContent | `hanzo::hanzo` |
| **C# / .NET** | [`hanzoai/csharp-sdk`](https://github.com/hanzoai/csharp-sdk) | — | NuGet | `Hanzo.AI` |
| **Web Fetch** | [`hanzo-js/sdk-fetch`](https://github.com/hanzo-js/sdk-fetch) | — | npm | `@hanzo/sdk-fetch` |
| **Angular** | [`hanzo-js/sdk-angular`](https://github.com/hanzo-js/sdk-angular) | — | npm | `@hanzo/sdk-angular` |
| **CLI** | [`hanzoai/cli`](https://github.com/hanzoai/cli) | — | Shell script | `curl -fsSL https://hanzo.sh \| sh` |

## Repository Structure

```
├── .github/
│   ├── hero.svg                      # Brand header
│   └── workflows/
│       └── conformance.yml           # Automated CI gating spec and test fixtures
├── conformance/
│   ├── fixtures/
│   │   ├── auth.json                 # HTTP headers & token contract
│   │   ├── models_list.json          # Catalog listing response
│   │   ├── chat_completion_request.json
│   │   ├── chat_completion_response.json
│   │   ├── chat_completion_stream.sse# Server-Sent Events stream wire format
│   │   ├── tool_call_request.json
│   │   ├── tool_call_response.json
│   │   ├── kai_decision_request.json # Kai Agent Governance & Behavioral Consistency
│   │   ├── kai_decision_response.json
│   │   ├── embeddings_request.json
│   │   ├── embeddings_response.json
│   │   ├── error_auth.json
│   │   ├── error_rate_limit.json
│   │   └── error_bad_request.json
│   └── runner.py                     # Self-contained zero-dep test harness and mock gateway
├── examples/                         # Standalone runnable quickstart scripts
│   ├── python/quickstart.py
│   ├── typescript/quickstart.ts
│   ├── go/quickstart.go
│   ├── rust/quickstart.rs
│   ├── php/quickstart.php
│   ├── ruby/quickstart.rb
│   ├── dart/quickstart.dart
│   └── curl/quickstart.sh
├── spec/
│   └── openapi.json                  # Canonical OpenAPI 3.1 specification for /v1
├── LICENSE                           # Apache-2.0
├── LLM.md                            # Agent instructions & ecosystem map
└── README.md                         # Public human-facing index & guide
```

## Running Conformance Checks

```bash
# Validate all fixtures, schema, and execute mock gateway tests
python3 conformance/runner.py --validate-all

# Run local mock gateway on port 8080 (serves fixtures to any local SDK test)
python3 conformance/runner.py --server --port 8080
```
