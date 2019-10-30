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
	"log"
)

// FindByIDWithMembers - finds account by its ID
func (s *store) FindByID(ctx context.Context, userID core.ID) (*core.NotificationToken, error) {
	var foundNotification core.NotificationToken

	// Creating filter which find / create an item with the same meeting.ID
	filter := bson.D{{
		"userid",
		bson.D{{
			"$eq",
			userID,
		}},
	}}

	err := s.Col.FindOne(ctx, filter).Decode(&foundNotification)
	if err != nil {
		return nil, err
	}

	log.Printf("%#v", foundNotification)

	return &foundNotification, nil
}
