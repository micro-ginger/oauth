package handler

import "time"

type config struct {
	Expiration time.Duration
	Challenge  struct {
		// Length is length of challenge to be generated
		Length int
		// Characters are list of characters to be used while generating challenge
		Characters string
	}
}

func (c *config) Initialize() {
	if c.Expiration == 0 {
		c.Expiration = time.Minute * 10
	}
	if c.Challenge.Length == 0 {
		c.Challenge.Length = 32
	}
	if len(c.Challenge.Characters) < 2 {
		c.Challenge.Characters =
			"0123456789ABCDEFGHIJKLMOPQRSTUXYZabcdefghijklmopqrstuxyz@*!-_=^"
	}
}

const (
	root = "_"
)
