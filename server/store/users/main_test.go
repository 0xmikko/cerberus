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
	"github.com/MikaelLazarev/cerberus/server/store/helpers"
	"log"
	"os"
	"testing"
)

var testDB *helpers.Store
var s *store

func TestMain(m *testing.M) {
	// Getting database settings for test database
	dbConfig, err := config.GetDBConfig(config.TEST)
	if err != nil {
		log.Fatalf("Can't get DB Settings! %s", err)
	}

	// Connecting Mongo DB and defer disconnecting
	testDB = helpers.Connect(dbConfig.String(), dbConfig.DBname)
	defer testDB.Client.Disconnect(context.Background())

	// Generate User Stores
	s = &store{testDB.DB, testDB.DB.Collection(collection)}

	code := m.Run()
	os.Exit(code)
}
