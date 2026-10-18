# circle

Agent-orchestration platform for the full software lifecycle.
Submit a user story → it is broken into tasks → coding agents (opencode,
claude code, raw LLM) execute them in git worktrees → diffs are reviewed
→ merge.

```mermaid
flowchart TB
    Web["Web UI<br/>(htmx / react)"]
    TUI["TUI<br/>(Bubble Tea)"]
    CLI["CLI<br/>(circle)"]

    subgraph Server["circle-server"]
        Intake["Intake<br/>stories / PRDs"]
        Queue["Queue<br/>(SQLite)"]
        Orch["Orchestrator<br/>budgets · retries"]
        Review["Review<br/>diffs · approvals"]
        Intake --> Queue --> Orch --> Review
    end

    subgraph Runners["runner adapters (pluggable)"]
        OC["opencode<br/>(serve)"]
        CC["claude code<br/>(-p stream-json)"]
        RAW["rawllm<br/>(own loop)"]
    end

    Worktrees["git worktrees<br/>one branch per task"]

    CLI -->|HTTP+SSE| Server
    Web -->|HTTP+SSE| Server
    TUI -->|HTTP+SSE| Server
    Orch --> Runners
    Runners --> Worktrees
```

## Layout

| Path                  | Purpose                                        |
| --------------------- | ---------------------------------------------- |
| `cmd/circle`           | swiss-army CLI (submit stories, watch runs)    |
| `cmd/circle-server`    | headless HTTP API + SSE server                 |
| `cmd/circle-tui`       | terminal UI client (Bubble Tea)                |
| `internal/domain`     | pure types: Story, Task, Run, Event, states    |
| `internal/store`      | SQLite persistence + migrations                |
| `internal/orchestrator`| queue, budgets, retry loop                    |
| `internal/runner`     | Runner interface + registry + adapters         |
| `internal/gitops`     | worktree/branch/diff helpers                   |
| `internal/api`        | HTTP + SSE handlers                            |
| `web/`                | web UI (later)                                 |
| `.circle/`             | project config                                 |

## Getting started

```sh
make build        # builds bin/circle, bin/circle-server, bin/circle-tui
make run-server   # starts API on :8080
curl :8080/healthz
```

## Status

Skeleton only — runner adapters and orchestrator are stubs.
See `docs/ROADMAP.md`.
