package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/johnnyfreeman/anvil/internal/cli"
	"github.com/johnnyfreeman/anvil/internal/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: anvil [--dry-run] <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  create-user [--group <group>] <username>")
		fmt.Println("  install-package [--update] <package>")
		fmt.Println("  recipe <recipe-name>")
		fmt.Println("  recipe --list")
		fmt.Println("  detect-os")
		fmt.Println("")
		fmt.Println("Global flags:")
		fmt.Println("  --dry-run    Show what would be executed without running commands")
		os.Exit(1)
	}

	ctx := context.Background()
	
	// Parse global flags
	args := os.Args[1:]
	var executor core.Executor = &core.LocalExecutor{}
	dryRun := false
	
	// Check for --dry-run flag
	if len(args) > 0 && args[0] == "--dry-run" {
		dryRun = true
		executor = &core.DryRunExecutor{Commands: make([]string, 0)}
		args = args[1:]
		if len(args) == 0 {
			log.Fatal("Command required after --dry-run")
		}
	}
	
	if dryRun {
		fmt.Println("🔍 DRY RUN MODE - No commands will be executed")
		fmt.Println("")
	}

	switch args[0] {
	case "create-user":
		cli.CreateUserCommand(ctx, executor, args[1:])
	case "install-package":
		cli.InstallPackageCommand(ctx, executor, args[1:])
	case "recipe":
		cli.RecipeCommand(ctx, executor, args[1:])
	case "detect-os":
		cli.DetectOSCommand(ctx, executor)
	default:
		log.Fatalf("Unknown command: %s", args[0])
	}
}

