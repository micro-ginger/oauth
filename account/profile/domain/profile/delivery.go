package profile

import (
	"context"

	"github.com/micro-blonde/auth/proto/auth/account/profile"
)

type GrpcProfileGetter interface {
	GetProfile(ctx context.Context,
		request *profile.GetRequest) (*profile.Profile, error)
}
