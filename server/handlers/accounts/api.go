/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package accounts

import (
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
)

var accountService core.AccountsService

func RegisterController(r gin.IRouter, as core.AccountsService) {

	accountService = as

	c := r.Group("/accounts/")
	c.GET("/", ListHandler)
	c.GET("/:id/", RetrieveHandler)
	c.POST("/", CreateHandler)

}
