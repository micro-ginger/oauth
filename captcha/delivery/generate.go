package delivery

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	"github.com/micro-ginger/oauth/captcha/domain/captcha"
)

type generateHandler struct {
	gateway.Responder
	logger log.Logger
	uc     captcha.UseCase
}

func NewGenerate(logger log.Logger, uc captcha.UseCase,
	responder gateway.Responder) gateway.Handler {
	h := &generateHandler{
		Responder: responder,
		logger:    logger,
		uc:        uc,
	}
	return h
}

func (h *generateHandler) Handle(request gateway.Request) (any, errors.Error) {
	ctx, cancel := context.WithTimeout(request.GetContext(), time.Second*3)
	c, err := h.uc.Generate(ctx)
	cancel()
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(make([]byte, 0))
	if _, err := c.Image.WriteTo(buffer); err != nil {
		return nil, errors.Internal(err).WithContext(request.GetContext())
	}

	response := gateway.NewResponse().
		WithContentType(gateway.ContentTypeJson).
		WithStatus(gateway.StatusOK).
		WithHeader("Cache-Control", "no-cache, no-store, must-revalidate").
		WithHeader("Pragma", "no-cache").
		WithHeader("Expires", "0")

	if _, ok := request.GetQuery("img"); ok {
		// send image
		resp := response.
			WithContentType(gateway.ContentTypeImageJpeg).
			WithHeader("X-Captcha-Secret", c.Secret).
			WithHeader("X-Captcha-Expiration-Seconds", fmt.Sprint(int(c.Expiration.Seconds()))).
			WithBody(buffer)
		if c.Code != "" {
			// debug mode
			resp = resp.WithHeader("X-Captcha-Code", c.Code)
		}
		return resp, nil
	}
	str := base64.StdEncoding.EncodeToString(buffer.Bytes())

	resp := &captcha.Response{
		Base64:            str,
		Secret:            c.Secret,
		ExpirationSeconds: int(c.Expiration.Seconds()),
		Code:              c.Code,
	}

	return response.
		WithBody(resp), nil
}
