package validator

import (
	"context"

	"github.com/ginger-core/errors"
)

type UseCase interface {
	// NewTemplate returns clone of base validation containing base validation entries
	NewTemplate() *Validation

	// ValidateRequest only handles the validation logic
	ValidateRequest(ctx context.Context, v *Validation) errors.Error
	// Requested only handles the after generation login
	Requested(ctx context.Context, v *Validation)
	// BeginRequest validates and saves the validation in db
	// calls ValidateRequest, then saves into db
	BeginRequest(ctx context.Context, key any) (*Validation, errors.Error)
	// EndRequest handles generation logic and saves data into db
	// calls Generated function, then saves into db
	EndRequest(ctx context.Context, v *Validation)

	// ResetVerifies resets the verification validation
	ResetVerifies(ctx context.Context, v *Validation)
	// ValidateVerify only handles the validation logic of verify
	ValidateVerify(ctx context.Context, v *Validation) errors.Error
	// VerifyChecked only handles after verification logic
	VerifyChecked(ctx context.Context, v *Validation)
	// BeginVerify validates and saves the validation in db
	// calls ValidateVerify, then saves into db
	BeginVerify(ctx context.Context, key any) (*Validation, errors.Error)
	// EndVerify handles generation logic and saves data into db
	// calls VerifyChecked function, then saves into db
	EndVerify(ctx context.Context, v *Validation)

	Delete(ctx context.Context, key any) errors.Error
}
