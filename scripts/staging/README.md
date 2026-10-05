# Staging for vkc — why there is no deployable staging target

vkc is a stateless single-binary CLI with no server component; it deploys nowhere and holds no persistent environment of its own. A "staging instance" cannot exist in the deploy.sh sense.

**Alternative validation the project uses instead** (per the two-stage pipeline mandate):

- **Prerelease artifacts as the staging stage:** prerelease tags (`vX.Y.Z-rc.N`) and their binaries are the release candidate; `main` always builds a releasable binary.
- **E2E smoke tests** (`scripts/e2e-smoke.sh` once wired): run the built binary's read-only commands against a declared Vikunja instance; write-operations are only exercised against instances explicitly configured for testing.
- **The exit-code contract** (0/1/2/3, see `docs/USER-EXPERIENCE.md`) is part of e2e validation, since agent automation depends on it.

Until CI wiring lands, `go test ./...` plus the e2e smoke script are the gate.
