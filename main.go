package main

import (
	"github.com/eldada/metrics-viewer/commands"
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "metrics-viewer"
	app.Description = "Easily present Open Metrics data in terminal."
	app.Version = "v0.2.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetGraphCommand(),
		commands.GetPrintCommand(),
	}
}
