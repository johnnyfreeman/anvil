package recipes

import (
	"github.com/johnnyfreeman/anvil/internal/actions"
	"github.com/johnnyfreeman/anvil/internal/core"
)

// BasicWebServer recipe for setting up just Apache
type BasicWebServer struct {
	core.BaseRecipe
}

func NewBasicWebServer() *BasicWebServer {
	webActions := []core.Action{
		// Update packages
		actions.NewInstallPackage("", actions.WithUpdate()),
		
		// Install and configure Apache
		actions.NewInstallPackage("apache2"),
		actions.NewEnableService("apache2"),
		actions.NewStartService("apache2"),
		
		// Install common utilities
		actions.NewInstallPackage("curl"),
		actions.NewInstallPackage("wget"),
	}

	baseRecipe := core.NewBaseRecipe(
		"webserver",
		"Basic Apache web server setup",
		webActions,
	)

	return &BasicWebServer{
		BaseRecipe: baseRecipe,
	}
}

var _ core.Recipe = (*BasicWebServer)(nil)

// NginxWebServer recipe for setting up Nginx instead of Apache
type NginxWebServer struct {
	core.BaseRecipe
}

func NewNginxWebServer() *NginxWebServer {
	nginxActions := []core.Action{
		// Update packages
		actions.NewInstallPackage("", actions.WithUpdate()),
		
		// Install and configure Nginx
		actions.NewInstallPackage("nginx"),
		actions.NewEnableService("nginx"),
		actions.NewStartService("nginx"),
		
		// Install common utilities
		actions.NewInstallPackage("curl"),
		actions.NewInstallPackage("wget"),
	}

	baseRecipe := core.NewBaseRecipe(
		"nginx-webserver",
		"Nginx web server setup",
		nginxActions,
	)

	return &NginxWebServer{
		BaseRecipe: baseRecipe,
	}
}

var _ core.Recipe = (*NginxWebServer)(nil)