# ADR-0001: Implement vkc in Go as a single static binary

## Context

vkc is a CLI users type many times a day, often inside agent automation where every spawn pays startup cost.

## Options

1. **Go** — static binary, instant startup, strong CLI ecosystem (cobra), painless binary distribution.
2. **TypeScript/Node** — fastest iteration for the maintainer, npm distribution, but runtime dependency and slower startup.
3. **Python** — simplest code, but pipx distribution and the slowest startup of the three.

## Decision

Go, standard library first; a CLI framework is only adopted when the dynamic command tree shows stdlib flag handling is no longer ergonomic.

## Consequences

Users install one file; no runtime deps. Iteration cost on CLI ergonomics is higher than TypeScript but bounded by the small component surface. LSP config = `go.mod`.
