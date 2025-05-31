package recipes

import (
	"github.com/johnnyfreeman/anvil/internal/core"
)

// DefaultRegistry returns a registry with all built-in recipes
func DefaultRegistry() *core.RecipeRegistry {
	registry := core.NewRecipeRegistry()
	
	// Register all available recipes
	registry.Register(NewLAMPServer())
	registry.Register(NewBasicWebServer())
	registry.Register(NewNginxWebServer())
	
	return registry
}