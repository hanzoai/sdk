<p align="center"><img src=".github/hero.svg" alt="sdk" width="880"></p>

# Hanzo SDKs

This repository is the index for the Hanzo SDKs. Each language has its own repository and
its own package; this is where you find the right one.

## Install

| Language | Install | Version | Source |
| --- | --- | --- | --- |
| Python | `pip install hanzoai` | 3.1.6 | [hanzoai/python-sdk](https://github.com/hanzoai/python-sdk) |
| TypeScript | `npm install hanzoai` | 2.0.7 | [hanzo-js/sdk](https://github.com/hanzo-js/sdk) |
| Go | `go get github.com/hanzoai/go-sdk` | v1.0.1 | [hanzo-go/sdk](https://github.com/hanzo-go/sdk) |
| Rust | `cargo add hanzo-client` | 0.1.1 | [hanzo-rs/sdk](https://github.com/hanzo-rs/sdk) |
| C++ | build from source | — | [hanzo-cpp/sdk](https://github.com/hanzo-cpp/sdk) |
| Swift | SwiftPM, from the repository | — | [hanzo-swift/sdk](https://github.com/hanzo-swift/sdk) |
| Kotlin | build from source (`ai.hanzo:sdk`) | — | [hanzo-kt/sdk](https://github.com/hanzo-kt/sdk) |

The Go module path is `github.com/hanzoai/go-sdk` even though the repository now lives at
`hanzo-go/sdk` — Go resolves the module path, not the repository URL, so use the line
above. C++, Swift and Kotlin are not published to a package registry yet; build them from
their repositories.

## First call

Every SDK talks to the same `/v1` API at `api.hanzo.ai` and reads the same environment
variable. Get a key from [hanzo.ai](https://hanzo.ai), then:

```bash
export HANZO_API_KEY=sk-...
```

```python
from hanzoai import (
    ApiClient, Configuration, AiOpenAICompatibleApi,
    AiChatCompletionRequest, AiChatMessage,
)

config = Configuration(host="https://api.hanzo.ai", access_token="sk-...")

with ApiClient(config) as client:
    ai = AiOpenAICompatibleApi(client)
    resp = ai.ai_create_chat_completion(
        AiChatCompletionRequest(
            model="zen5",
            messages=[AiChatMessage(role="user", content="hello")],
        )
    )
    print(resp.choices[0].message.content)
```

Model ids come from the gateway, which is the only authority on the catalog. Browse them
with `curl https://catalog.hanzo.ai/v1/models`; `zen5`, `zen5-coder` and `enso` are the
ones to start from.

If you want a terminal rather than a library, install the CLI instead:

```bash
curl -fsSL https://hanzo.sh | sh
hanzo auth login
```

## What is in this repository

`src/` holds an older prototype of a single cross-language package. It is not what ships,
and it is not published to any registry — the per-language repositories above are the ones
under development. Read this repository as the index, not as a dependency.

## Docs

- [docs.hanzo.ai/docs/sdks](https://docs.hanzo.ai/docs/sdks) — per-language guides
- [docs.hanzo.ai/docs/api-keys](https://docs.hanzo.ai/docs/api-keys) — getting a key
- [docs.hanzo.ai/docs](https://docs.hanzo.ai/docs) — everything else

## License

MIT — see [LICENSE](LICENSE).
