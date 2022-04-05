package main

import (
	"github.com/tasselsd/gorum/pkg/core"
)

func main() {
	// Prepare configs
	core.LoadConfig()
	// DB
	core.LoadDatabase()
	// App engine start
	core.StartEngine()
}
