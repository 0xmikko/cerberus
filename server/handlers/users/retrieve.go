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

func RetrieveHandler(c *gin.Context) {

	userID, ok := c.Params.Get("id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Wrong parameter"})
		return
	}

	profile, err := userService.GetProfile(context.TODO(), core.ID(userID))
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// ToDo asking for all services

	c.JSON(http.StatusOK, profile)

}
