# ADR-0004: Command names derive from the spec, not from a table

## Context

The Vikunja spec has no `operationId`s (171/172 missing). With 172 operations, hand-maintaining a command table violates the project's no-hardcoded-knowledge rule and would rot.

## Options

1. **Hand-maintained mapping** — precise but goes stale and violates the core rule.
2. **Derived naming** — resource from per-op tag (`task`, `project`, …) with path-segment fallback (`/projects/{id}/tasks` → `project task`); verb from summary text ("Creates a new project" → `create`, "Get one task" → `get`) with method fallback.
3. **Method/path verbatim commands** — lossless but unusable (`api-put-projects-id-tasks`).

## Decision

Derived naming (option 2), with `vkc raw` as the always-works escape hatch and `vkc find` as recovery for mis-derived names.

## Consequences

The naming heuristic is owner-maintained code with fixture tests: when a command name is wrong, the fix is in the heuristic, not an endpoint special case. Known ambiguity: summary text wins over method for POST-that-actually-updates cases. New endpoints appear automatically.
