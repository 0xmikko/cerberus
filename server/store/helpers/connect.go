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
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Store struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect(uri string, dbname string) *Store {
	mongoClient := GetClient(uri)
	err := TestConnectionDB(mongoClient)
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	mongoDB := mongoClient.Database(dbname)
	return &Store{mongoClient, mongoDB}
}

func GetClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func TestConnectionDB(c *mongo.Client) error {
	err := c.Ping(context.Background(), readpref.Primary())
	return err

}
