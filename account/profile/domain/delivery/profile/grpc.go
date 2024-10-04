package delivery

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/micro-blonde/auth/account"
	"github.com/micro-blonde/auth/profile"
	prof "github.com/micro-blonde/auth/proto/auth/account/profile"
	p "github.com/micro-ginger/oauth/account/profile/domain/profile"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetGrpcProfile[T profile.Model](
	a *p.Profile[T]) (*prof.Profile, errors.Error) {
	if a == nil {
		return nil, nil
	}
	r := &prof.Profile{
		Id: a.Id,
		T:  structpb.NewNullValue(),
	}
	if a.Photo != nil {
		r.Photo = *a.Photo
	}
	var v *structpb.Struct
	var t any = a.T
	if vg, ok := t.(account.StructValueGetter); ok {
		var err error
		v, err = structpb.NewStruct(vg.GetValues())
		if err != nil {
			return nil, errors.New(err).
				WithTrace("structpb.NewStruct")
		}
		r.T = structpb.NewStructValue(v)
	}
	return r, nil
}

func GetGrpcProfiles[T profile.Model](
	a []*p.Profile[T]) (*prof.Profiles, errors.Error) {
	r := &prof.Profiles{
		Items: make([]*prof.Profile, len(a)),
	}
	var err errors.Error
	for i, itm := range a {
		r.Items[i], err = GetGrpcProfile(itm)
		if err != nil {
			return nil, err.WithTrace("GetGrpcProfile")
		}
	}
	return r, nil
}

type BaseReadHandler[T profile.Model] interface {
	GetInstruction() instruction.Instruction

	GetProfile(a *p.Profile[T]) (*prof.Profile, errors.Error)
}
