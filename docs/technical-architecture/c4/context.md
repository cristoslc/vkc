# C4: System Context

```mermaid
flowchart LR
    user((Human user\nor agent))
    vkc["vkc\nsingle-binary CLI"]
    vikunja["Vikunja instance\nREST API + docs.json"]

    user -->|types commands| vkc
    vkc -->|"Bash/JSON out"| user
    vkc -->|"HTTPS: docs.json + API calls\nBearer token"| vikunja
```

## Actors and systems

- **User** — a human or a coding agent orchestrating Vikunja from a terminal.
- **vkc** — the system under design; discovers its own command surface from the spec instead of embedding it (ADR-0004).
- **Vikunja instance** — external system of record; every instance serves its own spec, so vkc works against any instance it is pointed at.

Trust boundary: the bearer token crosses to Vikunja on every API call; the spec fetch is an unauthenticated, read-only GET.
