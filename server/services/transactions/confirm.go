package transactions

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/MikaelLazarev/cerberus/server/services/errors"
)

func (s *service) Confirm(ctx context.Context, transactionID, userID core.ID, confirm bool) error {

	transaction, err := s.store.FindByID(ctx, transactionID)
	if err != nil {
		return errors.ErrorTransactionNotFound
	}

	if transaction.Owner != userID {
		return errors.ErrorHaveNoRights
	}

	if transaction.State == core.Cancelled {
		return errors.ErrorCantConfirmCancelled
	}

	newState := core.Cancelled
	if confirm == true {
		newState = core.Confirmed
	}

	transaction.State = newState
	err = s.store.Update(ctx, transaction)
	if err != nil {
		return errors.ErrorDBError
	}

	return nil

}
