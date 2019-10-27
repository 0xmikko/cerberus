/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package transactions

import (
	"context"
	"errors"
	"github.com/MikaelLazarev/cerberus/server/core"
)

func (s *service) Retrieve(ctx context.Context, transactionID core.ID, userID core.ID) (*core.TransactionItem, error) {

	transaction, err := s.store.FindByID(ctx, transactionID)
	if err != nil {
		return nil, err
	}

	if transaction.Owner != userID {
		return nil, errors.New("You have no rights")
	}

	return transaction, nil

}
