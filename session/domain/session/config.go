package session

import (
	"time"
)

type CreateConfig struct {
	Section         string
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration

	AccessTokenLength  int
	RefreshTokenLength int
	// DefaultRoles defines default roles to assign to account
	// if the required roles parameter is missing
	DefaultRoles []string
	// AdditionalScopes is scopes that must be added to user permission anyways
	AdditionalScopes []string
	// IncludeRoles containes roles to assign to
	// account if has permission
	IncludeRoles []string
}

func NewCreateConfigFromSession(s *Session) *CreateConfig {
	r := &CreateConfig{
		Section:            s.Section,
		AccessTokenExp:     s.AccessTokenExp,
		RefreshTokenExp:    s.RefreshTokenExp,
		AccessTokenLength:  len(s.AccessToken),
		RefreshTokenLength: len(s.RefreshToken),
		IncludeRoles:       s.Roles,
	}
	return r
}

func (c *CreateConfig) Initialize() {
	if c == nil {
		return
	}
	if c.Section == "" {
		c.Section = "DEFAULT"
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
