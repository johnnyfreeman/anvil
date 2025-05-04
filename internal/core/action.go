package core

import (
	"context"
)

type Action interface {
	Handle(context.Context, Executor, OS, ActionObserver) error
}

type ActionObserver interface {
	ExecutionObserver
	OnActionStart() error
	OnActionEnd() error
}
