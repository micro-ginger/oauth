package session

import "github.com/ginger-core/gateway"

type Account[Detail gateway.ResultGetter] struct {
	Id uint64

	Detail Detail
}
