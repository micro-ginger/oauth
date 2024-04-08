package session

type CreateRequest struct {
	Account Account

	CreateConfig *CreateConfig
	Old          *Session
	// RequestedScopes is scopes that user requested for. and
	// must be checked before giving the permission
	RequestedScopes []string
	// RequestedRoles is roles that user requested for. and
	// must be checked before giving the permission
	RequestedRoles []string
}
