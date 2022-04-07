package core

var HOOKS []func()

func RefreshApplication() {
	for _, hook := range HOOKS {
		hook()
	}
}
