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
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListHandler(c *gin.Context) {

	userID := c.MustGet("userId").(core.ID)

	goalsList, err := transactionsService.List(context.TODO(), userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, goalsList)

}
