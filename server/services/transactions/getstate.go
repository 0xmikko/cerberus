package transactions

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
)

func (s *service) GetState(ctx context.Context, transactionID core.ID) (bool, error) {
	transaction, err := s.store.FindByID(ctx, transactionID)
	if err != nil {
		return false, err
	}

	return transaction.State == core.Confirmed, nil
}
