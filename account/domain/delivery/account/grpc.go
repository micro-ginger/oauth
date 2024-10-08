package account

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	blondeAcc "github.com/micro-blonde/auth/account"
	acc "github.com/micro-blonde/auth/proto/auth/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetGrpcAccount[T a.Model](
	a *a.Account[T]) (*acc.Account, errors.Error) {
	var v *structpb.Struct
	var t any = a.T
	if vg, ok := t.(blondeAcc.StructValueGetter); ok {
		var err error
		v, err = structpb.NewStruct(vg.GetValues())
		if err != nil {
			return nil, errors.New(err).
				WithTrace("structpb.NewStruct")
		}
	}
	r := &acc.Account{
		Id:     a.Id,
		Status: a.Status.Uint64(),
		T:      structpb.NewStructValue(v),
	}
	return r, nil
}

func GetGrpcAccounts[T a.Model](
	a []*a.Account[T]) (*acc.Accounts, errors.Error) {
	r := &acc.Accounts{
		Items: make([]*acc.Account, len(a)),
	}
	var err errors.Error
	for i, itm := range a {
		r.Items[i], err = GetGrpcAccount(itm)
		if err != nil {
			return nil, err.WithTrace("GetGrpcAccount")
		}
	}
	return r, nil
}

type BaseReadHandler[T a.Model] interface {
	GetInstruction() instruction.Instruction

	GetAccount(a *a.Account[T]) (*acc.Account, errors.Error)
}
