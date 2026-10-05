# Plan 0001: vkc MVP

## Goal

Operator runs `vkc task list` against try.vikunja.io (or a local instance), sees pretty JSON, without any endpoint knowledge in the code.

## Milestones

1. **Spec Discovery** — fetch + TTL cache + parse Swagger 2.0 (`internal/spec/`), fixture tests from the real `docs.json` (trimmed).
2. **Naming heuristics** — resource from per-op tags / path segments; verb from summary text / method fallback; fixture-driven (ADR-0004). Failure case: ambiguous summary text must error, not guess.
3. **Command tree** — two-level builder, flags from query/path params and body schema, `--raw`/`--jq`/`--` passthrough, `vkc find`, `vkc raw`, grouped help. No cobra until ergonomics demand it.
4. **Auth** — `vkc login/logout` storing tokens (keychain on macOS, chmod-600 file fallback), profiles per host.
5. **Transport + rendering** — request assembly with passthrough, exit codes 0/1/2/3, pretty JSON default.
6. **E2E smoke** — against a declared instance; script under `scripts/`.

## Out of scope for MVP

Generated completions (post-MVP), `--jq` sugar, stdio MCP mode, generic `mcpcli` profiles.
