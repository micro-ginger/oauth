package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/ginger-core/query"
	"github.com/micro-ginger/oauth/account/domain/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	ad "github.com/micro-ginger/oauth/account/domain/delivery/account"
)

type resetPassword struct {
	gateway.Responder
	logger log.Logger
	uc     a.ResetPasswordHandler
}

func NewPasswordReset(logger log.Logger,
	uc a.ResetPasswordHandler, responder gateway.Responder) gateway.Handler {
	h := &resetPassword{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *resetPassword) Handle(request gateway.Request) (any, errors.Error) {
	ctx := request.GetContext()

	body := new(ad.ResetPassword)
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

	hashedPassword, err := account.HashPassword(body.NewPassword)
	if err != nil {
		return nil, err.WithTrace("account.HashPassword")
	}
	err = h.uc.ResetPassword(ctx, q, hashedPassword)
	if err != nil {
		return nil, err.WithTrace("uc.ResetPassword")
	}

	return nil, nil
}
