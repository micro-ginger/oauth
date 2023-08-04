package register

import (
	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/account/domain/account"
	"github.com/micro-ginger/oauth/register/domain/register"
)

type RequestModel interface {
}

type RequestModelHandler[R RequestModel,
	T register.Model, acc account.Model] interface {
	New() *Request[R]
	PopulateRequest(body *Request[R], req *register.Request[T, acc]) errors.Error
}

// var requestModelHandler RequestModelHandler[register.Model, account.Model]

// func SetRequestModelHandler(m RequestModelHandler[register.Model, account.Model]) {
// 	requestModelHandler = m
// }

// func NewRequest() *Request {
// 	r := new(Request)
// 	if requestModelHandler != nil {
// 		r.T = requestModelHandler.New()
// 	}
// 	return r
// }

// func PopulateRequest[T register.Model,
// 	acc account.Model](body *Request, req register.Request[T, acc]) {
// 	if requestModelHandler != nil {
// 		requestModelHandler.PopulateRequest(body)
// 	}
// }
