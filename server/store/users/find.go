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
)

// FindByID - finds story by its ID
func (s *store) FindByID(id core.ID) (*core.User, error) {

	// Creating filter which find / create an item with the same meeting.ID
	filter := bson.D{{
		"id", id,
	}}

	return s.findByFilter(filter)
}

// FindByID - finds user by its ID
func (s *store) FindByUsername(username string) (*core.User, error) {

	// Creating filter which find / create an item with the same meeting.ID
	filter := bson.D{{
		"username", username,
	}}

	return s.findByFilter(filter)
}

func (s *store) findByFilter(filter interface{}) (*core.User, error) {
	var foundUser core.User

	err := s.Col.FindOne(context.Background(), filter).Decode(&foundUser)
	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}
