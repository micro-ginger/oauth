package session

import (
	"time"

	a "github.com/micro-blonde/auth/account"
)

type Session[AccountDetail a.ExtentedModel] struct {
	Id        string
	CreatedAt time.Time

	Section string

	AccessToken    string
	AccessTokenExp time.Duration

	RefreshToken    string
	RefreshTokenExp time.Duration

	Account a.Account[AccountDetail]

	Roles  []string
	Scopes []string
}
