package steps

import (
	"github.com/micro-ginger/oauth/login/authentication/step"
	keyPw "github.com/micro-ginger/oauth/login/authentication/steps/key/password"
	mobileOtp "github.com/micro-ginger/oauth/login/authentication/steps/mobile/otp"
	mobileVerify "github.com/micro-ginger/oauth/login/authentication/steps/mobile/verify"
	"github.com/micro-ginger/oauth/login/authentication/steps/password"
	"github.com/micro-ginger/oauth/login/authentication/steps/refresh"
)

func GetByType(t step.Type) *step.Step {
	switch t {
	case keyPw.Type:
		return keyPw.Step
	case password.Type:
		return password.Step
	case mobileOtp.Type:
		return mobileOtp.Step
	case mobileVerify.Type:
		return mobileVerify.Step
	case refresh.Type:
		return refresh.Step
	}
	return nil
}
