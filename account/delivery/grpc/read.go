package grpc

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway/instruction"
	"github.com/ginger-core/log"
	blondeAcc "github.com/micro-blonde/auth/account"
	acc "github.com/micro-blonde/auth/proto/auth/account"
	ins "github.com/micro-ginger/oauth/account/delivery/instruction"
	a "github.com/micro-ginger/oauth/account/domain/account"
	ad "github.com/micro-ginger/oauth/account/domain/delivery/account"
	"google.golang.org/protobuf/types/known/structpb"
)

type baseRead[T a.Model] struct {
	instruction instruction.Instruction
	logger      log.Logger
	uc          a.UseCase[T]
}

func newBaseRead[T a.Model](
	logger log.Logger, uc a.UseCase[T]) ad.BaseReadHandler[T] {
	h := &baseRead[T]{
		instruction: ins.NewInstruction(),
		logger:      logger,
		uc:          uc,
	}
	return h
}

func (h *baseRead[T]) GetInstruction() instruction.Instruction {
	return h.instruction
}

func (h *baseRead[T]) GetAccount(a *a.Account[T]) (*acc.Account, errors.Error) {
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
