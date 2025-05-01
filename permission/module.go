package permission

import (
	"github.com/ginger-core/gateway"
	"github.com/ginger-core/log"
	dl "github.com/ginger-core/repository"
	"github.com/micro-ginger/oauth/permission/accountrole"
	"github.com/micro-ginger/oauth/permission/accountscope"
	"github.com/micro-ginger/oauth/permission/role"
	rs "github.com/micro-ginger/oauth/permission/rolescope"
	"github.com/micro-ginger/oauth/permission/scope"
)

type Module[SessionAccountDetail gateway.ResultGetter] struct {
	Scope        *scope.Module
	Role         *role.Module
	RoleScope    *rs.Module
	AccountScope *accountscope.Module[SessionAccountDetail]
	AccountRole  *accountrole.Module
}

func Initialize[SessionAccountDetail gateway.ResultGetter](
	logger log.Logger, baseDb dl.Repository,
) *Module[SessionAccountDetail] {
	mod := &Module[SessionAccountDetail]{
		Scope: scope.Initialize(
			logger.WithTrace("scope"),
			baseDb,
		),
		Role: role.Initialize(
			logger.WithTrace("role"),
			baseDb,
		),
		RoleScope: rs.Initialize(
			logger.WithTrace("roleScope"),
			baseDb,
		),
		AccountScope: accountscope.Initialize[SessionAccountDetail](
			logger.WithTrace("accountScope"),
			baseDb,
		),
		AccountRole: accountrole.Initialize(
			logger.WithTrace("accountRole"),
			baseDb,
		),
	}
	return mod
}
