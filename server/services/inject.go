/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package services

import (
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/MikaelLazarev/cerberus/server/services/accounts"
	"github.com/MikaelLazarev/cerberus/server/services/transactions"
	"github.com/MikaelLazarev/cerberus/server/services/users"
	"github.com/MikaelLazarev/cerberus/server/store"
)

type Services struct {
	TransactionsService core.TransactionsService
	AccountsService     core.AccountsService
	UserService         core.UserService
}

func InjectServices(Store store.GlobalStore) *Services {
	return &Services{
		transactions.New(Store.TransactionsStore),
		accounts.New(Store.AccountStore),
		users.New(Store.UserStore),
	}
}
