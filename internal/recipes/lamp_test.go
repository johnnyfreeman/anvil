package recipes

import (
	"context"
	"strings"
	"testing"

	"github.com/johnnyfreeman/anvil/internal/core"
	"github.com/johnnyfreeman/anvil/internal/testutil"
)

func Test_LAMPServer_Recipe(t *testing.T) {
	recipe := NewLAMPServer()
	
	// Test basic properties
	if recipe.Name() != "lamp-server" {
		t.Fatalf("Expected name 'lamp-server', got '%s'", recipe.Name())
	}
	
	if !strings.Contains(recipe.Description(), "LAMP") {
		t.Fatalf("Expected description to contain 'LAMP', got '%s'", recipe.Description())
	}
	
	// Test that recipe has multiple actions
	actions := recipe.Actions()
	if len(actions) == 0 {
		t.Fatal("Expected LAMP recipe to have actions")
	}
	
	// We expect actions for: update, apache2, mysql, php packages, services
	// At minimum: update + apache2 + mysql + php + several php modules + service actions
	if len(actions) < 10 {
		t.Fatalf("Expected at least 10 actions for LAMP setup, got %d", len(actions))
	}
}

func Test_LAMPServer_Execute(t *testing.T) {
	recipe := NewLAMPServer()
	
	executor := &core.FakeExecutor{
		Responses: make(map[string]core.FakeResponse),
	}
	
	os := &testutil.MockOS{}
	observer := &testutil.MockObserver{}
	
	// Execute recipe
	err := recipe.Execute(context.Background(), executor, os, observer)
	if err != nil {
		t.Fatalf("Expected no error executing LAMP recipe, got %v", err)
	}
	
	// Verify that commands were executed
	if len(executor.History) == 0 {
		t.Fatal("Expected commands to be executed")
	}
	
	// Check for key LAMP components in the executed commands
	hasApache := false
	hasMySQL := false
	hasPHP := false
	
	for _, cmd := range executor.History {
		if strings.Contains(cmd, "apache2") {
			hasApache = true
		}
		if strings.Contains(cmd, "mysql") {
			hasMySQL = true
		}
		if strings.Contains(cmd, "php") {
			hasPHP = true
		}
	}
	
	if !hasApache {
		t.Error("Expected Apache installation commands")
	}
	if !hasMySQL {
		t.Error("Expected MySQL installation commands")
	}
	if !hasPHP {
		t.Error("Expected PHP installation commands")
	}
}

func Test_BasicWebServer_Recipe(t *testing.T) {
	recipe := NewBasicWebServer()
	
	// Test basic properties
	if recipe.Name() != "webserver" {
		t.Fatalf("Expected name 'webserver', got '%s'", recipe.Name())
	}
	
	actions := recipe.Actions()
	if len(actions) == 0 {
		t.Fatal("Expected web server recipe to have actions")
	}
}

func Test_NginxWebServer_Recipe(t *testing.T) {
	recipe := NewNginxWebServer()
	
	// Test basic properties
	if recipe.Name() != "nginx-webserver" {
		t.Fatalf("Expected name 'nginx-webserver', got '%s'", recipe.Name())
	}
	
	actions := recipe.Actions()
	if len(actions) == 0 {
		t.Fatal("Expected Nginx recipe to have actions")
	}
}

func Test_DefaultRegistry(t *testing.T) {
	registry := DefaultRegistry()
	
	// Test that all expected recipes are registered
	expectedRecipes := []string{
		"lamp-server",
		"webserver", 
		"nginx-webserver",
	}
	
	for _, recipeName := range expectedRecipes {
		recipe, exists := registry.Get(recipeName)
		if !exists {
			t.Fatalf("Expected recipe '%s' to be registered", recipeName)
		}
		
		if recipe.Name() != recipeName {
			t.Fatalf("Expected recipe name '%s', got '%s'", recipeName, recipe.Name())
		}
	}
	
	// Test that we have the expected number of recipes
	recipes := registry.List()
	if len(recipes) != len(expectedRecipes) {
		t.Fatalf("Expected %d recipes, got %d", len(expectedRecipes), len(recipes))
	}
}

