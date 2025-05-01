package login

import "github.com/ginger-core/gateway"

type Response[SessionAccountDetail gateway.ResultGetter] struct {
	Sessions map[string]*Session[SessionAccountDetail] `json:"sessions"`
}
