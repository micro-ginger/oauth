package app

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-gateway/ginger"
	"github.com/micro-blonde/auth/authorization"
)

func (a *App[acc, prof, regReq, reg, f]) initializeServer() {
	a.initializeHTTP()
	a.initializeAuthenticator()
}

func (a *App[acc, prof, regReq, reg, f]) initializeHTTP() {
	a.HTTP = ginger.NewHTTP(
		a.Logger.WithTrace("ginger.http"),
		a.Registry.ValueOf("gateway.http"),
	)
	responder := a.HTTP.NewResponder()
	controller := gateway.NewController(responder).WithLanguageBundle(a.Language)
	a.HTTP.SetController(controller)
}

func (a *App[acc, prof, regReq, reg, f]) initializeGrpc() {
	a.GRPC = ginger.NewGRPC(
		a.Logger.WithTrace("ginger.grpc"),
		a.Registry.ValueOf("gateway.grpc"),
	)
	a.GRPC.Initialize(a.GRPC.NewResponder())
}

func (a *App[acc, prof, regReq, reg, f]) initializeAuthenticator() {
	a.Authenticator = authorization.New[acc](
		a.HTTP, a.Registry.ValueOf("gateway.authorization"))
}
