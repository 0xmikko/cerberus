/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package store

import (
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/MikaelLazarev/cerberus/server/store/accounts"
	"github.com/MikaelLazarev/cerberus/server/store/transactions"
	"github.com/MikaelLazarev/cerberus/server/store/users"
	"go.mongodb.org/mongo-driver/mongo"
)

// GlobalStore contains all repository objects for injecting
type GlobalStore struct {
	AccountStore      core.AccountsStore
	TransactionsStore core.TransactionsStore
	UserStore         core.UserStore
}

func InjectStore(db *mongo.Database) *GlobalStore {
	return &GlobalStore{
		accounts.New(db),
		transactions.New(db),
		users.New(db),
	}

}
