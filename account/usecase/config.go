package usecase

import "github.com/micro-blonde/auth/account"

type config struct {
	Password struct {
		MinLen        int
		MinComplexity int
	}
	InternalStatus map[account.InternalStatus]struct {
		AddRoles []string
	}
}

func (c *config) initialize() {
	if c.Password.MinLen == 0 {
		c.Password.MinLen = 6
	}
	if c.Password.MinComplexity == 0 {
		c.Password.MinComplexity = 2
	}
}
