# ADR-0003: HTTP transport first; stdio second

## Context

The dynamic-CLI concept supports both a plain HTTP client and an MCP stdio mode.

## Decision

Ship HTTP-first: the operator's declared use is direct terminal use and agent orchestration via shell, which plain HTTPS serves fully. Stdio MCP mode is deferred until there is a consumer that needs it.

## Consequences

Simplest possible v0.1: fetch spec, resolve command, Bearer-authenticated request, JSON out. No server lifecycle to manage. Exit-code contract (0 success / 1 usage / 2 transport / 3 auth) defined in `USER-EXPERIENCE.md` becomes the agent-automation interface.
