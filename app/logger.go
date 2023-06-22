package app

import "github.com/ginger-core/log"

func (a *app) initializeLogger() {
	a.logger = log.NewLogger(a.registry.ValueOf("logger"))
	a.logger.SetSource("auth")
	a.logger.Start()
}
