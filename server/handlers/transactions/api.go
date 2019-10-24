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
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
)

var transactionsService core.TransactionsService

func RegisterController(r gin.IRouter, rs core.TransactionsService) {

	transactionsService = rs

	c := r.Group("/transactions")
	c.GET("/", ListHandler)
	c.GET("/:id/", RetrieveHandler)
	c.PUT("/:id/", UpdateHandler)

}
