package usecase

import (
	"fmt"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type useCase[AccountDetail gateway.ResultGetter] struct {
	logger log.Logger
	config session.Config
	repo   session.Repository[AccountDetail]

	handlerFuncs []session.SessionHandlerFunc[AccountDetail]
}

func New[AccountDetail gateway.ResultGetter](logger log.Logger, registry registry.Registry,
	repo session.Repository[AccountDetail]) session.UseCase[AccountDetail] {
	uc := &useCase[AccountDetail]{
		logger:       logger,
		repo:         repo,
		handlerFuncs: make([]session.SessionHandlerFunc[AccountDetail], 0),
	}
	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.Initialize()
	return uc
}

func (uc *useCase[AccountDetail]) RegisterSessionHandlers(
	hs ...session.SessionHandlerFunc[AccountDetail],
) {
	uc.handlerFuncs = append(uc.handlerFuncs, hs...)
}

func (uc *useCase[AccountDetail]) getSessionKey(accId uint64, id string) string {
	return fmt.Sprintf("session_%d_%s", accId, id)
}

func (uc *useCase[AccountDetail]) getAccessKey(token string) string {
	return "access_" + token
}

func (uc *useCase[AccountDetail]) getRefreshKey(token string) string {
	return "refresh_" + token
}
