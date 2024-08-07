package grpc

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/ginger-core/errors/grpc"
	"github.com/ginger-core/log"
	"github.com/ginger-core/log/logger"
	"github.com/ginger-core/query"
	"github.com/micro-blonde/auth/profile"
	prof "github.com/micro-blonde/auth/proto/auth/account/profile"
	"github.com/micro-blonde/file"
	fileClient "github.com/micro-blonde/file/client"
	profDlv "github.com/micro-ginger/oauth/account/profile/domain/delivery/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
)

type ListHandler[T profile.Model, F file.Model] interface {
	p.GrpcProfilesGetter
	Initialize(file fileClient.Client[F])
}

type list[T profile.Model, F file.Model] struct {
	logger log.Logger
	uc     p.UseCase[T]
	file   fileClient.Client[F]
}

func NewList[T profile.Model, F file.Model](logger log.Logger,
	uc p.UseCase[T]) ListHandler[T, F] {
	h := &list[T, F]{
		logger: logger,
		uc:     uc,
	}
	return h
}

func (h *list[T, F]) Initialize(file fileClient.Client[F]) {
	h.file = file
}

func (h *list[T, F]) ListProfiles(ctx context.Context,
	request *prof.ListRequest) (*prof.Profiles, error) {
	r, err := h.listProfiles(ctx, request)
	if err != nil {
		h.logger.
			WithContext(ctx).
			With(logger.Field{
				"error": err.Error(),
				"ids":   request.Ids,
			}).
			Errorf("accounts request")
		return r, grpc.Generate(err)
	}
	h.logger.
		WithContext(ctx).
		With(logger.Field{
			"ids": request.Ids,
		}).
		Infof("accounts request")
	return r, nil
}

func (h *list[T, F]) listProfiles(ctx context.Context,
	request *prof.ListRequest) (*prof.Profiles, errors.Error) {
	var err errors.Error
	var ps []*p.Profile[T]
	r := new(prof.Profiles)
	if len(request.Ids) > 0 {
		q := query.NewFilter(query.New(ctx)).
			WithMatch(&query.Match{
				Key:      "id",
				Operator: query.In,
				Value:    request.Ids,
			})
		ps, err = h.uc.List(ctx, q)
		if err != nil {
			return r, err.
				WithTrace("uc.List")
		}
	} else {
		return r, errors.Validation().
			WithMessage("no reference given")
	}
	r, err = profDlv.GetGrpcProfiles(ps)
	if err != nil {
		return nil, err.
			WithTrace("delivery.GetGrpcAccounts")
	}
	for _, itm := range r.Items {
		if itm.Photo != "" {
			url, err := h.file.GetDownloadUrlByKey(itm.Photo)
			if err != nil {
				h.logger.
					With(logger.Field{
						"error": err.Error(),
					}).
					WithTrace("file.GetDownloadUrlByKey").
					Errorf("error on get download url by key")
				itm.Photo = ""
			} else {
				itm.Photo = url
			}
		}
	}
	return r, nil
}
