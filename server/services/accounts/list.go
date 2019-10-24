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
)

func (s *service) ListByUser(ctx context.Context, userID core.ID) (result []*core.Account, err error) {

	accounts, err := s.store.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}
