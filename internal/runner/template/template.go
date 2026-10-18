// Package template is a copy-paste starting point for a new runner adapter.
//
// To add a runner:
//  1. Copy this directory to internal/runner/<name> and rename the package.
//  2. Set Binary/Args in New (or read them from config).
//  3. Translate the agent CLI's output into domain.Event values in run().
//  4. Register the factory in cmd/circle/main.go:
//     reg := runner.NewRegistry()
//     reg.Register("myname", myname.New)
package template

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"os/exec"
	"sync"
	"time"

	"github.com/jalusw/circle/internal/domain"
	"github.com/jalusw/circle/internal/runner"
)

type Adapter struct {
	Binary string
	Args   []string
	Dir    string
}

func New() (runner.Runner, error) {
	return &Adapter{Binary: "myagent", Dir: "."}, nil
}

func (a *Adapter) Name() string { return "template" }

func (a *Adapter) Start(ctx context.Context, spec domain.RunSpec) (runner.Handle, error) {
	ctx, cancel := context.WithCancel(ctx)

	args := append(append([]string{}, a.Args...), spec.Prompt)
	cmd := exec.CommandContext(ctx, a.Binary, args...)
	cmd.Dir = a.Dir

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, err
	}
	if err := cmd.Start(); err != nil {
		cancel()
		return nil, err
	}

	h := &handle{
		ctx:    ctx,
		cancel: cancel,
		cmd:    cmd,
		events: make(chan domain.Event, 64),
		done:   make(chan struct{}),
	}
	go h.run(stdout, spec)
	return h, nil
}

type handle struct {
	ctx    context.Context
	cancel context.CancelFunc
	cmd    *exec.Cmd

	events chan domain.Event
	done   chan struct{}

	mu  sync.Mutex
	res domain.RunResult
	err error
}

func (h *handle) Events() <-chan domain.Event { return h.events }

func (h *handle) Wait(ctx context.Context) (domain.RunResult, error) {
	select {
	case <-h.done:
	case <-ctx.Done():
		_ = h.Kill()
		<-h.done
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.res, h.err
}

func (h *handle) Kill() error {
	h.cancel()
	return nil
}

func (h *handle) run(stdout io.Reader, spec domain.RunSpec) {
	defer close(h.done)
	defer close(h.events)

	h.emit(domain.EventRunStarted, map[string]string{"runner": "template", "title": spec.Title})

	sc := bufio.NewScanner(stdout)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		h.emit(domain.EventAssistantMsg, map[string]string{"line": sc.Text()})
	}

	err := h.cmd.Wait()
	h.mu.Lock()
	h.res = domain.RunResult{Success: err == nil, Summary: "template run"}
	h.err = err
	h.mu.Unlock()

	h.emit(domain.EventRunFinished, map[string]any{"success": err == nil})
}

func (h *handle) emit(kind domain.EventKind, v any) {
	payload, _ := json.Marshal(v)
	select {
	case h.events <- domain.Event{Kind: kind, Payload: payload, TS: time.Now()}:
	case <-h.ctx.Done():
	}
}
