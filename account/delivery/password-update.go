package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	ad "github.com/micro-ginger/oauth/account/domain/delivery/account"
)

type updatePassword[T account.Model] struct {
	gateway.Responder
	logger log.Logger
	uc     a.PasswordUpdateHandler[T]
}

func NewPasswordUpdate[T account.Model](logger log.Logger,
	uc a.PasswordUpdateHandler[T], responder gateway.Responder) gateway.Handler {
	h := &updatePassword[T]{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *updatePassword[T]) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	body := new(ad.UpdatePassword)
	if err := request.ProcessBody(body); err != nil {
		return nil, errors.Validation(err)
	}

	err := h.uc.ValidatePassword(ctx, body.NewPassword)
	if err != nil {
		return nil, err.WithTrace("uc.ValidatePassword")
	}

	q := query.NewFilter(query.New(ctx)).
		WithMatch(&query.Match{
			Key:      "id",
			Operator: query.Equal,
			Value:    request.GetAuthorization().GetApplicantId(),
		})

	acc, err := h.uc.Get(ctx, q)
	if err != nil {
		return nil, err.WithTrace("uc.Get")
	}
	if err := acc.MatchPassword(body.OldPassword); err != nil {
		return nil, errors.Forbidden(err).
			WithId("InvalidOldPasswordError").
			WithMessage("entered old password is not currect").
			WithTrace("InvalidOldPassword").
			WithExtra("oldPassword",
				errors.NewExtra(
					errors.Forbidden(err).
						WithId("InvalidOldPasswordError").
						WithMessage("entered old password is not currect"),
				),
			)
	}

	hashedPassword, err := a.HashPassword(body.NewPassword)
	if err != nil {
		return nil, err.WithTrace("account.HashPassword")
	}
	err = h.uc.UpdatePassword(ctx, q, hashedPassword)
	if err != nil {
		return nil, err.WithTrace("uc.UpdatePassword")
	}

	return nil, nil
}
