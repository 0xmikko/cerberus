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
	"go.mongodb.org/mongo-driver/mongo"
)

const collection = "accounts"

type store struct {
	DB  *mongo.Database
	Col *mongo.Collection
}

// NewStore - creates New store
func New(dbase *mongo.Database) core.AccountsStore {

	// Migrate model
	return &store{dbase, dbase.Collection(collection)}
}
