package core

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"slices"
	"strings"
	"sync"
)

type Executor interface {
	Execute(ctx context.Context, command string, observer ExecutionObserver) (string, error)
}

type SshExecutor struct {
	Host string
	User string
}

func (e SshExecutor) Execute(ctx context.Context, command string, observer ExecutionObserver) (string, error) {
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

	// Use ssh command for now - can be improved with golang.org/x/crypto/ssh later
	sshCmd := fmt.Sprintf("ssh %s@%s '%s'", e.User, e.Host, command)
	cmd := exec.CommandContext(ctx, "sh", "-c", sshCmd)
	
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

type DryRunExecutor struct {
	Commands       []string
	simulatedError bool
}

func (e *DryRunExecutor) Execute(ctx context.Context, command string, observer ExecutionObserver) (string, error) {
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

	e.Commands = append(e.Commands, command)

	if observer != nil {
		if err := observer.OnExecutionOutput("[DRY RUN] Command would be executed\n"); err != nil {
			// Log error but continue
		}
	}

	// Simulate realistic command behavior for dry run
	// This helps actions make proper decisions about next steps
	if strings.Contains(command, "id -u") && !e.simulatedError {
		// Simulate user not found to trigger user creation (only once)
		e.simulatedError = true
		return "", fmt.Errorf("exit status 1")
	}

	return "[DRY RUN] Command would be executed", nil
}

// ParallelSshExecutor executes commands on multiple SSH hosts concurrently
type ParallelSshExecutor struct {
	Hosts []SshHost
}

type SshHost struct {
	Host string
	User string
}

type ParallelResult struct {
	Host   string
	Output string
	Error  error
}

func (e ParallelSshExecutor) Execute(ctx context.Context, command string, observer ExecutionObserver) (string, error) {
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

	results := make(chan ParallelResult, len(e.Hosts))
	var wg sync.WaitGroup

	// Execute command on all hosts concurrently
	for _, host := range e.Hosts {
		wg.Add(1)
		go func(h SshHost) {
			defer wg.Done()
			
			executor := SshExecutor{Host: h.Host, User: h.User}
			output, err := executor.Execute(ctx, command, nil) // Don't pass observer to avoid duplicate notifications
			
			results <- ParallelResult{
				Host:   h.Host,
				Output: output,
				Error:  err,
			}
		}(host)
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	var allOutputs []string
	var errors []string
	
	// Collect results
	for result := range results {
		hostOutput := fmt.Sprintf("[%s] %s", result.Host, result.Output)
		allOutputs = append(allOutputs, hostOutput)
		
		if observer != nil {
			if err := observer.OnExecutionOutput(hostOutput); err != nil {
				// Log error but continue
			}
		}
		
		if result.Error != nil {
			errors = append(errors, fmt.Sprintf("[%s] %v", result.Host, result.Error))
		}
	}

	// Combine all outputs
	combinedOutput := strings.Join(allOutputs, "\n")
	
	// If any host had errors, return combined error
	if len(errors) > 0 {
		return combinedOutput, fmt.Errorf("execution failed on some hosts: %s", strings.Join(errors, "; "))
	}

	return combinedOutput, nil
}
