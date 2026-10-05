# Contributing to vkc

Thanks for contributing. This file serves both human contributors and coding agents.

## For humans

- **Issues:** please use the issue templates (blank issues are disabled); fill in environment, reproduction steps, expected result, and actual result.
- **Branches:** name them `your-topic/short-slug` or use a worktree under `.worktrees/`. The default branch is `main` and must always stay releasable.
- **Commits:** conventional commits (`feat:`, `fix:`, `docs:`, `test:`, `chore:`).
- **Pull requests:** use the PR template. Every PR adds an entry to `CHANGELOG.md` under `Unreleased` and updates affected docs.
- **DCO/CLA:** none. By opening a PR you license your contribution under the project license (MIT).

## For coding agents

- Read the project `AGENTS.md` first; it carries project-specific rules that override your defaults.
- Test command: `go test ./...` — run it before every commit and re-run before requesting review.
- Agent-originated bug fixes MUST arrive with reproduction steps and a failing test written before the fix (red-green-refactor). A fix without a first-failing test will be sent back.
- Work in `.worktrees/` for anything beyond a one-file fix; keep project root on `main`.
- Never commit tokens, keys, or spec fixtures containing real credentials.

## Review and merge policy

Maintainer (repository owner) reviews and merges. Squash-merge with a conventional-commit title is the default.
