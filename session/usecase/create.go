package usecase

import (
	"context"
	"time"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/session/domain/session"
)

func (uc *useCase) Create(ctx context.Context,
	request *session.CreateRequest) (*session.Session, errors.Error) {
	conf := uc.config.Create
	if request.CreateConfig != nil {
		conf = *request.CreateConfig
	}
	s := request.Old
	if s == nil {
		s = &session.Session{
			CreatedAt:      time.Now().UTC(),
			AccessTokenExp: conf.AccessTokenExp,
			Account:        request.Account,
			Roles:          make([]string, 0),
			Scopes:         make([]string, 0),
		}
		if len(request.RequestedScopes) > 0 {
			s.Scopes = append(s.Scopes, request.RequestedScopes...)
		}
		if len(request.RequestedRoles) > 0 {
			s.Roles = append(s.Roles, request.RequestedRoles...)
		}
		// call handlers
		for _, h := range uc.handlerFuncs {
			if err := h(ctx, s); err != nil {
				return nil, err.WithTrace("session.Create.handlerFunc")
			}
		}
		// additional scopes will be added anyways
		s.Scopes = append(s.Scopes, conf.AdditionalScopes...)
	}
	//
	s.Id = uc.randomId()
	s.Section = conf.Section
	s.AccessToken = uc.generateToken(conf.AccessTokenLength)

	exp := conf.RefreshTokenExp
	if conf.AccessTokenExp > exp {
		exp = conf.AccessTokenExp
	}
	if err := uc.repo.Create(ctx,
		uc.getSessionKey(s.Account.Id, s.Id), s, exp); err != nil {
		return nil, err
	}
	if err := uc.repo.Create(ctx,
		uc.getAccessKey(s.AccessToken),
		s, conf.AccessTokenExp); err != nil {
		return nil, err
	}

	if conf.RefreshTokenExp > 0 {
		s.RefreshToken = uc.generateToken(conf.RefreshTokenLength)
		s.RefreshTokenExp = conf.RefreshTokenExp
		if err := uc.repo.Create(ctx,
			uc.getRefreshKey(s.RefreshToken),
			s, conf.RefreshTokenExp); err != nil {
			return nil, err
		}
	}
	return s, nil
}
