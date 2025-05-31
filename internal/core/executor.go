package core

import (
	"context"
	"io"
	"os/exec"
	"slices"
	"strings"
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
	if observer != nil {
		if err := observer.OnExecutionStart(command); err != nil {
			return "", err
		}
		defer func() {
			if err := observer.OnExecutionEnd(); err != nil {
				// Log error but don't fail execution
			}
		}()
	}

	cmd := exec.CommandContext(ctx, "sh", "-c", command)
	
	var output strings.Builder
	cmd.Stdout = &output
	cmd.Stderr = &output

	if observer != nil {
		cmd.Stdout = io.MultiWriter(&output, &observerWriter{observer: observer})
		cmd.Stderr = cmd.Stdout
	}

	err := cmd.Run()
	return output.String(), err
}

type observerWriter struct {
	observer ExecutionObserver
}

func (w *observerWriter) Write(p []byte) (n int, err error) {
	if err := w.observer.OnExecutionOutput(string(p)); err != nil {
		// Log error but continue writing
		return len(p), nil
	}
	return len(p), nil
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
		if err := observer.OnExecutionStart(command); err != nil {
			return "", err
		}
		defer func() {
			if err := observer.OnExecutionEnd(); err != nil {
				// Log error but don't fail execution
			}
		}()
	}

	e.History = append(e.History, command)

	if resp, ok := e.Responses[command]; ok {
		if observer != nil {
			if err := observer.OnExecutionOutput(resp.Output); err != nil {
				// Log error but continue
			}
		}
		return resp.Output, resp.Err
	}

	if observer != nil {
		if err := observer.OnExecutionOutput(""); err != nil {
			// Log error but continue
		}
	}
	return "", nil
}

func (e *FakeExecutor) Executed(command string) bool {
	return slices.Contains(e.History, command)
}
