<p align="center"><img src=".github/hero.svg" alt="sdk" width="880"></p>

# Hanzo SDKs

The index for the Hanzo SDKs. Each language has its own repository and its own
package; this is where you find the right one. Nothing ships from here.

| Language | Package | Install | Repository |
| --- | --- | --- | --- |
| Python | `hanzoai` | `pip install hanzoai` | [hanzoai/python-sdk](https://github.com/hanzoai/python-sdk) |
| TypeScript | `hanzoai` | `npm install hanzoai` | [hanzoai/js-sdk](https://github.com/hanzoai/js-sdk) |
| Go | `github.com/hanzoai/go-sdk` | `go get github.com/hanzoai/go-sdk` | [hanzo-go/sdk](https://github.com/hanzo-go/sdk) |
| Rust | `hanzo-client` | `cargo add hanzo-client` | `hanzo-rs/sdk` |
| Java | `ai.hanzo:hanzo-java-cloud` | Gradle or Maven | [hanzoai/java-sdk](https://github.com/hanzoai/java-sdk) |
| Kotlin | `ai.hanzo:hanzo-kotlin-cloud` | Gradle or Maven | [hanzo-kotlin/sdk](https://github.com/hanzo-kotlin/sdk) |
| Swift | `Hanzo` | SwiftPM, from tag `v8.0.0` | [hanzo-swift/sdk](https://github.com/hanzo-swift/sdk) |
| C++ | `hanzo::hanzo` | CMake `FetchContent`, tag `v8.0.0` | [hanzo-cpp/sdk](https://github.com/hanzo-cpp/sdk) |

Three things that do not fit in a cell:

- The Go module path stays `github.com/hanzoai/go-sdk` while the repository sits
  at `hanzo-go/sdk`. Go resolves the module path, not the repository URL, so use
  the line above.
- Java and Kotlin are not on Maven Central yet, and C++ and Swift have no
  registry to be on. All four resolve from their repositories, and each README
  carries the coordinate to copy.
- The Rust repository is not public. The crate is:
  [docs.rs/hanzo-client](https://docs.rs/hanzo-client).

## The parts every SDK shares

They all talk to `/v1` at `api.hanzo.ai` and they all read `HANZO_API_KEY`. Get a
key from [hanzo.ai](https://hanzo.ai):

```bash
export HANZO_API_KEY=sk-...
```

Model ids come from the gateway, which is the only authority on the catalog —
`curl https://catalog.hanzo.ai/v1/models`. Start from `zen5`, `zen5-coder`, `enso`.

If you want a terminal rather than a library, install the CLI
([hanzoai/cli](https://github.com/hanzoai/cli)):

```bash
curl -fsSL https://hanzo.sh | sh
hanzo auth login
```

## Docs

- [docs.hanzo.ai/docs/sdks](https://docs.hanzo.ai/docs/sdks) — per-language guides
- [docs.hanzo.ai/docs/api-keys](https://docs.hanzo.ai/docs/api-keys) — getting a key
- [docs.hanzo.ai/docs](https://docs.hanzo.ai/docs) — everything else

## License

MIT — see [LICENSE](LICENSE).
