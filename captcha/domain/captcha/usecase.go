package captcha

import (
	"context"

	"github.com/ginger-core/errors"
	cpt "github.com/micro-blonde/auth/captcha"
)

type UseCase interface {
	Initialize(generator cpt.Generator)
	Generate(ctx context.Context) (*cpt.Captcha, errors.Error)
}
