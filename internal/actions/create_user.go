package actions

import (
	"context"

	"github.com/johnnyfreeman/anvil/internal/core"
)

type CreateUserOpts struct {
	Username string
	Group    *string
}

type CreateUserOptsFunc func(*CreateUserOpts)

func WithGroup(group string) CreateUserOptsFunc {
	return func(o *CreateUserOpts) {
		o.Group = &group
	}
}

type CreateUser struct {
	CreateUserOpts
}

func DefaultCreateUserOpts() CreateUserOpts {
	return CreateUserOpts{
		Username: "",
	}
}

func NewCreateUser(username string, opts ...CreateUserOptsFunc) *CreateUser {
	o := DefaultCreateUserOpts()
	o.Username = username
	for _, fn := range opts {
		fn(&o)
	}
	return &CreateUser{
		CreateUserOpts: o,
	}
}

func (a CreateUser) Handle(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	if observer != nil {
		observer.OnActionStart()
		defer observer.OnActionEnd()
	}

	_, err := ex.Execute(ctx, os.CheckUser(a.Username), observer)

	if err != nil {
		_, err = ex.Execute(ctx, os.CreateUser(a.Username), observer)
		if err != nil {
			return err
		}
	}

	if a.CreateUserOpts.Group != nil {
		_, err = ex.Execute(ctx, os.GroupUser(a.CreateUserOpts.Username, *a.CreateUserOpts.Group), observer)
		if err != nil {
			return err
		}
	}

	return nil
}

var _ core.Action = (*CreateUser)(nil)
