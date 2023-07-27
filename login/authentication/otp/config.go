package otp

import (
	"fmt"
	"math"

	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type config struct {
	// Debug if set true otp code will be always
	Debug bool

	Challenge struct {
		// Length is length of challenge to be generated
		Length int
		// Characters are list of characters to be used while generating challenge
		Characters string
	}

	Code struct {
		// CodeLength length of code while generating one
		Length int
		// will be generated if not set in config
		Min int
		Max int
		// Format auto generated being used to format generated code.
		// e.g `%06d` is for 6 character long codes which is zero padded from left
		Format string
	}

	Validators struct {
		Session validator.Config
		Global  validator.Config
	}
}

func (c *config) Initialize() {
	if c.Code.Length == 0 {
		c.Code.Length = 6
	}
	if c.Code.Min == 0 || c.Code.Max == 0 {
		c.Code.Min = int(math.Pow10(c.Code.Length - 1))
		c.Code.Max = c.Code.Min * 10
	}
	c.Code.Format = "%0" + fmt.Sprint(c.Code.Length) + "d"

	if c.Challenge.Length == 0 {
		c.Challenge.Length = 32
	}
	if len(c.Challenge.Characters) < 2 {
		c.Challenge.Characters = "0123456789ABCDEFGHIJKLMOPQRSTUXYZabcdefghijklmopqrstuxyz@*!-_=^"
	}
}
