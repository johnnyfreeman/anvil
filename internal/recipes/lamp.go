package recipes

import (
	"context"

	"github.com/johnnyfreeman/anvil/internal/actions"
	"github.com/johnnyfreeman/anvil/internal/core"
)

// LAMPServer recipe for setting up a complete LAMP (Linux, Apache, MySQL, PHP) stack
type LAMPServer struct {
	core.BaseRecipe
}

func NewLAMPServer() *LAMPServer {
	lampActions := []core.Action{
		// Update package list first
		actions.NewInstallPackage("", actions.WithUpdate()),
		
		// Install Apache web server
		actions.NewInstallPackage("apache2"),
		actions.NewEnableService("apache2"),
		actions.NewStartService("apache2"),
		
		// Install MySQL database server
		actions.NewInstallPackage("mysql-server"),
		actions.NewEnableService("mysql"),
		actions.NewStartService("mysql"),
		
		// Install PHP and common modules
		actions.NewInstallPackage("php"),
		actions.NewInstallPackage("libapache2-mod-php"),
		actions.NewInstallPackage("php-mysql"),
		actions.NewInstallPackage("php-cli"),
		actions.NewInstallPackage("php-curl"),
		actions.NewInstallPackage("php-gd"),
		actions.NewInstallPackage("php-mbstring"),
		actions.NewInstallPackage("php-xml"),
		actions.NewInstallPackage("php-zip"),
		
		// Restart Apache to load PHP module
		actions.NewRestartService("apache2"),
	}

	baseRecipe := core.NewBaseRecipe(
		"lamp-server",
		"Complete LAMP stack with Apache, MySQL, and PHP",
		lampActions,
	)

	return &LAMPServer{
		BaseRecipe: baseRecipe,
	}
}

// Custom execute method for LAMP recipe with additional setup
func (r *LAMPServer) Execute(ctx context.Context, ex core.Executor, os core.OS, observer core.ActionObserver) error {
	// Execute base recipe actions
	if err := r.BaseRecipe.Execute(ctx, ex, os, observer); err != nil {
		return err
	}

	// Additional LAMP-specific configuration could go here
	// For example: setting up virtual hosts, configuring PHP settings, etc.
	
	return nil
}

var _ core.Recipe = (*LAMPServer)(nil)