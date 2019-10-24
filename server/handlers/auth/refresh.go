/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package auth

import (
	"github.com/MikaelLazarev/cerberus/server/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshTokenHandler(c *gin.Context) {
	type tokenReqBody struct {
		RefreshToken string `json:"refresh"`
	}
	tokenReq := tokenReqBody{}
	err := c.BindJSON(&tokenReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Cant extract refresh token"})
		return
	}

	tokens, err := middlewares.RefreshToken(tokenReq.RefreshToken)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Auth error"})
		return
	}

	c.JSON(http.StatusOK, tokens)
}
