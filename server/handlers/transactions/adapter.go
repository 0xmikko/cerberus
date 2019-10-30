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
	"fmt"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdapterHandler(c *gin.Context) {

	transactionID, ok := c.Params.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Wrong parameter"})
		return
	}
	fmt.Println("Adapter confirmation Request for ", transactionID)

	confirmation, err := transactionsService.GetState(context.TODO(), core.ID(transactionID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"confirmation": confirmation})

}
