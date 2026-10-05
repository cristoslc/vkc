# docs/ — vkc Documentation

The docs tree splits into two architectural surfaces plus cross-cutting hubs.
`AGENTS.md` at the repo root is the entry point for agent guidance; `PURPOSE.md` is the intent statement.

| Surface | Hub | Spokes |
|---|---|---|
| Technical architecture | `ARCHITECTURE.md` | `technical-architecture/` (C4, tech stack) |
| Domain architecture | `DOMAIN-MODEL-L1.md` | `domain-architecture/` (context map, events, L2) |
| Cross-cutting | `DEVELOPER-WORKFLOWS.md`, `USER-EXPERIENCE.md` | `developer-workflows/`, `user-experience/` |

Unnumbered-but-structured directories (no hub file): `adr/`, `plans/`, `musings/`, `retros/`, `tech-debt/`, `agents-detail/`.
