package session

import "time"

type CreateConfig struct {
	Key             string
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration

	AccessTokenLength  int
	RefreshTokenLength int
	// AdditionalScopes is scopes that must be added to user permission anyways
	AdditionalScopes []string
	// IsAnonymousOK check if requested account must exist in our system or not
	IsAnonymousOK bool
}

func (c *CreateConfig) Initialize() {
	if c == nil {
		return
	}
	if c.Key == "" {
		c.Key = "DEFAULT"
	}
	if c.AccessTokenExp == 0 {
		c.AccessTokenExp = time.Hour
	}
	//if c.RefreshTokenExp == 0 {
	//	c.RefreshTokenExp = time.Hour * 24 * 7
	//}
	if c.AccessTokenLength == 0 {
		c.AccessTokenLength = 64
	}
	if c.RefreshTokenLength == 0 {
		c.RefreshTokenLength = 128
	}
}

type Config struct {
	Create CreateConfig
}

func (c *Config) Initialize() {
	c.Create.Initialize()
}
