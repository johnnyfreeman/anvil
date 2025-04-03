package main

import (
	"context"
	"slices"
)

type Executor interface {
	Execute(ctx context.Context, command string) (string, error)
}

type SshExecutor struct{}

func (e SshExecutor) Execute(ctx context.Context, command string) (string, error) {
	// execute action
	return "", nil
}

type LocalExecutor struct{}

func (e LocalExecutor) Execute(ctx context.Context, command string) (string, error) {
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

func (e *FakeExecutor) Execute(ctx context.Context, command string) (string, error) {
	e.History = append(e.History, command)

	if resp, ok := e.Responses[command]; ok {
		return resp.Output, resp.Err
	}

	return "", nil
}

func (e *FakeExecutor) Executed(command string) bool {
	return slices.Contains(e.History, command)
}
