package actions

import (
	"context"

	"github.com/johnnyfreeman/anvil/internal/core"
)

// StartService action for starting a system service
type StartService struct {
	ServiceName string
}

func NewStartService(serviceName string) *StartService {
	return &StartService{
		ServiceName: serviceName,
	}
}

func (a StartService) Handle(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	return core.WithObserver(observer, func() error {
		_, err := ex.Execute(ctx, os.StartService(a.ServiceName), observer)
		return err
	})
}

var _ core.Action = (*StartService)(nil)

// StopService action for stopping a system service
type StopService struct {
	ServiceName string
}

func NewStopService(serviceName string) *StopService {
	return &StopService{
		ServiceName: serviceName,
	}
}

func (a StopService) Handle(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	return core.WithObserver(observer, func() error {
		_, err := ex.Execute(ctx, os.StopService(a.ServiceName), observer)
		return err
	})
}

var _ core.Action = (*StopService)(nil)

// EnableService action for enabling a system service to start on boot
type EnableService struct {
	ServiceName string
}

func NewEnableService(serviceName string) *EnableService {
	return &EnableService{
		ServiceName: serviceName,
	}
}

func (a EnableService) Handle(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	return core.WithObserver(observer, func() error {
		_, err := ex.Execute(ctx, os.EnableService(a.ServiceName), observer)
		return err
	})
}

var _ core.Action = (*EnableService)(nil)

// RestartService action for restarting a system service
type RestartService struct {
	ServiceName string
}

func NewRestartService(serviceName string) *RestartService {
	return &RestartService{
		ServiceName: serviceName,
	}
}

func (a RestartService) Handle(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	return core.WithObserver(observer, func() error {
		_, err := ex.Execute(ctx, os.RestartService(a.ServiceName), observer)
		return err
	})
}

var _ core.Action = (*RestartService)(nil)