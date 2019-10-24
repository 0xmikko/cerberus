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
	"github.com/MikaelLazarev/cerberus/server/core"
)

type service struct {
	store core.AccountsStore
}

func New(accountStore core.AccountsStore) core.AccountsService {

	cs := &service{accountStore}
	return cs
}
