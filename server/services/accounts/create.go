package accounts

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
)

func (s *service) Create(ctx context.Context, userID core.ID, account string) (core.ID, error) {
	return s.store.Insert(ctx, &core.Account{
		ID:      core.ID(account),
		Address: account,
		Owner:   userID,
	})

}
