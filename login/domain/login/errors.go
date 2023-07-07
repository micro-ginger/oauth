package login

import "github.com/ginger-core/errors"

var (
	InvalidCredentialError = errors.Validation().
				WithId("InvalidCredentialsError").
				WithMessage("Invalid credentials, Invalid username or password")
	YourAccountDisabledError = errors.Validation().
					WithId("YourAccountDisabledError").
					WithMessage("Your account has been disabled to login, " +
			"Contact administrator and verify your information to access your account.")
)
