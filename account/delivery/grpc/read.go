package grpc

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/log"
	blondeAcc "github.com/micro-blonde/auth/account"
	acc "github.com/micro-blonde/auth/proto/auth/account"
	ins "github.com/micro-ginger/oauth/account/delivery/instruction"
	"github.com/micro-ginger/oauth/account/domain/account"
	a "github.com/micro-ginger/oauth/account/domain/account"
	"google.golang.org/protobuf/types/known/structpb"
)

type BaseReadHandler[T account.Model] interface {
}

type baseRead[T account.Model] struct {
	instruction instruction.Instruction
	logger      log.Logger
	uc          a.UseCase[T]
}

func newBaseRead[T account.Model](
	logger log.Logger, uc a.UseCase[T]) *baseRead[T] {
	h := &baseRead[T]{
		instruction: ins.NewInstruction(),
		logger:      logger,
		uc:          uc,
	}
	return h
}

func (h *baseRead[T]) getAccount(a *account.Account[T]) (*acc.Account, errors.Error) {
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
