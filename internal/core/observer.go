package core

type ExecutionObserver interface {
	OnExecutionStart(command string) error
	OnExecutionOutput(output string) error
	OnExecutionEnd() error
}
