package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/johnnyfreeman/anvil/internal/core"
)

type Runner struct {
	ctx      context.Context
	executor core.Executor
	observer core.ActionObserver
}

func NewRunner(ctx context.Context, executor core.Executor, observer core.ActionObserver) *Runner {
	return &Runner{
		ctx:      ctx,
		executor: executor,
		observer: observer,
	}
}

func (r *Runner) ExecuteAction(action core.Action, successMsg string) {
	osInfo, err := core.DetectOS(r.ctx, r.executor)
	if err != nil {
		log.Fatalf("Failed to detect OS: %v", err)
	}

	if err := action.Handle(r.ctx, r.executor, osInfo.Detected, r.observer); err != nil {
		log.Fatalf("Action failed: %v", err)
	}

	fmt.Println(successMsg)
}