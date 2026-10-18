package orchestrator

import (
	"github.com/jalusw/cycle/internal/runner"
)

type Orchestrator struct {
	runners *runner.Registry
}

func New(runners *runner.Registry) *Orchestrator {
	return &Orchestrator{runners: runners}
}
