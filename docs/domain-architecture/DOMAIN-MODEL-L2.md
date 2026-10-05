# Domain Model — L2

## Entities and value objects

```mermaid
erDiagram
    SERVER_PROFILE ||--o| STORED_TOKEN : authenticates
    SERVER_PROFILE ||--o| CACHED_SPEC : describes
    CACHED_SPEC ||--|{ PARSED_SPEC : parses-into
    PARSED_SPEC ||--|{ OPERATION : contains
    OPERATION }|--|| HTTP_CALL : resolves-to
    SERVER_PROFILE {
        string host PK
        string alias
    }
    STORED_TOKEN {
        string serverHost FK
        bytes secret
        datetime storedAt
    }
    CACHED_SPEC {
        string host FK
        datetime fetchedAt
        string etag
    }
    PARSED_SPEC {
        string version "2.0 expected"
        string basePath
    }
    OPERATION {
        string method
        string path
        string summary
        string tag
        string resourceName "derived"
        string verbName "derived"
    }
    HTTP_CALL {
        string method
        string url
        string query
        json body
    }
```

## Ubiquitous language (developer layer)

| Term | Definition | Invariant |
|---|---|---|
| **Server profile** | a host + alias pair the operator logs into | exactly one token and one cache entry per profile |
| **Operation** | one method+path pair from the spec with its summary, tags, params, body schema | immutable once parsed |
| **Resource** | a grouping of operations derived from per-op tags or path segments | names must be plural, lowercase, collision-free |
| **Verb** | the second command level, derived from summary text or method semantics | every resource has ≥1 verb; ambiguous summaries fail loudly |
| **Passthrough** | everything after `--` or the first unknown flag, forwarded verbatim | never rewritten or dropped |
| **Escape hatch** | the `vkc raw METHOD /path` command shape | bypasses naming entirely; still goes through Auth and Transport |

## Aggregates

- **`ServerProfile` aggregate** — profile, its token, and its cached spec form one consistency unit: `vkc logout` must clear all three artifacts for that host.
- **`ParsedSpec` aggregate** — the immutable parsed result; command resolution never mutates it.

## Invariants

1. A `Command Resolution` decision may only reference operations from the `ParsedSpec` of the profile being addressed.
2. `Passthrough` bytes are unchanged between the user's argv and the outgoing request.
3. Help text must be traceable to spec content (description fields); hand-written help is a bug.
