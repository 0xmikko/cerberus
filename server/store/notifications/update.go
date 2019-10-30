/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */

package notifications

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// FindByIDWithMembers - finds account by its ID
func (s *store) Update(ctx context.Context, token *core.NotificationToken) error {
	filter := bson.D{{
		"userid", bson.D{{
			"$eq",
			token.UserID,
		}},
	}}

	update := bson.D{{
		"$set", token,
	}}

	upsert := true

	res, err := s.Col.UpdateOne(context.Background(), filter, update, &options.UpdateOptions{
		Upsert: &upsert,
	})

	if res.ModifiedCount == 0 {
		log.Printf("UserStore.Update: user %v wasn't modified. Check!", token)
	}

	return err

}
