package login

import (
	"time"

	"github.com/ginger-core/gateway"
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

func NewSession[SessionAccountDetail gateway.ResultGetter](
	session *session.Session[SessionAccountDetail],
) *Session[SessionAccountDetail] {
	return &Session[SessionAccountDetail]{
		Id:                 session.Id,
		CreatedAt:          session.CreatedAt,
		AccessToken:        session.AccessToken,
		AccessTokenExpSec:  uint(session.AccessTokenExp.Seconds()),
		RefreshToken:       session.RefreshToken,
		RefreshTokenExpSec: uint(session.RefreshTokenExp.Seconds()),
		Scopes:             session.Scopes,
		Account:            session.Account.Detail.GetResult(),
	}
}
