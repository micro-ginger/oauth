package oauth

import (
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/app"
)

func NewApp[acc account.Model](configType string) app.Application {
	return app.New[acc](configType)
}
