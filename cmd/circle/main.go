package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jalusw/circle/internal/runner"
	_ "github.com/jalusw/circle/internal/runner/claude"
	_ "github.com/jalusw/circle/internal/runner/opencode"
	_ "github.com/jalusw/circle/internal/runner/rawllm"
)

func main() {
	ctx := context.Background()
	_ = ctx
	_ = runner.NewRegistry()
	fmt.Fprintln(os.Stderr, "circle: not implemented yet")
}
