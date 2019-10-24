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
	"github.com/google/uuid"
)

func (s *store) Insert(user *core.User) (core.ID, error) {

	user.ID = core.ID(uuid.New().String())
	_, err := s.Col.InsertOne(context.Background(), user)

	return user.ID, err
}
