package captcha

import (
	"encoding/json"
)

type Response struct {
	Base64            string `json:"base64"`
	Secret            string `json:"secret"`
	ExpirationSeconds int    `json:"remaining"`
	Code              string `json:"code,omitempty"`
}

func (r *Response) Bytes() []byte {
	result, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return result
}
