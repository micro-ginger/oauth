package authorization

import (
	"github.com/ginger-core/errors"
	"github.com/ginger-core/gateway"
	"github.com/micro-ginger/oauth/account/domain/account"
	ld "github.com/micro-ginger/oauth/login/domain/delivery/login"
	"github.com/micro-ginger/oauth/login/flow"
	s "github.com/micro-ginger/oauth/login/session/domain/session"
	"github.com/micro-ginger/oauth/session/domain/session"
)

type Manager[acc account.Model] interface {
	BeforeHandle(request gateway.Request) errors.Error
	BeforeStart(request gateway.Request,
		queries *ld.Request, flow *flow.Flow) errors.Error
	BeforeLogin(request gateway.Request, sess *s.Session[acc]) errors.Error
	BeforeSessionCreate(request gateway.Request,
		sess *s.Session[acc], sessions []*session.CreateRequest) errors.Error
	AfterSessionCreate(request gateway.Request,
		sess *s.Session[acc], resp *ld.Response) errors.Error
}
