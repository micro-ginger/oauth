package flow

type Login struct {
	// roles
	// scopes
	// IncludeRoles containes roles to assign to account after login
	IncludeRoles []string
	// DefaultRoles defines default roles to give to
	// logged-in account if not passed the required roles
	DefaultRoles []string
}
