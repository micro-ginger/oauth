package oauth

import (
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/app"
	rdd "github.com/micro-ginger/oauth/register/domain/delivery/register"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func NewApp[acc account.Model,
	regReq rdd.RequestModel, reg register.Model](configType string) app.Application {
	return app.New[acc, regReq, reg](configType)
}
