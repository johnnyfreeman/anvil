package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/johnnyfreeman/anvil/internal/actions"
	"github.com/johnnyfreeman/anvil/internal/cli"
	"github.com/johnnyfreeman/anvil/internal/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: anvil [--dry-run] <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  create-user [--group <group>] <username>")
		fmt.Println("  install-package [--update] <package>")
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
		createUserCmd(ctx, executor, args[1:])
	case "install-package":
		installPackageCmd(ctx, executor, args[1:])
	case "detect-os":
		detectOSCmd(ctx, executor)
	default:
		log.Fatalf("Unknown command: %s", args[0])
	}
}

func createUserCmd(ctx context.Context, executor core.Executor, args []string) {
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
	runner := cli.NewRunner(ctx, executor, &cliObserver{})
	runner.ExecuteAction(action, fmt.Sprintf("✓ User %s created successfully", username))
}

func installPackageCmd(ctx context.Context, executor core.Executor, args []string) {
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
	runner := cli.NewRunner(ctx, executor, &cliObserver{})
	runner.ExecuteAction(action, fmt.Sprintf("✓ Package %s installed successfully", packageName))
}

func detectOSCmd(ctx context.Context, executor core.Executor) {
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
