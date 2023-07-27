package info

import (
	"context"
	"fmt"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/account/domain/account"
)

type Handler[acc account.Model] interface {
	Generate(ctx context.Context) (*Info[acc], errors.Error)

	Save(ctx context.Context, info *Info[acc]) errors.Error
	Set(ctx context.Context, challenge, key string, value any) errors.Error

	Get(ctx context.Context, challenge string) (*Info[acc], errors.Error)
	GetItem(ctx context.Context, challenge, key string, ref any) errors.Error

	Delete(ctx context.Context, challenge string) errors.Error
	DeleteItem(ctx context.Context, challenge, key string) errors.Error
}

type handler[acc account.Model] struct {
	logger log.Logger
	config config

	cache repository.Cache

	challengeGenerator ChallengeGenerator
}

func New[acc account.Model](logger log.Logger, registry registry.Registry,
	cache repository.Cache) Handler[acc] {
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

func (h *handler[acc]) RegisterChallengeGenerator(generator ChallengeGenerator) {
	h.challengeGenerator = generator
}

func (h *handler[acc]) getChallengeKey(challenge string) string {
	key := fmt.Sprintf("challenge_%s", challenge)
	return key
}
