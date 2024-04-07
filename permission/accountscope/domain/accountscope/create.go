package accountscope

import "context"

type CreateScope struct {
	ScopeId      uint64
	IsAuthorized *bool
}

type CreateScopeBulk []CreateScope

type CreatedScopeEventHandle func(ctx context.Context, scopeId, accountId uint64)
