/*
* Cerberus Wallet App
* Designed for CoinList ChainLink Hackathon
*
* https://github.com/MikaelLazarev/cerberus
*
* Copyright (c) 2019, Mikhail Lazarev
 */

package users

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (s *store) Update(user *core.User) error {

	filter := bson.D{{
		"id", bson.D{{
			"$eq",
			user.ID,
		}},
	}}

	update := bson.D{{
		"$set", user,
	}}
	res, err := s.Col.UpdateOne(context.Background(), filter, update)
	if res.ModifiedCount == 0 {
		log.Printf("UserStore.Update: user %v wasn't modified. Check!", user)
	}
	return err
}
