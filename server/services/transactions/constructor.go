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
	store               core.TransactionsStore
	notificationService core.NotificationService
}

func New(transactionStore core.TransactionsStore, ns core.NotificationService) core.TransactionsService {

	cs := &service{
		store:               transactionStore,
		notificationService: ns,
	}

	return cs
}
