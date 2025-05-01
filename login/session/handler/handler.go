package handler

import (
	"fmt"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type handler[acc account.Model, SessionAccountDetail gateway.ResultGetter] struct {
	logger log.Logger
	config config

	cache repository.Cache

	challengeGenerator session.ChallengeGenerator
}

func New[acc account.Model, SessionAccountDetail gateway.ResultGetter](
	logger log.Logger, registry registry.Registry, cache repository.Cache,
) session.Handler[acc, SessionAccountDetail] {
	h := &handler[acc, SessionAccountDetail]{
		logger: logger,
		cache:  cache,
	}
	h.RegisterChallengeGenerator(h.GenerateChallenge)

	if registry != nil {
		if err := registry.Unmarshal(&h.config); err != nil {
			panic(err)
		}
	}
	h.config.Initialize()
	return h
}

func (h *handler[acc, SessionAccountDetail]) RegisterChallengeGenerator(generator session.ChallengeGenerator) {
	h.challengeGenerator = generator
}

func (h *handler[acc, SessionAccountDetail]) getChallengeKey(challenge string) string {
	key := fmt.Sprintf("login.sessions.%s", challenge)
	return key
}
