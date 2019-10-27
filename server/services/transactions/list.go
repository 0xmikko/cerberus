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
	"github.com/MikaelLazarev/cerberus/server/core"
)

func (s *service) List(ctx context.Context, userID core.ID) ([]*core.TransactionItem, error) {

	transactions, err := s.store.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
