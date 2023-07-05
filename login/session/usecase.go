package session

import (
	"context"

	"github.com/ginger-core/compound/registry"
	"github.com/ginger-core/errors"
	"github.com/ginger-core/log"
	"github.com/ginger-core/repository"
)

type UseCase interface {
	Generate(request *GenerateRequest) (*Session, errors.Error)
	Store(ctx context.Context, session *Session) (err errors.Error)
	Delete(ctx context.Context, session *Session) errors.Error
}

type useCase struct {
	logger log.Logger
	config config

	cache              repository.Cache
	challengeGenerator ChallengeGenerator
}

func New(logger log.Logger, registry registry.Registry,
	cache repository.Cache) UseCase {
	uc := &useCase{
		logger: logger,
		cache:  cache,
	}

	if err := registry.Unmarshal(&uc.config); err != nil {
		panic(err)
	}
	uc.config.initialize()

	uc.RegisterChallengeGenerator(uc.generateChallenge)
	return uc
}

func (uc *useCase) RegisterChallengeGenerator(generator ChallengeGenerator) {
	uc.challengeGenerator = generator
}
