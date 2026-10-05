# Developer Workflows

## Build

```sh
go build ./...        # compile check
go build -o vkc .     # binary to cwd
```

## Test

```sh
go test ./...         # declared test command (also in AGENTS.md)
go test -race ./...   # before merging anything touching transport
```

Unit tests use fixtures (trimmed real spec copies under `internal/fixtures/`); live-instance checks are e2e only.

## Deploy / release

No servers to deploy. Releases are git tags (`vX.Y.Z`, prereleases `vX.Y.Z-rc.N` as the staging stage); see `scripts/staging/README.md`.

## CI/CD

Intended: GitHub Actions running `go test ./...` on PRs (not wired yet).

## Local dev

```sh
go run . --help
VKC_SPEC_URL=… vkc …   # planned: point vkc at an arbitrary spec URL for testing
```

The spec cache is disposable: `rm -rf ~/.cache/vkc` never loses user data.
