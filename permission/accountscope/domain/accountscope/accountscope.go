package accountscope

type AccountScope struct {
	AccountId    uint64
	ScopeId      uint64
	IsAuthorized *bool
}

func (*AccountScope) GetId() any {
	return nil
}
