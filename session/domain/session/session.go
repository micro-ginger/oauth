package session

import "time"

type Session struct {
	Id        string
	CreatedAt time.Time

	AccessToken    string
	AccessTokenExp time.Duration

	RefreshToken    string
	RefreshTokenExp time.Duration

	Account Account

	Roles  []string
	Scopes []string
}
