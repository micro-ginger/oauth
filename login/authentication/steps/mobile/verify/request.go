package verify

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log/logger"
	"github.com/micro-ginger/oauth/login/authentication/response"
	"github.com/micro-ginger/oauth/login/session/domain/session"
)

func (h *_handler[acc]) request(ctx context.Context, request gateway.Request,
	sess *session.Session[acc]) (response.Response, errors.Error) {
	o, remaining, err := h.otp.Generate(ctx, sess.Key, sess.Challenge, otpType)
	if err != nil {
		return nil, err
	}

	a, err := h.GetAccount(ctx, sess.Info, request, nil)
	if err != nil && !err.IsType(errors.TypeNotFound) {
		return nil, err
	}
	if a != nil {
		// validate
		if err := h.CheckVerifyAccount(ctx, a); err != nil {
			return nil, err
		}
	}

	msgType := sess.Info.GetTemp("msgType")
	if msgType == nil {
		msgType = h.config.NotificationMessageType
	}

	var mobile string
	if a != nil && a.T.GetMobile() != nil {
		mobile = *a.T.GetMobile()
		sess.Info.SetTemp("mobile", mobile)
	} else {
		mob := sess.Info.GetTemp("mobile")
		if mob != nil {
			mobile = mob.(string)
		}
	}

	if mobile == "" {
		return nil, errors.Validation().
			WithTrace("otp.request.empty.mobile")
	}

	// if !h.config.Debug {
	// 	// TODO send otp
	// 	msg := &message.Message{
	// 		Type: fmt.Sprint(msgType),
	// 		Receiver: &message.Receiver{
	// 			Id:     sess.Info.AccountId,
	// 			Mobile: mobile,
	// 		},
	// 		Meta: map[string]*message.Meta{
	// 			"code": {
	// 				String_: o.Code,
	// 			},
	// 		},
	// 	}
	// 	if err := h.notification.Send(composition.NewBackgroundContext(ctx), msg); err != nil {
	// 		return nil, errors.Internal(err).
	// 			WithTrace("mobile.otp.send.Err").
	// 			WithDesc("error on send otp")
	// 	}
	// }
	h.logger.With(logger.Field{"otp": o}).Debugf("generated otp")

	detail := make(map[string]any)
	resp := &response.BaseResponse{
		State:     response.StateOtpSent,
		Challenge: sess.Challenge,
		Remaining: uint(remaining.Seconds()),
		Detail:    detail,
	}
	if mobile := sess.Info.GetTemp("mobile"); mobile != nil {
		detail["mobile"] = a.T.MaskMobile()
	}

	return resp, nil
}
