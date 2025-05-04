package core

import (
	"context"
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
