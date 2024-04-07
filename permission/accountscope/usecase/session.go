package usecase

import (
	"context"

	"github.com/ginger-core/errors"
	"github.com/micro-ginger/oauth/permission/scope/domain/scope"
	"github.com/micro-ginger/oauth/session/domain/session"
)

func (uc *useCase) SessionRemoveUnauthorized(ctx context.Context,
	session *session.Session) errors.Error {
	scopes, err := uc.GetAccountScopesFromScopes(ctx,
		session.Account.Id, session.Scopes, false)
	if err != nil {
		return err.
			WithTrace("SessionRemoveUnauthorized.GetAccountScopesFromScopes")
	}

	session.Scopes = make([]string, 0)
	for _, s := range scopes {
		session.Scopes = append(session.Scopes, s.Name)
	}

	// unauthorizedScopes, err := uc.repo.
	// 	ListUnauthorizedAccountScopes(ctx, session.Account.Id)
	// if err != nil {
	// 	return err.
	// 		WithTrace("SessionRemoveUnauthorized")
	// }

	// for i := len(session.Scopes) - 1; i >= 0; i-- {
	// 	es := session.Scopes[i]
	// 	unauthorized := false
	// 	for _, s := range unauthorizedScopes {
	// 		if es == s.Name {
	// 			unauthorized = true
	// 			break
	// 		}
	// 	}
	// 	if unauthorized {
	// 		session.Scopes = append(session.Scopes[:i], session.Scopes[i+1:]...)
	// 	}
	// }

	return nil
}

func (uc *useCase) SessionAddRequestedRoleScopes(ctx context.Context,
	session *session.Session) errors.Error {
	var scopes []*scope.Detailed
	var err errors.Error

	session.Roles = removeDuplicate(session.Roles)
	session.Scopes = removeDuplicate(session.Scopes)

	if len(session.Scopes) == 1 && session.Scopes[0] == "*" {
		scopes, err = uc.GetAllAccountScopes(ctx, session.Account.Id, false)
		if err != nil {
			return err.
				WithTrace("SessionAddRequestedRoleScopes.GetAllAccountScopes")
		}
	} else {
		for _, r := range session.Roles {
			if r == "*" {
				session.Roles = nil
				break
			}
		}
		scopes, err = uc.GetAccountScopesFromRoles(ctx,
			session.Account.Id, session.Roles, false)
		if err != nil {
			return err.
				WithTrace("SessionAddRequestedRoleScopes.GetAccountScopesFromRoles")
		}
	}

	seenRoles := make(map[string]bool)
	for _, s := range scopes {
		session.Scopes = append(session.Scopes, s.Name)
		if !seenRoles[s.RoleName] {
			session.Roles = append(session.Roles, s.RoleName)
		}
		seenRoles[s.RoleName] = true

	}
	return nil
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
