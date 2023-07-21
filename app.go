package oauth

import (
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/app"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func NewApp[acc account.Model, reg register.Model](configType string) app.Application {
	return app.New[acc, reg](configType)
}
