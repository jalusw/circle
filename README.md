# circle

Agent-orchestration platform for the full SDLC, TDD-driven.
Submit a user story → it is planned into tasks with acceptance tests
(failing tests first) → tech/platform options are analyzed against the
requirements → coding agents (opencode, claude code, raw LLM) implement
them in git worktrees red → green → refactor → tests gate progress →
human QA sign-off → diffs are reviewed → merge.

```mermaid
flowchart TB
    Web["Web UI<br/>(htmx / react)"]
    TUI["TUI<br/>(Bubble Tea)"]
    CLI["CLI<br/>(circle)"]

    subgraph Server["circle-server"]
        Intake["Ideation<br/>stories / PRDs"]
        Plan["Planning<br/>story → tasks<br/>acceptance tests written first"]
        Design["Design · tech analysis<br/>stack · platform · libraries<br/>→ ADR"]
        Queue["Queue<br/>(SQLite)"]
        Orch["Orchestrator<br/>budgets · retries"]
        Test["Testing · TDD gate<br/>red → green → refactor"]
        QA["QA / QC<br/>human acceptance sign-off"]
        Review["Review<br/>diffs · approvals"]
        Merge["Merge<br/>back to main"]
        Intake --> Plan --> Design --> Queue --> Orch --> Test --> QA --> Review --> Merge
        Test -->|"red — feed errors<br/>retry (cap N)"| Orch
        QA -->|"criteria not met"| Orch
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
    Test -->|"green — diff extracted"| Worktrees
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
