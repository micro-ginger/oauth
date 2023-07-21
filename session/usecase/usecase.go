package usecase

import (
	"fmt"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type useCase struct {
	logger log.Logger
	config session.Config
	repo   session.Repository

	handlerFuncs []session.SessionHandlerFunc
}

func New(logger log.Logger, registry registry.Registry,
	repo session.Repository) session.UseCase {
	uc := &useCase{
		logger:       logger,
		repo:         repo,
		handlerFuncs: make([]session.SessionHandlerFunc, 0),
	}
	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.Initialize()
	return uc
}

func (uc *useCase) RegisterSessionHandlers(hs ...session.SessionHandlerFunc) {
	uc.handlerFuncs = append(uc.handlerFuncs, hs...)
}

func (uc *useCase) getSessionKey(accId uint64, id string) string {
	return fmt.Sprintf("session_%d_%s", accId, id)
}

func (uc *useCase) getAccessKey(token string) string {
	return "access_" + token
}

func (uc *useCase) getRefreshKey(token string) string {
	return "refresh_" + token
}
