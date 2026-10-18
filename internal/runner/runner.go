package runner

import (
	"context"

	"github.com/jalusw/circle/internal/domain"
)

type Runner interface {
	Name() string
	Start(ctx context.Context, spec domain.RunSpec) (Handle, error)
}

type Handle interface {
	Events() <-chan domain.Event
	Wait(ctx context.Context) (domain.RunResult, error)
	Kill() error
}

type Factory func() (Runner, error)

type Registry struct {
	factories map[string]Factory
}

func NewRegistry() *Registry {
	return &Registry{factories: map[string]Factory{}}
}

func (r *Registry) Register(name string, f Factory) {
	r.factories[name] = f
}

func (r *Registry) Get(name string) (Runner, error) {
	f, ok := r.factories[name]
	if !ok {
		return nil, ErrUnknownRunner{name}
	}
	return f()
}

func (r *Registry) Names() []string {
	names := make([]string, 0, len(r.factories))
	for name := range r.factories {
		names = append(names, name)
	}
	return names
}

type ErrUnknownRunner struct {
	Name string
}

func (e ErrUnknownRunner) Error() string {
	return "unknown runner: " + e.Name
}
