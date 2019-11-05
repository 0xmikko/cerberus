/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package handlers

import (
	"github.com/MikaelLazarev/cerberus/server/config"
	"github.com/MikaelLazarev/cerberus/server/handlers/accounts"
	"github.com/MikaelLazarev/cerberus/server/handlers/auth"
	"github.com/MikaelLazarev/cerberus/server/handlers/transactions"
	"github.com/MikaelLazarev/cerberus/server/handlers/users"
	"github.com/MikaelLazarev/cerberus/server/middlewares"
	"github.com/MikaelLazarev/cerberus/server/services"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func StartServer(services services.Services, port string) {

	router := gin.Default()

	// CORS setup
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if config.GetConfigType() == config.PROD {
		securityConfig, err := config.GetSecurityConfig()
		if err != nil {
			log.Fatal("Can't get sentry DNS options!")
		}

		if err := sentry.Init(sentry.ClientOptions{
			Dsn: securityConfig.SentryDSN,
		}); err != nil {
			log.Fatal("Sentry initialization failed: %v\n", err)
		}

		router.Use(sentrygin.New(sentrygin.Options{}))
		log.Println("Sentry service was started")
	}

	// Auth handlers
	authRouter := router.Group("/auth")
	auth.RegisterController(authRouter, services.UserService)

	// Api handlers. Restricted area
	apiRouter := router.Group("/api")
	apiRouter.Use(middlewares.AuthHandler(services.UserService))
	accounts.RegisterController(apiRouter, services.AccountsService)
	transactions.RegisterController(apiRouter, services.TransactionsService)
	users.RegisterController(apiRouter, services.UserService, services.NotificationService)

	// External adapter check
	adapter := router.Group("/adapter")
	transactions.AdapterController(adapter)

	// Temporary added to serve landing page

	staticPath := "./landing"
	if config.GetConfigType() == config.PROD {
		staticPath = "/app/server/landing"
	}

	walletPath := "../web/build"
	if config.GetConfigType() == config.PROD {
		walletPath = "/app/web/build"
	}

	router.Use(static.Serve("/", static.LocalFile(staticPath, false)))
	router.Use(static.Serve("/wallet", static.LocalFile(walletPath, false)))
	router.NoRoute(func(c *gin.Context) {
		c.File(walletPath + "/index.html")
	})

	log.Println("Starting service at :" + port)
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
		return
	}

}
