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

func (s *service) Retrieve(ctx context.Context, userID, accountID core.ID) (*core.Account, error) {

	account, err := s.store.FindByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return account, nil

}
