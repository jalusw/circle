package gitops

import (
	"fmt"
	"os/exec"
	"strings"
)

func Root(dir string) (string, error) {
	out, err := exec.Command("git", "-C", dir, "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("gitops: not inside a git repo: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}
