# Roadmap

The lifecycle is TDD-driven: **ideation → planning → implementation → testing
→ review → merge**, with failing tests written before implementation. See
[LIFECYCLE.md](LIFECYCLE.md) for the full cycle and state transitions.

## v0 — vertical slice
- [ ] store: story/task/run CRUD + migrations
- [ ] api: POST /stories, GET /events (SSE)
- [ ] runner: one real adapter (opencode serve) driving a task in a worktree
- [ ] cli: `circle submit`, `circle watch`
- [ ] gitops: create worktree, extract diff, merge locally

## v1 — solid lifecycle
- [ ] planning: story → tasks with acceptance tests (written first)
- [ ] TDD gate: red → green → refactor before review
- [ ] budgets: $, tokens, time, retries; kill-switch
- [ ] failure loop: red tests → feed errors → retry (cap N)
- [ ] approval gates + diff review UI
- [ ] circle-tui: live event streams
- [ ] web ui parity

## v2 — scale
- [ ] credentials: remote origin + SSH key / token handling (see [CREDENTIALS.md](CREDENTIALS.md))
- [ ] containers instead of raw worktrees
- [ ] GitHub PR automation
- [ ] remote server (multi-user)
- [ ] planner promotion: rawllm planner → planner runner
