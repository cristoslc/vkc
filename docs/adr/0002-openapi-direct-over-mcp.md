# ADR-0002: Go OpenAPI-direct, not MCP-wrapped

## Context

The initial concept wrapped vkc's brain in an MCP server. Community Vikunja MCP servers exist, but they derive their own tool descriptions from the same spec every Vikunja instance serves at `/api/v1/docs.json`.

## Options

1. **MCP server wrapping the API** — extra transport layer (HTTP or stdio), lossy translation at description time.
2. **OpenAPI-direct** — vkc consumes `/api/v1/docs.json` itself; spec + token + server are the only moving parts.

## Decision

OpenAPI-direct. MCP remains a possible *expose* mode later (serve vkc's derived commands as MCP tools), never the *source*.

## Consequences

One fewer process and protocol; the CLI inherits spec quality limits directly: Swagger **2.0** (no `openapi` field), no per-op `operationId` (see ADR-0004), no `required` arrays in definitions (help cannot mark required body fields; server 4xx surfaces misuse). If OAS3 tooling is ever needed, normalize once at cache time.
