package otp

import (
	"fmt"
	"math/rand"
)

func (h *handler[acc]) GenerateCode() string {
	var code int
	if !h.config.Debug {
		code = rand.Intn(h.config.Code.Max-h.config.Code.Min) + h.config.Code.Min
	}
	return fmt.Sprintf(h.config.Code.Format, code)
}
