package cli

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/johnnyfreeman/anvil/internal/actions"
	"github.com/johnnyfreeman/anvil/internal/core"
	"github.com/johnnyfreeman/anvil/internal/recipes"
)

func CreateUserCommand(ctx context.Context, executor core.Executor, args []string) {
	fs := flag.NewFlagSet("create-user", flag.ExitOnError)
	group := fs.String("group", "", "Optional group to add user to")
	
	if err := fs.Parse(args); err != nil {
		log.Fatal(err)
	}
	
	if fs.NArg() < 1 {
		log.Fatal("Username required")
	}
	
	username := fs.Arg(0)
	
	// Create action
	var opts []actions.CreateUserOptsFunc
	if *group != "" {
		opts = append(opts, actions.WithGroup(*group))
	}
	
	action := actions.NewCreateUser(username, opts...)
	
	// Execute action
	runner := NewRunner(ctx, executor, &cliObserver{})
	runner.ExecuteAction(action, fmt.Sprintf("✓ User %s created successfully", username))
}

func InstallPackageCommand(ctx context.Context, executor core.Executor, args []string) {
	fs := flag.NewFlagSet("install-package", flag.ExitOnError)
	update := fs.Bool("update", false, "Update package lists before installing")
	
	if err := fs.Parse(args); err != nil {
		log.Fatal(err)
	}
	
	if fs.NArg() < 1 {
		log.Fatal("Package name required")
	}
	
	packageName := fs.Arg(0)
	
	// Create action
	var opts []actions.InstallPackageOptsFunc
	if *update {
		opts = append(opts, actions.WithUpdate())
	}
	
	action := actions.NewInstallPackage(packageName, opts...)
	
	// Execute action
	runner := NewRunner(ctx, executor, &cliObserver{})
	runner.ExecuteAction(action, fmt.Sprintf("✓ Package %s installed successfully", packageName))
}

func RecipeCommand(ctx context.Context, executor core.Executor, args []string) {
	registry := recipes.DefaultRegistry()
	
	if len(args) == 0 {
		log.Fatal("Recipe name required or --list flag")
	}
	
	// Handle --list flag
	if args[0] == "--list" {
		fmt.Println("Available recipes:")
		for _, recipe := range registry.List() {
			fmt.Printf("  %-20s %s\n", recipe.Name(), recipe.Description())
		}
		return
	}
	
	recipeName := args[0]
	
	// Check if recipe exists
	recipe, exists := registry.Get(recipeName)
	if !exists {
		fmt.Printf("Recipe '%s' not found.\n\n", recipeName)
		fmt.Println("Available recipes:")
		for _, r := range registry.List() {
			fmt.Printf("  %-20s %s\n", r.Name(), r.Description())
		}
		os.Exit(1)
	}
	
	// Execute recipe
	runner := NewRunner(ctx, executor, &cliObserver{})
	action := actions.NewExecuteRecipe(recipeName, registry)
	
	fmt.Printf("🚀 Executing recipe: %s\n", recipe.Description())
	fmt.Printf("📋 Actions: %d\n\n", len(recipe.Actions()))
	
	runner.ExecuteAction(action, fmt.Sprintf("✓ Recipe '%s' completed successfully", recipeName))
}

func DetectOSCommand(ctx context.Context, executor core.Executor) {
	osInfo, err := core.DetectOS(ctx, executor)
	if err != nil {
		log.Fatalf("Failed to detect OS: %v", err)
	}
	
	fmt.Printf("OS: %s\n", osInfo.Pretty)
	fmt.Printf("ID: %s\n", osInfo.ID)
	fmt.Printf("Version: %s\n", osInfo.Version)
	fmt.Printf("Detected Type: %T\n", osInfo.Detected)
}

type cliObserver struct{}

func (o *cliObserver) OnActionStart() error {
	fmt.Println("→ Starting action...")
	return nil
}

func (o *cliObserver) OnActionEnd() error {
	return nil
}

func (o *cliObserver) OnExecutionStart(command string) error {
	fmt.Printf("  $ %s\n", command)
	return nil
}

func (o *cliObserver) OnExecutionOutput(output string) error {
	fmt.Print("  " + output)
	return nil
}

func (o *cliObserver) OnExecutionEnd() error {
	return nil
}