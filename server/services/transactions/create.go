package transactions

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
)

func (s *service) Create(ctx context.Context, newTransaction *core.TransactionItem) error {

	// TransactionID is always TransactionID like in Ethereum
	_, err := s.store.Insert(ctx, newTransaction)

	err = s.notificationService.SendCode(ctx, newTransaction.Owner, "New Transaction"+string(newTransaction.To))
	return err
}
