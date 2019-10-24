/*
* Cerberus Wallet App
* Designed for CoinList ChainLink Hackathon
*
* https://github.com/MikaelLazarev/cerberus
*
* Copyright (c) 2019, Mikhail Lazarev
 */

package helpers

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertOIDToString(OID interface{}) (string, error) {
	if oid, ok := OID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}
	return "", errors.New("Error converting Mongo OID into string")
}
