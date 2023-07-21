package handler

import "github.com/ginger-core/errors"

var (
	NotImplementedError = errors.Validation().
				WithId("InvalidCredentialsError").
				WithMessage("Invalid credentials, Invalid username or password").
				WithTrace("NotImplemented")
	InvalidCredentialError = errors.Validation().
				WithId("InvalidCredentialsError").
				WithMessage("Invalid credentials, Invalid username or password")
	YourAccountBlockedError = errors.Validation().
				WithId("YourAccountBlockedError").
				WithMessage("Your account has been disabled, " +
			"Contact administrator and verify your information to access your account.")
	YourAccountDisabledError = errors.Validation().
					WithId("YourAccountDisabledError").
					WithMessage("Your account has been disabled to login, " +
			"Contact administrator and verify your information to access your account.")
	InvalidCredentialCanMigrateError = errors.Validation().
						WithCode(644).
						WithId("InvalidCredentialCanMigrateError").
						WithMessage("Can not login, because your account has not " +
			"migrated to this system. You need to migrate your account first.")
)
