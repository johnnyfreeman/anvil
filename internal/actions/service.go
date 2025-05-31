package actions

import (
	"context"

	"github.com/johnnyfreeman/anvil/internal/core"
)

// ServiceOperation represents the type of service operation to perform
type ServiceOperation int

const (
	StartService ServiceOperation = iota
	StopService
	EnableService
	RestartService
)

// ServiceAction performs operations on system services
type ServiceAction struct {
	ServiceName string
	Operation   ServiceOperation
}

func NewStartService(serviceName string) *ServiceAction {
	return &ServiceAction{
		ServiceName: serviceName,
		Operation:   StartService,
	}
}

func NewStopService(serviceName string) *ServiceAction {
	return &ServiceAction{
		ServiceName: serviceName,
		Operation:   StopService,
	}
}

func NewEnableService(serviceName string) *ServiceAction {
	return &ServiceAction{
		ServiceName: serviceName,
		Operation:   EnableService,
	}
}

func NewRestartService(serviceName string) *ServiceAction {
	return &ServiceAction{
		ServiceName: serviceName,
		Operation:   RestartService,
	}
}

func (a ServiceAction) Handle(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	return core.WithObserver(observer, func() error {
		var command string
		switch a.Operation {
		case StartService:
			command = os.StartService(a.ServiceName)
		case StopService:
			command = os.StopService(a.ServiceName)
		case EnableService:
			command = os.EnableService(a.ServiceName)
		case RestartService:
			command = os.RestartService(a.ServiceName)
		}
		_, err := ex.Execute(ctx, command, observer)
		return err
	})
}

var _ core.Action = (*ServiceAction)(nil)