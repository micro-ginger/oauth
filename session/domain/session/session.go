package session

import (
	"time"

	"github.com/ginger-core/gateway"
)

type Session[AccountDetail gateway.ResultGetter] struct {
	Id        string
	CreatedAt time.Time

	Section string

	AccessToken    string
	AccessTokenExp time.Duration

	RefreshToken    string
	RefreshTokenExp time.Duration

	Account Account[AccountDetail]

	Roles  []string
	Scopes []string
}
