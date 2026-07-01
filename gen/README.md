# gen/ — Spec-driven SDK generation

Generates a Hanzo **API-client SDK per language** from the ONE unified OpenAPI
spec ([`hanzoai/openapi`](https://github.com/hanzoai/openapi) `hanzo.yaml`, all
36 services), using open-source [openapi-generator](https://openapi-generator.tech).
This is our **Stainless replacement** — self-owned, no SaaS.

```bash
cd gen
make all       # fetch spec + generate typescript, go, python, rust
make every     # + java, kotlin, ruby, php, csharp, swift
make go        # one language
```

Output → `gen/clients/<language>/`. The engine is `scripts/generate.sh` (language
matrix + runner resolution: local `openapi-generator-cli` → Docker → `npx`).

Add a language: one `lang:generator` line in `scripts/generate.sh`.

> The hand-written cross-language **`hanzo` CLI** lives at the repo root
> (`src/{py,js,rs,go}`). This `gen/` tree is the generated **API client** — the
> two are complementary: the CLI is hand-tuned UX; the clients track the API 1:1.
