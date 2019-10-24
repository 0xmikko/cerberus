/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package transactions

import (
	"github.com/MikaelLazarev/cerberus/server/core"
)

type service struct {
	store core.TransactionsStore
}

func New(transactionStore core.TransactionsStore) core.TransactionsService {

	cs := &service{transactionStore}
	return cs
}
