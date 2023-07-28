package verify

import (
	"encoding/json"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/login/authentication/otp"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *_handler[acc]) getOtp(sess *session.Session[acc]) (*otp.Otp, errors.Error) {
	_otp := new(otp.Otp)
	otpM := sess.Info.GetTemp(otpType)
	if otpM != nil {
		otpStr := otpM.(string)
		if err := json.Unmarshal([]byte(otpStr), _otp); err != nil {
			return nil, errors.New(err).
				WithTrace("_otp.json.Unmarshal")
		}
	}
	return _otp, nil
}
