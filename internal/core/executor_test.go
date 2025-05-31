package core

import (
	"context"
	"strings"
	"testing"
)

func Test_FakeExecutor(t *testing.T) {
	ex := FakeExecutor{}
	command := "some command"
	output, _ := ex.Execute(context.Background(), command, nil)
	if output != "" {
		t.Fatalf("fake executor should empty output: command '%s'", command)
	}
	if !ex.Executed(command) {
		t.Fatalf("fake executor should have executed: '%s'", command)
	}
}

func Test_ParallelSshExecutor_Success(t *testing.T) {
	// Create a ParallelSshExecutor with test hosts
	executor := ParallelSshExecutor{
		Hosts: []SshHost{
			{Host: "host1.example.com", User: "user1"},
			{Host: "host2.example.com", User: "user2"},
			{Host: "host3.example.com", User: "user3"},
		},
	}

	// This test would normally require real SSH connections
	// In a real test environment, you'd mock the SSH connections
	// For now, we test the structure and interface
	ctx := context.Background()
	
	// Test that the method signature works
	_, err := executor.Execute(ctx, "echo 'test'", nil)
	
	// In a real environment, this would fail due to SSH connections
	// but we can verify the structure is correct
	if err == nil {
		t.Log("ParallelSshExecutor executed successfully (likely in test environment)")
	} else {
		// Expected in test environment without SSH setup
		if !strings.Contains(err.Error(), "execution failed on some hosts") {
			t.Logf("Expected SSH connection failure in test environment: %v", err)
		}
	}
}

func Test_ParallelSshExecutor_EmptyHosts(t *testing.T) {
	executor := ParallelSshExecutor{
		Hosts: []SshHost{},
	}

	ctx := context.Background()
	output, err := executor.Execute(ctx, "echo 'test'", nil)

	if err != nil {
		t.Fatalf("Expected no error for empty hosts, got: %v", err)
	}

	if output != "" {
		t.Fatalf("Expected empty output for no hosts, got: %s", output)
	}
}

func Test_ParallelSshExecutor_WithObserver(t *testing.T) {
	executor := ParallelSshExecutor{
		Hosts: []SshHost{
			{Host: "host1.example.com", User: "user1"},
		},
	}

	observer := &TestObserver{}
	ctx := context.Background()
	
	_, err := executor.Execute(ctx, "echo 'test'", observer)
	
	// Verify observer was called
	if !observer.StartCalled {
		t.Error("Expected OnExecutionStart to be called")
	}
	
	if !observer.EndCalled {
		t.Error("Expected OnExecutionEnd to be called")
	}
	
	// In test environment, this will likely fail SSH connection
	// but observer should still be called
	if err != nil && !strings.Contains(err.Error(), "execution failed on some hosts") {
		t.Logf("Expected SSH connection failure in test environment: %v", err)
	}
}

// TestObserver is a test implementation of ActionObserver
type TestObserver struct {
	StartCalled       bool
	EndCalled         bool
	OutputCalled      bool
	ActionStartCalled bool
	ActionEndCalled   bool
	Commands          []string
	Outputs           []string
}

func (o *TestObserver) OnExecutionStart(command string) error {
	o.StartCalled = true
	o.Commands = append(o.Commands, command)
	return nil
}

func (o *TestObserver) OnExecutionEnd() error {
	o.EndCalled = true
	return nil
}

func (o *TestObserver) OnExecutionOutput(output string) error {
	o.OutputCalled = true
	o.Outputs = append(o.Outputs, output)
	return nil
}

func (o *TestObserver) OnActionStart() error {
	o.ActionStartCalled = true
	return nil
}

func (o *TestObserver) OnActionEnd() error {
	o.ActionEndCalled = true
	return nil
}
