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
	"github.com/MikaelLazarev/cerberus/server/core"
)

type service struct {
	store core.UserStore
}

func New(UserStore core.UserStore) core.UserService {

	cs := &service{UserStore}
	return cs
}
