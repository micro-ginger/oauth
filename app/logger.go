package app

import "github.com/ginger-core/log"

func (a *app[acc]) initializeLogger() {
	a.Logger = log.NewLogger(a.Registry.ValueOf("logger"))
	a.Logger.SetSource("auth")
	a.Logger.Start()
}
