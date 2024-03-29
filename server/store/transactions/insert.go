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
	"github.com/MikaelLazarev/cerberus/server/store/helpers"
)

func (s *store) Insert(ctx context.Context, transaction *core.TransactionItem) (transactionID core.ID, err error) {

	result, err := s.Col.InsertOne(context.Background(), transaction)
	_, err = helpers.ConvertOIDToString(result.InsertedID)

	return transaction.ID, err
}
