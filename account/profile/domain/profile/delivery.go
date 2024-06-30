package profile

import (
	"context"

	"github.com/micro-blonde/auth/proto/auth/account/profile"
)

type GrpcProfileGetter interface {
	GetProfile(ctx context.Context,
		request *profile.GetRequest) (*profile.Profile, error)
}

type GrpcProfilesGetter interface {
	ListProfiles(ctx context.Context,
		request *profile.ListRequest) (*profile.Profiles, error)
}
