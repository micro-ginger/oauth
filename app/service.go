package app

import (
	fileClient "github.com/micro-blonde/file/client"
)

func (a *App[acc, prof, regReq, reg, f, SessionAccountDetail]) initializeServices() {
	a.initiateFileClient()
	//
	a.File.Initialize()
}

func (a *App[acc, prof, regReq, reg, f, SessionAccountDetail]) initiateFileClient() {
	a.File = fileClient.New[f](
		a.Logger.WithTrace("services.file"),
		a.Registry.ValueOf("services.file"),
	)
}
