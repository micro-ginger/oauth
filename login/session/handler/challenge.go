package handler

import (
	"crypto/rand"

	"github.com/ginger-core/errors"
)

func (h *handler[acc]) GenerateChallenge(chars string,
	length int) (string, errors.Error) {
	charLen := len(chars)
	maxCharInd := 255 - (256 % charLen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			return "", errors.New(err)
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxCharInd {
				// Skip this number to avoid modulo bias.
				continue
			}
			b[i] = chars[c%charLen]
			i++
			if i == length {
				return string(b), nil
			}
		}
	}
}
