package actions

import (
	"context"
	"fmt"

	"github.com/johnnyfreeman/anvil/internal/core"
)

// ExecuteRecipe action for running a recipe by name
type ExecuteRecipe struct {
	RecipeName string
	Registry   *core.RecipeRegistry
}

func NewExecuteRecipe(recipeName string, registry *core.RecipeRegistry) *ExecuteRecipe {
	return &ExecuteRecipe{
		RecipeName: recipeName,
		Registry:   registry,
	}
}

func (a ExecuteRecipe) Handle(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	return core.WithObserver(observer, func() error {
		recipe, exists := a.Registry.Get(a.RecipeName)
		if !exists {
			return fmt.Errorf("recipe '%s' not found", a.RecipeName)
		}
		
		return recipe.Execute(ctx, ex, os, observer)
	})
}

var _ core.Action = (*ExecuteRecipe)(nil)