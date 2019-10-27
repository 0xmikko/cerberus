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
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *store) Update(ctx context.Context, transaction *core.TransactionItem) error {
	filter := bson.D{{
		"id",
		bson.D{{
			"$eq",
			transaction.ID,
		}},
	}}

	update := bson.D{{
		"$set", transaction,
	}}
	_, err := s.Col.UpdateOne(context.Background(), filter, update)
	return err
}
