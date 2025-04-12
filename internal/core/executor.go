package core

import (
	"context"
	"slices"
	"time"
)

type Executor interface {
	Execute(ctx context.Context, command string) (string, error)
	Channel() chan string
}

type SshExecutor struct {
	Server
	C chan string
}

func (e SshExecutor) Execute(ctx context.Context, command string) (string, error) {
	// execute action
	return "", nil
}

func (e SshExecutor) Channel() chan string {
	return e.C
}

type LocalExecutor struct {
	C chan string
}

func (e LocalExecutor) Execute(ctx context.Context, command string) (string, error) {
	// execute action
	return "", nil
}

func (e LocalExecutor) Channel() chan string {
	return e.C
}

type FakeExecutor struct {
	History   []string
	Responses map[string]FakeResponse
	C         chan string
}

type FakeResponse struct {
	Output string
	Err    error
}

func (e *FakeExecutor) Execute(ctx context.Context, command string) (string, error) {
	time.Sleep(1 * time.Second)
	e.History = append(e.History, command)

	if resp, ok := e.Responses[command]; ok {
		if resp.Output != "" {
			e.Channel() <- resp.Output
		}
		if resp.Err != nil {
			e.Channel() <- resp.Err.Error()
		}
		return resp.Output, resp.Err
	}

	e.Channel() <- command
	return "", nil
}

func (e FakeExecutor) Channel() chan string {
	return e.C
}

func (e *FakeExecutor) Executed(command string) bool {
	return slices.Contains(e.History, command)
}
