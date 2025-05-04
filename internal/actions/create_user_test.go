package actions

import (
	"errors"
	"testing"

	"github.com/johnnyfreeman/anvil/internal/core"
)

func Test_CreateUser(t *testing.T) {
	username := "john"
	os := core.Ubuntu{}
	ex := &core.FakeExecutor{
		Responses: map[string]core.FakeResponse{
			os.CheckUser(username): {Err: errors.New("no such user")},
		},
	}
	action := NewCreateUser(username)

	if err := action.Handle(t.Context(), ex, os, nil); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !ex.Executed(os.CheckUser(username)) {
		t.Error("did not execute command to check if user exists")
	}

	if !ex.Executed(os.CreateUser(username)) {
		t.Error("did not execute command to create user")
	}
}

func Test_CreateUserWithGroup(t *testing.T) {
	username := "john"
	os := core.Ubuntu{}
	ex := &core.FakeExecutor{
		Responses: map[string]core.FakeResponse{
			os.CheckUser(username): {Err: errors.New("no such user")},
		},
	}
	group := "audio"
	action := NewCreateUser(username, WithGroup(group))

	if err := action.Handle(t.Context(), ex, os, nil); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !ex.Executed(os.CheckUser(username)) {
		t.Error("did not execute command to check if user exists")
	}

	if !ex.Executed(os.CreateUser(username)) {
		t.Error("did not execute command to create user")
	}

	if !ex.Executed(os.GroupUser(username, group)) {
		t.Error("did not execute command to group user")
	}
}

func Test_CreateUser_UserExists(t *testing.T) {
	username := "john"
	os := core.Ubuntu{}
	ex := &core.FakeExecutor{
		Responses: map[string]core.FakeResponse{
			os.CheckUser(username): {Output: "1002"},
		},
	}
	action := NewCreateUser(username)

	// When Handle is called, it should see that the user exists and not try to create it.
	if err := action.Handle(t.Context(), ex, os, nil); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// It should always check if the user exists.
	if !ex.Executed(os.CheckUser(username)) {
		t.Error("did not execute command to check if user exists")
	}

	// If the user exists, the command to create the user should not have been executed.
	if ex.Executed(os.CreateUser(username)) {
		t.Error("should not execute command to create user when user already exists")
	}
}

func Test_CreateUser_StreamOutput(t *testing.T) {
	username := "john"
	os := core.Ubuntu{}
	ex := &core.FakeExecutor{
		Responses: map[string]core.FakeResponse{
			os.CheckUser(username): {Output: "1002"},
		},
	}
	action := NewCreateUser(username)

	// When Handle is called, it should see that the user exists and not try to create it.
	if err := action.Handle(t.Context(), ex, os, nil); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// It should always check if the user exists.
	if !ex.Executed(os.CheckUser(username)) {
		t.Error("did not execute command to check if user exists")
	}

	// If the user exists, the command to create the user should not have been executed.
	if ex.Executed(os.CreateUser(username)) {
		t.Error("should not execute command to create user when user already exists")
	}
}
