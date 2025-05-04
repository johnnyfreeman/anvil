package core

import (
	"context"
	"slices"
)

type Executor interface {
	Execute(ctx context.Context, command string, observer ExecutionObserver) (string, error)
}

type SshExecutor struct{}

func (e SshExecutor) Execute(ctx context.Context, command string, observer ExecutionObserver) (string, error) {
	// execute action
	return "", nil
}

type LocalExecutor struct{}

func (e LocalExecutor) Execute(ctx context.Context, command string, observer ExecutionObserver) (string, error) {
	// execute action
	return "", nil
}

type FakeExecutor struct {
	History   []string
	Responses map[string]FakeResponse
}

type FakeResponse struct {
	Output string
	Err    error
}

func (e *FakeExecutor) Execute(ctx context.Context, command string, observer ExecutionObserver) (string, error) {
	if observer != nil {
		observer.OnExecutionStart(command)
		defer observer.OnExecutionEnd()
	}

	e.History = append(e.History, command)

	if resp, ok := e.Responses[command]; ok {
		if observer != nil {
			observer.OnExecutionOutput(resp.Output)
		}
		return resp.Output, resp.Err
	}

	if observer != nil {
		observer.OnExecutionOutput("")
	}
	return "", nil
}

func (e *FakeExecutor) Executed(command string) bool {
	return slices.Contains(e.History, command)
}
