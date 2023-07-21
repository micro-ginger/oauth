package validator

import "github.com/ginger-core/errors"

const (
	RetryRequestAfterErrorCode = 732
	TooManyRequestsErrorCode   = 429
)

var TooManyRequestsError = errors.TooManyRequests().
	WithId("TooManyRequestsError").
	WithCode(TooManyRequestsErrorCode).
	WithMessage("You are temporarily blocked as " +
		"you have made too many requests lately. Please try again later")

var RequestExpiredError = errors.Validation().
	WithId("ValidatorRequestExpiredError").
	WithMessage("Invalid request, Please try again.")
