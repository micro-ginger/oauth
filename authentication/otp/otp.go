package otp

import (
	"encoding/json"

	"github.com/micro-ginger/oauth/validator/domain/validator"
)

type Otp struct {
	Key  any
	Code string

	Validation *validator.Validation
}

func (o *Otp) MarshalBinary() (data []byte, err error) {
	var bytes []byte
	bytes, err = json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
