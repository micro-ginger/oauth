package login

import (
	"time"

	"github.com/ginger-core/gateway"
	a "github.com/micro-blonde/auth/account"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Session[SessionAccountDetail gateway.ResultGetter] struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`

	AccessToken       string `json:"accessToken"`
	AccessTokenExpSec uint   `json:"accessTokenExpSec"`

	RefreshToken       string `json:"refreshToken,omitempty"`
	RefreshTokenExpSec uint   `json:"refreshTokenExpSec,omitempty"`

	Scopes []string `json:"scopes"`

	Account any `json:"account,omitempty"`
}

func NewSession[acc a.ExtentedModel](
	session *session.Session[acc],
) *Session[acc] {
	return &Session[acc]{
		Id:                 session.Id,
		CreatedAt:          session.CreatedAt,
		AccessToken:        session.AccessToken,
		AccessTokenExpSec:  uint(session.AccessTokenExp.Seconds()),
		RefreshToken:       session.RefreshToken,
		RefreshTokenExpSec: uint(session.RefreshTokenExp.Seconds()),
		Scopes:             session.Scopes,
		Account:            session.Account.GetDeliveryResult(),
	}
}
