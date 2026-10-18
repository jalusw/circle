package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jalusw/cycle/internal/runner"
	_ "github.com/jalusw/cycle/internal/runner/claude"
	_ "github.com/jalusw/cycle/internal/runner/opencode"
	_ "github.com/jalusw/cycle/internal/runner/rawllm"
)

func main() {
	ctx := context.Background()
	_ = ctx
	_ = runner.NewRegistry()
	fmt.Fprintln(os.Stderr, "cycle: not implemented yet")
}
