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

// FindByIDWithMembers - finds transaction by its ID
func (s *store) FindByID(ctx context.Context, transactionID core.ID) (*core.TransactionItem, error) {
	var foundTransaction core.TransactionItem

	// Creating filter which find / create an item with the same meeting.ID
	filter := bson.D{{
		"id",
		bson.D{{
			"$eq",
			transactionID,
		}},
	}}

	err := s.Col.FindOne(context.Background(), filter).Decode(&foundTransaction)
	if err != nil {
		return nil, err
	}

	return &foundTransaction, nil
}
