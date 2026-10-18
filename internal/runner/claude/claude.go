package claude

import (
	"context"
	"fmt"

	"github.com/jalusw/circle/internal/runner"
	"github.com/jalusw/circle/internal/domain"
)

type Adapter struct{}

func New() (runner.Runner, error)             { return &Adapter{}, nil }
func (a *Adapter) Name() string               { return "claude" }
func (a *Adapter) Start(ctx context.Context, spec domain.RunSpec) (runner.Handle, error) {
	return nil, fmt.Errorf("claude adapter: not implemented")
}
