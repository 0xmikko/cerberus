/*
* Cerberus Wallet App
* Designed for CoinList ChainLink Hackathon
*
* https://github.com/MikaelLazarev/cerberus
*
* Copyright (c) 2019, Mikhail Lazarev
 */
package main

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/config"
	"github.com/MikaelLazarev/cerberus/server/handlers"
	"github.com/MikaelLazarev/cerberus/server/services"
	"github.com/MikaelLazarev/cerberus/server/store"
	"github.com/MikaelLazarev/cerberus/server/store/helpers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {

	configType := config.GetConfigType()

	// Setup release mode for Gin in production
	if configType == config.PROD {
		gin.SetMode(gin.ReleaseMode)
	}

	// Getting database settings
	dbConfig, err := config.GetDBConfig(configType)
	if err != nil {
		log.Fatal("Can't get DB Settings! %s", err)
	}

	// Connecting Mongo DB and defer disconnecting
	mongoDB := helpers.Connect(dbConfig.String(), dbConfig.DBname)
	defer mongoDB.Client.Disconnect(context.Background())

	// Inject store
	store := store.InjectStore(mongoDB.DB)

	// Inject services
	services := services.InjectServices(*store)

	// Run server
	Port := "8080"
	if os.Getenv("PORT") != "" {
		Port = os.Getenv("PORT")
	}

	// Setup handlers
	handlers.StartServer(*services, Port)

}
