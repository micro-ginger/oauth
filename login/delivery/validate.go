package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	s "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/login/validation"
)

func (h *lh[acc]) validate(request gateway.Request,
	sess *s.Session[acc]) errors.Error {
	if sess.Info.Account == nil {
		if sess.Info.AccountId == 0 {
			return nil
		}
		// get & fill account
		acc, err := h.account.GetById(request.GetContext(), sess.Info.AccountId)
		if err != nil {
			return err.WithTrace("account.GetById")
		}
		sess.Info.Account = acc
	}
	isValid := true
	for _, v := range sess.Flow.Login.Validations {
		if validation.StepBeforeDone != "" {
			if (sess.IsDone() && v.Step == validation.StepBeforeDone) ||
				(!sess.IsDone() && v.Step == validation.StepAfterDone) {
				continue
			}
		}
		isValid = true
		switch v.Type {
		case validation.TypeStatusPresent:
			isValid = sess.Info.Account.Status.Is(v.Status)
		case validation.TypeStatusAbsent:
			isValid = !sess.Info.Account.Status.Has(v.Status)
		}
		if !isValid {
			return errors.Forbidden().
				WithId(v.ErrorKey).
				WithDesc("account does not have required status")
		}
	}
	return nil
}
