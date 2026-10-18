# Lifecycle

circle drives the full SDLC with a TDD approach:

**ideation → planning → implementation → testing → review → merge**

```mermaid
flowchart TB
    Web["Web UI<br/>(htmx / react)"]
    TUI["TUI<br/>(Bubble Tea)"]
    CLI["CLI<br/>(circle)"]

    subgraph Server["circle-server"]
        Intake["Ideation<br/>stories / PRDs"]
        Plan["Planning<br/>story → tasks<br/>acceptance tests written first"]
        Queue["Queue<br/>(SQLite)"]
        Orch["Orchestrator<br/>budgets · retries"]
        Test["Testing · TDD gate<br/>red → green → refactor"]
        Review["Review<br/>diffs · approvals"]
        Merge["Merge<br/>back to main"]
        Intake --> Plan --> Queue --> Orch --> Test --> Review --> Merge
        Test -->|"red — feed errors<br/>retry (cap N)"| Orch
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

## Stages

| Stage          | Input                    | Output                              |
| -------------- | ------------------------ | ----------------------------------- |
| Ideation       | idea / PRD               | `Story` (draft)                     |
| Planning       | `Story`                  | `Task`s + acceptance tests (red)    |
| Implementation | `Task`                   | worktree changes (red → green)      |
| Testing        | worktree                 | green tests + diff, or retry        |
| Review         | diff                     | approval or re-queue                |
| Merge          | approved diff            | change on main                      |

## Test gate

Failing acceptance tests are written at planning time. The gate only releases a
task to review on **green**; a **red** result feeds the errors back to the
orchestrator for a bounded retry (`budgets.max_retries` in
`.circle/config.yaml`).

## State transitions

Story (`internal/domain/story.go`): `draft → triaged → planned → done`, plus
`cancelled`.

Task (`internal/domain/task.go`): `queued → running → review → merged`, with
`failed` and re-queue loops. The TDD gate sits inside `running → review`.
