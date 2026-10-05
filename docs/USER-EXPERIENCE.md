# User Experience

vkc's users are terminal-dwellers: humans and agents operating Vikunja. The interface is the command line, generated help, and JSON output only.

## UX principles

1. **The instance is the truth.** What vkc can do is exactly what the connected instance can do — no more, no less.
2. **Typeability beats completeness.** Two-level commands (`vkc task list`) with grouped bare `vkc` output; completion scripts generated from the spec.
3. **JSON out, jq exists.** Default output is pretty JSON the user pipes to `jq`; `--raw` avoids re-encoding; no interactive TUI is planned.
4. **Fail with a path forward.** Unknown command → suggest `vkc find <word>`; missing token → name the exact `vkc login` form.

## Key journeys

- **First contact:** bare `vkc` shows resource groups with help hints (works without login; spec from cache or direct fetch).
- **Daily add-task:** `vkc task create --list=42 --title="…"` → JSON of created task.
- **Lost operation:** `vkc find buckets` → matching operations with their commands.
- **Emergency:** `vkc raw PATCH /tasks/1 …` when naming or flags don't fit.

## Accessibility

Terminal-native: honors `NO_COLOR`; help output is width-aware; errors go to stderr with non-zero exit codes so scripts and agents can rely on exit-status contracts (0 = success, 1 = usage error, 2 = transport failure, 3 = auth failure).
