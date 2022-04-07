package main

import (
	"os"

	"github.com/tasselsd/gorum/api"
	"github.com/tasselsd/gorum/pkg/core"
)

func main() {
	// Prepare configs
	configPath := "config.yaml"
	if len(os.Args) >= 2 {
		configPath = os.Args[1]
	}
	core.LoadConfig(configPath)

	// DB
	core.LoadDatabase()

	// Refresh Application
	core.RefreshApplication()

	// App engine start
	api.StartEngine()
}
