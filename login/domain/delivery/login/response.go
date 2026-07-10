package login

import a "github.com/micro-blonde/auth/account"

type Response[acc a.ExtentedModel] struct {
	Sessions map[string]*Session[acc] `json:"sessions"`
}
