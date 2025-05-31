package core

import (
	"context"
)

// Recipe defines a reusable set of actions for common server configurations
type Recipe interface {
	Name() string
	Description() string
	Actions() []Action
	Execute(context.Context, Executor, OS, ActionObserver) error
}

// BaseRecipe provides common functionality for all recipes
type BaseRecipe struct {
	name        string
	description string
	actions     []Action
}

func NewBaseRecipe(name, description string, actions []Action) BaseRecipe {
	return BaseRecipe{
		name:        name,
		description: description,
		actions:     actions,
	}
}

func (r BaseRecipe) Name() string {
	return r.name
}

func (r BaseRecipe) Description() string {
	return r.description
}

func (r BaseRecipe) Actions() []Action {
	return r.actions
}

func (r BaseRecipe) Execute(ctx context.Context, ex Executor, os OS, observer ActionObserver) error {
	for _, action := range r.actions {
		if err := action.Handle(ctx, ex, os, observer); err != nil {
			return err
		}
	}
	return nil
}

// RecipeRegistry manages available recipes
type RecipeRegistry struct {
	recipes map[string]Recipe
}

func NewRecipeRegistry() *RecipeRegistry {
	return &RecipeRegistry{
		recipes: make(map[string]Recipe),
	}
}

func (r *RecipeRegistry) Register(recipe Recipe) {
	r.recipes[recipe.Name()] = recipe
}

func (r *RecipeRegistry) Get(name string) (Recipe, bool) {
	recipe, exists := r.recipes[name]
	return recipe, exists
}

func (r *RecipeRegistry) List() []Recipe {
	recipes := make([]Recipe, 0, len(r.recipes))
	for _, recipe := range r.recipes {
		recipes = append(recipes, recipe)
	}
	return recipes
}

func (r *RecipeRegistry) Names() []string {
	names := make([]string, 0, len(r.recipes))
	for name := range r.recipes {
		names = append(names, name)
	}
	return names
}