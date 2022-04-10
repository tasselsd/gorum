package main

import (
	"os"

	"github.com/tasselsd/gorum/api"
	"github.com/tasselsd/gorum/pkg/core"
)

func main() {
	defer core.GracefulShutdown()
	core.LoadConfig(os.Args)
	core.LoadDatabase()
	core.RefreshApplication()
	api.StartEngine()
}
