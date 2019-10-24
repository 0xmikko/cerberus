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
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/MikaelLazarev/cerberus/server/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	//"os"
)

// Sign up new user & returns new JWT token
func SignUpHandler(c *gin.Context) {

	signUpDTO := &core.SignUpDTO{}

	if err := c.ShouldBind(signUpDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Authentificate user
	user, err := userService.SignUp(context.Background(), signUpDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Generate new JWT Token
	token, err := middlewares.GenerateTokenPair(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, token)

}
