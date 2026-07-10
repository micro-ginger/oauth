package handler

import (
	"fmt"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

type handler[acc a.ExtentedModel] struct {
	logger log.Logger
	config config

	cache repository.Cache

	challengeGenerator session.ChallengeGenerator
}

func New[acc a.ExtentedModel](
	logger log.Logger, registry registry.Registry, cache repository.Cache,
) session.Handler[acc] {
	h := &handler[acc]{
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

func (h *handler[acc]) RegisterChallengeGenerator(generator session.ChallengeGenerator) {
	h.challengeGenerator = generator
}

func (h *handler[acc]) getChallengeKey(challenge string) string {
	key := fmt.Sprintf("login.sessions.%s", challenge)
	return key
}
