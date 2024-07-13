package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	cpt "github.com/micro-blonde/auth/captcha"
)

func (uc *useCase) Generate(ctx context.Context) (*cpt.Captcha, errors.Error) {
	c, err := uc.generator.New(ctx)
	if err != nil {
		return nil, err.WithTrace("generator.New")
	}
	if !uc.config.Debug {
		c.Code = ""
	}
	return c, nil
}
