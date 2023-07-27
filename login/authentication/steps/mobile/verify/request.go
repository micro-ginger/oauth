package verify

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log/logger"
	"github.com/micro-ginger/oauth/login/authentication/info"
	"github.com/micro-ginger/oauth/login/authentication/response"
)

func (h *_handler[acc]) request(ctx context.Context,
	request gateway.Request, inf *info.Info[acc]) (response.Response, errors.Error) {
	o, remaining, err := h.otp.Generate(ctx, inf.Key, inf.Challenge, otpType)
	if err != nil {
		return nil, err
	}

	a, err := h.GetAccount(ctx, inf, request, nil)
	if err != nil && !err.IsType(errors.TypeNotFound) {
		return nil, err
	}
	if a != nil {
		// validate
		if err := h.CheckVerifyAccount(ctx, a); err != nil {
			return nil, err
		}
	}

	msgType := inf.GetTemp("msgType")
	if msgType == nil {
		msgType = h.config.NotificationMessageType
	}

	var mobile string
	if a != nil && a.T.GetMobile() != nil {
		mobile = *a.T.GetMobile()
		inf.SetTemp("mobile", mobile)
	} else {
		mob := inf.GetTemp("mobile")
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
	// 			Id:     inf.AccountId,
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
		Challenge: inf.Challenge,
		Remaining: uint(remaining.Seconds()),
		Detail:    detail,
	}
	if mobile := inf.GetTemp("mobile"); mobile != nil {
		detail["mobile"] = a.T.MaskMobile()
	}

	return resp, nil
}
