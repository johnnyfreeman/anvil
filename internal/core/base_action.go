package core

import "context"

func WithObserver(observer ActionObserver, fn func() error) error {
	if observer == nil {
		return fn()
	}

	if err := observer.OnActionStart(); err != nil {
		return err
	}
	defer func() {
		if err := observer.OnActionEnd(); err != nil {
			// Log error but don't fail action
		}
	}()

	return fn()
}

func ExecuteAction(ctx context.Context, ex Executor, os OS, observer ActionObserver, fn func(context.Context, Executor, OS, ActionObserver) error) error {
	return WithObserver(observer, func() error {
		return fn(ctx, ex, os, observer)
	})
}