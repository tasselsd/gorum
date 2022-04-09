package core

import "fmt"

var STARTUP_HOOKS []func()

var SHUTDOWN_HOOKS []func()

func RefreshApplication() {
	for _, hook := range STARTUP_HOOKS {
		hook()
	}
}

func GracefulShutdown() {
	fmt.Println()
	for _, hook := range SHUTDOWN_HOOKS {
		hook()
	}
	fmt.Printf("[INFO] graceful shutdown\n")
}
