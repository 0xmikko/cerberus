/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */

package auth_test

import (
	"context"
	config2 "github.com/MikaelLazarev/cerberus/server/config"
	"github.com/MikaelLazarev/cerberus/server/handlers"
	"github.com/MikaelLazarev/cerberus/server/services"
	"github.com/MikaelLazarev/cerberus/server/store"
	"github.com/MikaelLazarev/cerberus/server/store/helpers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"testing"
)

var router *gin.Engine
var Store *store.GlobalStore
var Services *services.Services

func TestMain(m *testing.M) {
	// Getting database settings for test database
	dbConfig, err := config2.GetDBConfig(config2.TEST)
	if err != nil {
		log.Fatalf("Can't get DB Settings! %s", err)
	}

	// Connecting Mongo DB and defer disconnecting
	testDB := helpers.Connect(dbConfig.String(), dbConfig.DBname)
	defer testDB.Client.Disconnect(context.Background())

	err = testDB.DB.Collection("users").Drop(context.TODO())
	if err != nil {
		log.Fatal("Cant clear test DB", err)
	}

	// Inject store
	Store = store.InjectStore(testDB.DB)

	// Inject services
	Services = services.InjectServices(*Store)

	// Setup handlers
	router = handlers.StartServer(*Services)

	code := m.Run()
	os.Exit(code)
}
