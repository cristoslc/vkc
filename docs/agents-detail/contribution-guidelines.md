# Contribution Guidelines (agent-facing spoke)

Human-facing process lives in `CONTRIBUTING.md` at the repo root; this spoke covers only what a coding agent must do differently.

## Before every PR

1. Read the project `AGENTS.md` and check the ADRs touched by the change.
2. Run the declared test command (`go test ./...`) and record the result in the PR body.
3. For bug fixes: write the reproduction test first, watch it fail, then fix, then watch it pass. Note the failing-test commit in the PR.

## Worktree discipline

- Project root stays on `main`; feature work happens in `.worktrees/<topic>/` or a topic branch.
- `.worktrees/` is gitignored and never committed.

## PR conventions

- Conventional-commit titles (`feat:` / `fix:` / `docs:` / `test:` / `chore:`).
- Body references related issues by number and title, lists evidence (test output), and states any spec-version assumptions (`try.vikunja.io` spec dated in the commit message when fixtures were refreshed).
- Every PR carries a `CHANGELOG.md` `Unreleased` entry.

## Chronicling

When a thread makes a decision worth keeping outside the diff (e.g. "naming heuristic now prefers the per-op tag over the path segment"), record it in the relevant ADR or a `docs/plans/` update — not only in the commit message — so future sessions find it through the docs tree.
