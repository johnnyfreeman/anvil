package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/johnnyfreeman/anvil/internal/actions"
	"github.com/johnnyfreeman/anvil/internal/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: anvil <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  create-user <username> [--group <group>]")
		fmt.Println("  detect-os")
		os.Exit(1)
	}

	ctx := context.Background()
	executor := core.LocalExecutor{}

	switch os.Args[1] {
	case "create-user":
		createUserCmd(ctx, executor, os.Args[2:])
	case "detect-os":
		detectOSCmd(ctx, executor)
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
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
	
	// Detect OS
	osInfo, err := core.DetectOS(ctx, executor)
	if err != nil {
		log.Fatalf("Failed to detect OS: %v", err)
	}
	
	// Create action
	var opts []actions.CreateUserOptsFunc
	if *group != "" {
		opts = append(opts, actions.WithGroup(*group))
	}
	
	action := actions.NewCreateUser(username, opts...)
	
	// Execute with basic observer
	observer := &cliObserver{}
	if err := action.Handle(ctx, executor, osInfo.Detected, observer); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	
	fmt.Printf("✓ User %s created successfully\n", username)
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
