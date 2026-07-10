package session

import (
	"github.com/micro-blonde/auth/account"
	a "github.com/micro-blonde/auth/account"
)

type CreateRequest[AccountDetail account.ExtentedModel] struct {
	Account *a.Account[AccountDetail]

	CreateConfig *CreateConfig
	Old          *Session[AccountDetail]
	// RequestedScopes is scopes that user requested for. and
	// must be checked before giving the permission
	RequestedScopes []string
	// RequestedRoles is roles that user requested for. and
	// must be checked before giving the permission
	RequestedRoles []string
}
