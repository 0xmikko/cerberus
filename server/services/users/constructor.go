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
	accountService core.AccountsService
	transactionsService core.TransactionsService
}

func New(UserStore core.UserStore, as core.AccountsService, ts core.TransactionsService) core.UserService {

	cs := &service{
		store: UserStore,
		accountService: as,
		transactionsService: ts,
		}

	return cs
}
