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
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
