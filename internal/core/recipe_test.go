package core

import (
	"context"
	"testing"

	"github.com/johnnyfreeman/anvil/internal/testutil"
)

func Test_RecipeRegistry(t *testing.T) {
	registry := NewRecipeRegistry()
	
	// Create a test recipe
	testRecipe := NewBaseRecipe(
		"test-recipe",
		"Test recipe for testing",
		[]Action{},
	)
	
	// Test registration
	registry.Register(testRecipe)
	
	// Test retrieval
	retrieved, exists := registry.Get("test-recipe")
	if !exists {
		t.Fatal("Expected recipe to exist after registration")
	}
	
	if retrieved.Name() != "test-recipe" {
		t.Fatalf("Expected recipe name 'test-recipe', got '%s'", retrieved.Name())
	}
	
	// Test non-existent recipe
	_, exists = registry.Get("non-existent")
	if exists {
		t.Fatal("Expected non-existent recipe to not exist")
	}
	
	// Test list functionality
	recipes := registry.List()
	if len(recipes) != 1 {
		t.Fatalf("Expected 1 recipe in list, got %d", len(recipes))
	}
	
	// Test names functionality
	names := registry.Names()
	if len(names) != 1 || names[0] != "test-recipe" {
		t.Fatalf("Expected names to contain 'test-recipe', got %v", names)
	}
}

func Test_BaseRecipe(t *testing.T) {
	actions := []Action{
		&TestAction{name: "action1"},
		&TestAction{name: "action2"},
	}
	
	recipe := NewBaseRecipe(
		"test-recipe",
		"Test description",
		actions,
	)
	
	// Test getters
	if recipe.Name() != "test-recipe" {
		t.Fatalf("Expected name 'test-recipe', got '%s'", recipe.Name())
	}
	
	if recipe.Description() != "Test description" {
		t.Fatalf("Expected description 'Test description', got '%s'", recipe.Description())
	}
	
	if len(recipe.Actions()) != 2 {
		t.Fatalf("Expected 2 actions, got %d", len(recipe.Actions()))
	}
}

func Test_BaseRecipe_Execute(t *testing.T) {
	testActions := []*TestAction{
		{name: "action1"},
		{name: "action2"},
		{name: "action3"},
	}
	
	actions := make([]Action, len(testActions))
	for i, action := range testActions {
		actions[i] = action
	}
	
	recipe := NewBaseRecipe(
		"test-recipe",
		"Test description",
		actions,
	)
	
	executor := &FakeExecutor{
		Responses: make(map[string]FakeResponse),
	}
	
	os := &testutil.MockOS{}
	observer := &testutil.MockObserver{}
	
	// Execute recipe
	err := recipe.Execute(context.Background(), executor, os, observer)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	
	// Verify all actions were executed
	for _, action := range testActions {
		if !action.executed {
			t.Fatalf("Expected action '%s' to be executed", action.name)
		}
	}
}

// TestAction is a test implementation of Action
type TestAction struct {
	name     string
	executed bool
	err      error
}

func (a *TestAction) Handle(ctx context.Context, ex Executor, os OS, observer ActionObserver) error {
	a.executed = true
	return a.err
}


