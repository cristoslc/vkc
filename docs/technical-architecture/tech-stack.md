# Tech Stack

## Language and toolchain

- **Go 1.27+** (ADR-0001: single static binary, fast startup, no runtime dependencies for users).
- LSP config is the `go.mod` created by `go mod init github.com/cristoslc/vkc`.
- Standard library only until a hard need appears (`net/http`, `encoding/json`, `time`, `os/exec` are sufficient for the MVP). A CLI framework (cobra) is permitted if dynamic command ergonomics demand it — decide by evidence, not by habit (ADR-0001).

## External systems

- **Vikunja API** — the only external system. Spec endpoint `/api/v1/docs.json` (Swagger 2.0), API base path `/api/v1`, Bearer token auth. Spec format caveats in ADR-0002.

## Local state

- **Spec cache:** `~/.cache/vkc/<server-host>/docs.json` + fetched-at timestamp, TTL-based refresh.
- **Token store:** platform keychain where available (macOS Keychain), with a plain-file fallback under `~/.config/vkc/` (chmod 600) documented as less safe. Tokens are never written to stdout or logs.

## Infrastructure

None: no CI runners declared yet (GitHub Actions is the intended default once the MVP lands), no containers, no databases, no message queues.
