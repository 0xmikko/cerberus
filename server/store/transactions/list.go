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

// Get teasers of all actual stories
func (s *store) ListByUser(ctx context.Context, userID core.ID) (result []*core.TransactionItem, err error) {

	results := make([]*core.TransactionItem, 0)

	filter := bson.D{{
		"owner",
		bson.D{{
			"$eq",
			userID,
		}},
	}}

	cur, err := s.Col.Find(context.TODO(), filter) // querying the "folders" collection
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem core.TransactionItem
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil

}
