# vkc — Agent Guidance

This file supplements `~/.agents/AGENTS.md` (the global rules). It also depends on the project's intent; start with `PURPOSE.md`.

## Orientation

- `README.md` — user-facing landing page
- `docs/ARCHITECTURE.md` → `docs/technical-architecture/` — how vkc is built
- `docs/DOMAIN-MODEL-L1.md` → `docs/domain-architecture/` — what vkc's domain means (spec discovery, command resolution, auth, transport, output)
- `docs/plans/0001-vkc-mvp.md` — current roadmap
- `docs/adr/` — why the stack looks the way it does

## Project-specific rules

- **No hardcoded Vikunja knowledge in the command layer.** Command names, flags, and help text must derive from the OpenAPI spec at runtime. Fixed tables of endpoints are a design violation (ADR-0004).
- **The spec is Swagger 2.0.** Every instance serves it at `/api/v1/docs.json`. Do not assume OpenAPI 3 fields (`openapi`, `paths[].servers`, OAS3 components); read from `definitions/`, `basePath`, and Swagger-2-style parameters. If OAS3 normalization is ever needed, do it once at cache time (ADR-0002).
- **The heuristic layers are owner-maintained.** The spec has no `operationId`s (resource+verb naming is derived from path+summary text) and no `required` arrays (help cannot mark required body fields). When a command name is wrong, fix the naming heuristic and add a test with the real spec fixture — do not special-case the endpoint.
- **Fixtures, not live calls, in unit tests.** Store trimmed copies of the real spec (`internal/fixtures/`) for tests; live `try.vikunja.io` calls belong in e2e only.

## Test command

```sh
go test ./...
```

Run this before every commit. A failing suite means commit, not push.
