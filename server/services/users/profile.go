/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package users

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
)

func (s *service) GetProfile(ctx context.Context, UserId core.ID) (*core.User, error) {
	user, err := s.store.FindByID(UserId)
	if err != nil {
		return nil, err
	}

	return user, err
}
