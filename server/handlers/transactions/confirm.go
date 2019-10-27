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
	"log"
	"net/http"
)

func ConfirmHandler(c *gin.Context) {

	type Confirmation struct {
		Status bool `json:"status"`
	}

	userID := c.MustGet("userId").(core.ID)
	transactionID, ok := c.Params.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Wrong parameter"})
		return
	}

	log.Println("Confirm Request for ", transactionID)

	var confirmation Confirmation
	err := c.BindJSON(&confirmation)
	if err != nil {
		c.AbortWithStatusJSON(403, gin.H{"error": "Wrong parameters"})
		return
	}

	err = transactionsService.Confirm(context.TODO(), core.ID(transactionID), userID, confirmation.Status)
	if err != nil {
		c.AbortWithStatusJSON(403, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
