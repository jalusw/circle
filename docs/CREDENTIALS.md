# Credentials

circle needs to push task branches back to a remote origin. Secrets should be
referenced, not copied, wherever possible.

## Option A — delegate to the host (recommended, v0)

circle stores only the remote URL. The SSH key lives in `~/.ssh/` — ideally a
**deploy key scoped to a single repository** — and is never stored by circle.

```yaml
# .circle/config.yaml
repo:
  remote: git@github.com:jalusw/circle.git
  ssh_key: ~/.ssh/id_ed25519_circle   # reference only, not copied
```

gitops sets a per-run command instead of reading a stored secret:

```
GIT_SSH_COMMAND="ssh -i <path> -o IdentitiesOnly=yes" git push ...
```

Works with ssh-agent / `~/.ssh/config` too. For container mode, mount the key
read-only into the container.

## Option B — store encrypted at rest

If keys must be stored (e.g. multi-repo, container mode), encrypt them with
[`filippo.io/age`](https://pkg.go.dev/filippo.io/age) (pure Go, X25519) before
writing to SQLite:

```sql
CREATE TABLE IF NOT EXISTS credentials (
    id         TEXT PRIMARY KEY,
    name       TEXT NOT NULL,        -- e.g. "github-oss"
    kind       TEXT NOT NULL,        -- ssh_key | https_token
    remote     TEXT,                 -- optional: scoped to an origin URL
    secret     TEXT NOT NULL,        -- age-encrypted PEM / token
    created_at DATETIME NOT NULL
);
```

- Master key from `CIRCLE_MASTER_KEY` (env) or `~/.circle/master.key`
  (`0600`, gitignored)
- On push, decrypt to a temp key file (`0600`, deleted after use) and pass it
  via `GIT_SSH_COMMAND`
- For HTTPS tokens, use `git credential.helper` or
  `http.extraHeader=Authorization: Basic ...`

## Rules

- File permissions `0600`; never log secrets
- Redact secrets (and tokens embedded in URLs) from API responses
- Never commit key material; verify with `git check-ignore`
- Prefer a fine-grained PAT over a long-lived SSH key when you need easy
  revocation
