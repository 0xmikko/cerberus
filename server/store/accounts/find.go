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
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"go.mongodb.org/mongo-driver/bson"
)

// FindByIDWithMembers - finds account by its ID
func (s *store) FindByID(ctx context.Context, accountID core.ID) (*core.Account, error) {
	var foundAccount core.Account

	// Creating filter which find / create an item with the same meeting.ID
	filter := bson.D{{
		"id",
		bson.D{{
			"$eq",
			accountID,
		}},
	}}

	err := s.Col.FindOne(context.Background(), filter).Decode(&foundAccount)
	if err != nil {
		return nil, err
	}

	return &foundAccount, nil
}
