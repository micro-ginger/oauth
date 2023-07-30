package login

import (
	"time"

	"github.com/micro-ginger/oauth/session/domain/session"
)

type Session struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`

	AccessToken       string `json:"accessToken"`
	AccessTokenExpSec uint   `json:"accessTokenExpSec"`

	RefreshToken       string `json:"refreshToken,omitempty"`
	RefreshTokenExpSec uint   `json:"refreshTokenExpSec,omitempty"`

	Scopes []string `json:"scopes"`
}

func NewSession(session *session.Session) *Session {
	return &Session{
		Id:                 session.Id,
		CreatedAt:          session.CreatedAt,
		AccessToken:        session.AccessToken,
		AccessTokenExpSec:  uint(session.AccessTokenExp.Seconds()),
		RefreshToken:       session.RefreshToken,
		RefreshTokenExpSec: uint(session.RefreshTokenExp.Seconds()),
		Scopes:             session.Scopes,
	}
}
