package main

import (
	"context"
)

type Action interface {
	Handle(ctx context.Context, executor Executor) error
}

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

type CreateUserAction struct {
	CreateUserOpts
}

func DefaultCreateUserOpts() CreateUserOpts {
	return CreateUserOpts{
		Username: "",
	}
}

func NewCreateUserAction(username string, opts ...CreateUserOptsFunc) *CreateUserAction {
	o := DefaultCreateUserOpts()
	o.Username = username
	for _, fn := range opts {
		fn(&o)
	}
	return &CreateUserAction{
		CreateUserOpts: o,
	}
}

func (a CreateUserAction) Handle(ctx context.Context, ex Executor, os OS) error {
	_, err := ex.Execute(ctx, os.CheckUser(a.Username))

	if err != nil {
		_, err = ex.Execute(ctx, os.CreateUser(a.Username))
		if err != nil {
			return err
		}
	}

	if a.CreateUserOpts.Group != nil {
		_, err = ex.Execute(ctx, os.GroupUser(a.CreateUserOpts.Username, *a.CreateUserOpts.Group))
		if err != nil {
			return err
		}
	}

	return nil
}
