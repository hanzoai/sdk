# sdk

**Org:** hanzoai · **Repo:** `hanzoai/sdk` · **Contents:** `README.md` and nothing else that runs.

This repo is the index for the Hanzo SDKs — a table pointing at the per-language
repository, package name and install line. There is no build, no test, no
package published from here, and no code to import.

## Why there is no `src/`

It used to carry a four-language prototype (`src/{go,js,py,rs}`) plus `gen/`, an
openapi-generator wrapper. Every part of it claimed an identity that a live repo
already owns, so the two could only disagree:

| What was here | Who owns that identity |
| --- | --- |
| `src/go/go.mod` → `module github.com/hanzoai/go-sdk` | `hanzo-go/sdk`, and it is the module the Go proxy serves |
| `src/py` → package `hanzo` | `hanzoai/python-sdk` `pkg/hanzo` |
| `package.json` → `@hanzo/cli` | `hanzoai/cli`, published on npm |
| `src/rs` → a `hanzo` binary + client crate | `hanzoai/cli` (binary), `hanzo-rs/sdk` (crate `hanzo-client`) |
| `gen/` → generate a client per language from the spec | `hanzoai/openapi` `generate.py` + `sdks.yaml` |

Two repositories cannot both own one import path. The dedicated repo wins every
row, so `src/`, `gen/`, and the build files that existed only to compile them
were deleted. `sdks.yaml`'s own header describes what the second copy costs: two
generators disagreed about output layout and built an orphan copy of all 2143
files.

## No hanzo.yml, no .hanzo/workflows

The canonical CI pair gated `npm ci && npm run build:ts`, `npm test` and the
`src/py` tests — all of it the prototype. With the prototype gone there is
nothing here that can go red, so the pair went with it. If this repo ever gains
something buildable, the pair comes back in its canonical shape: `hanzo.yml`
plus a ~7-line caller under `.hanzo/workflows/` on the git-runner fleet, never
under `.github/workflows/` (github.com has no runner for those labels, and the
forge reads only the first of WORKFLOW_DIRS, so a file there gates nothing).

## Editing the index

The package names are facts held in the SDK repos, not here — read them from the
manifest before changing a row:

- Python: `python-sdk/pyproject.toml` `name`
- TypeScript: `js-sdk/package.json` `name`
- Go: `go-sdk/go.mod` `module`
- Rust: `rust-sdk/crates/hanzo-client/Cargo.toml` `name`
- Java: `java-sdk/hanzo-java-cloud/build.gradle` `group` + module name
- Kotlin: `kotlin-sdk/hanzo-kotlin-cloud/build.gradle.kts` `group` + module name

Versions are deliberately absent from the table: an index that restates a
version is an index that is wrong a week later. Each repo's README carries its
own current one.

`hanzoai/go-sdk` and `hanzoai/kotlin-sdk` on GitHub are redirects to
`hanzo-go/sdk` and `hanzo-kotlin/sdk`. `hanzo-kt/sdk` and `hanzo-js/sdk` are
older generator output — not the repos the published packages come from.
