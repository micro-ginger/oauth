package account

import (
	"context"

	"github.com/micro-blonde/auth/proto/auth/account"
)

type GrpcAccountGetter interface {
	GetAccount(ctx context.Context,
		request *account.GetRequest) (*account.Account, error)
}

type GrpcAccountsGetter interface {
	ListAccounts(ctx context.Context,
		request *account.ListRequest) (*account.Accounts, error)
	ListAccountProfiles(ctx context.Context,
		request *account.ListRequest) (*account.AccountProfiles, error)
}
