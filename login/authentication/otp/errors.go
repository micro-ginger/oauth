package otp

import "github.com/ginger-core/errors"

var InvalidCodeError = errors.Validation().
	WithId("InvalidCodeError").
	WithMessage("Invalid code.")

var CodeExpiredError = errors.Validation().
	WithId("CodeExpiredError").
	WithMessage("Your request has expired, Please try again.")
