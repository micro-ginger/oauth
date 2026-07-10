package oauth

import (
	"github.com/micro-blonde/auth/account"
	"github.com/micro-blonde/auth/profile"
	"github.com/micro-blonde/file"
	"github.com/micro-ginger/oauth/app"
	rdd "github.com/micro-ginger/oauth/register/domain/delivery/register"
	"github.com/micro-ginger/oauth/register/domain/register"
)

func NewApp[acc account.ExtentedModel, prof profile.Model,
	regReq rdd.RequestModel, reg register.Model, f file.Model]() app.Application {
	return app.New[acc, prof, regReq, reg, f]()
}
