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
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

func RegisterIOSHandler(c *gin.Context) {

	userID := c.MustGet("userId").(core.ID)
	if userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't get updatedProfile id from request"})
		return
	}

	var iosToken Token
	err := c.BindJSON(&iosToken)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't get updatedProfile id from request"})
		return
	}

	err = notificationService.AddIOSToken(context.TODO(), userID, iosToken.Token)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ok"})

}
