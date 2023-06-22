package app

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger"
)

func (a *app) initializeServer() {
	a.initializeGinger()
}

func (a *app) initializeGinger() {
	logger := a.logger.WithTrace("ginger")
	a.ginger = ginger.NewServer(logger, a.registry.ValueOf("gateway.http"))

	responder := a.ginger.NewResponder()
	controller := gateway.NewController(responder).WithLanguageBundle(a.language)
	a.ginger.SetController(controller)
}
