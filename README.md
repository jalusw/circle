# cycle

Agent-orchestration platform for the full software lifecycle.
Submit a user story → it is broken into tasks → coding agents (opencode,
claude code, raw LLM) execute them in git worktrees → diffs are reviewed
→ merge.

```
        ┌──────┐   ┌──────┐
        │  Web │   │ TUI  │      dumb clients (HTTP + SSE)
        └───┬──┘   └──┬───┘
            │         │
┌───────────▼─────────▼───────────┐
│           cycle-server          │
│  intake → queue → orchestrator  │
│     → review → approval         │
└───────────┬─────────────────────┘
            │ runner adapters (pluggable)
      ┌─────┼──────────┐
 opencode  claude     rawllm
 (serve)   (-p JSON)  (own loop)
```

## Layout

| Path                  | Purpose                                        |
| --------------------- | ---------------------------------------------- |
| `cmd/cycle`           | swiss-army CLI (submit stories, watch runs)    |
| `cmd/cycle-server`    | headless HTTP API + SSE server                 |
| `cmd/cycle-tui`       | terminal UI client (Bubble Tea)                |
| `internal/domain`     | pure types: Story, Task, Run, Event, states    |
| `internal/store`      | SQLite persistence + migrations                |
| `internal/orchestrator`| queue, budgets, retry loop                    |
| `internal/runner`     | Runner interface + registry + adapters         |
| `internal/gitops`     | worktree/branch/diff helpers                   |
| `internal/api`        | HTTP + SSE handlers                            |
| `web/`                | web UI (later)                                 |
| `.cycle/`             | project config                                 |

## Getting started

```sh
make build        # builds bin/cycle, bin/cycle-server, bin/cycle-tui
make run-server   # starts API on :8080
curl :8080/healthz
```

## Status

Skeleton only — runner adapters and orchestrator are stubs.
See `docs/ROADMAP.md`.
