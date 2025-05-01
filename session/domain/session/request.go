package session

import "github.com/ginger-core/gateway"

type CreateRequest[AccountDetail gateway.ResultGetter] struct {
	Account Account[AccountDetail]

	CreateConfig *CreateConfig
	Old          *Session[AccountDetail]
	// RequestedScopes is scopes that user requested for. and
	// must be checked before giving the permission
	RequestedScopes []string
	// RequestedRoles is roles that user requested for. and
	// must be checked before giving the permission
	RequestedRoles []string
}
