# C4: Deployment

vkc is a stateless single-binary CLI; there is no deployed service. Distribution surfaces:

- `go install github.com/cristoslc/vkc@latest`
- Release artifacts per OS/arch from GitHub Releases (planned, from v0.1.0)
- From-source build (`go build`)

"Staging" for this project means prerelease binaries and tags (`v0.x.y-rc.N`) plus e2e smoke tests against a declared Vikunja instance — see `../../../scripts/staging/README.md`. Validation cannot damage any persistent environment beyond the target Vikunja instance the user points at.
