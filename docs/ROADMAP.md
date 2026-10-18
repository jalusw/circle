# Roadmap

## v0 — vertical slice
- [ ] store: story/task/run CRUD + migrations
- [ ] api: POST /stories, GET /events (SSE)
- [ ] runner: one real adapter (opencode serve) driving a task in a worktree
- [ ] cli: `circle submit`, `circle watch`
- [ ] gitops: create worktree, extract diff, merge locally

## v1 — solid lifecycle
- [ ] budgets: $, tokens, time, retries; kill-switch
- [ ] failure loop: CI red → feed errors → retry (cap N)
- [ ] approval gates + diff review UI
- [ ] circle-tui: live event streams
- [ ] web ui parity

## v2 — scale
- [ ] containers instead of raw worktrees
- [ ] GitHub PR automation
- [ ] remote server (multi-user)
- [ ] planner promotion: rawllm planner → planner runner
