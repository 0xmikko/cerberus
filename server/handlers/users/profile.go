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
	"fmt"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProfileHandler(c *gin.Context) {

	userID := c.MustGet("userId").(core.ID)
	if userID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't get user id from request"})
	}

	profile, err := userService.GetProfile(context.TODO(), userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, profile)

}
