/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */

package accounts

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/MikaelLazarev/cerberus/server/store/helpers"
	"github.com/google/uuid"
)

func (s *store) Insert(ctx context.Context, account *core.Account) (AccountID core.ID, err error) {
	account.ID = core.ID(uuid.New().String())

	result, err := s.Col.InsertOne(context.Background(), account)
	_, err = helpers.ConvertOIDToString(result.InsertedID)

	return account.ID, err
}
