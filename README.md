# vkc — Vikunja Command-line

`vkc` is a dynamic CLI for [Vikunja](https://vikunja.io). It reads the OpenAPI spec your Vikunja instance serves at `/api/v1/docs.json`, derives its commands from it at runtime, and talks to the API with a locally stored token. Nothing about Vikunja is hardcoded: if your instance exposes new endpoints, `vkc` exposes them too.

## Status

Scaffold — pre-alpha. The core loop (spec fetch → command tree generation → authenticated calls) is not implemented yet. See `docs/plans/0001-vkc-mvp.md` for the roadmap.

## Design intent

- **Two-level commands**: `vkc task list`, `vkc project create`, `vkc team update` — resources as groups, verbs beneath them.
- **Generated help**: `vkc task list --help` is rendered from the spec's parameter and model descriptions, not maintained by hand.
- **Search**: `vkc find <keyword>` greps operation summaries so obscure endpoints stay findable.
- **Escape hatch**: `vkc raw GET /tasks/42 --expand=buckets` issues any method/path directly.
- **Auth**: `vkc login` stores a Vikunja API token locally (macOS Keychain on macOS); requests go out as Bearer auth.

See `PURPOSE.md` for the one-paragraph outcome statement and `docs/ARCHITECTURE.md` / `docs/DOMAIN-MODEL-L1.md` for the technical and domain views.

## Install

```sh
go install github.com/cristoslc/vkc@latest        # once v0.1.0 exists
# or from source:
git clone https://github.com/cristoslc/vkc && cd vkc && go build -o /usr/local/bin/vkc .
```

## Usage (planned shape)

```sh
vkc                                      # resource groups with op counts
vkc task                                 # verbs for the task resource
vkc task list --help                     # parameter help from the live spec
vkc find "buckets"                       # search operation summaries
vkc raw GET /tasks/42 --expand=buckets   # direct method/path escape hatch
vkc login                                # store an API token
```

## Contributing

Read `CONTRIBUTING.md` and the project `AGENTS.md`. The test command is `go test ./...`; agent-originated bug fixes must arrive with reproduction steps and a failing test written before the fix.
