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
	"github.com/MikaelLazarev/cerberus/server/handlers/errors"
	"github.com/MikaelLazarev/cerberus/server/middlewares"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	//"os"
)

// LoginHandler handler gets code and authorise app
// Returns new JWT token
func LoginHandler(c *gin.Context) {

	loginDTO := &core.LoginDTO{}

	if err := c.ShouldBindJSON(loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": errors.ErrorWrongRequest.Error()})
		log.Printf("%v#", err)
		return
	}

	// Authentificate user
	user, err := userService.Authenticate(context.Background(), loginDTO)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": errors.ErrorUserNotFoundOrIncorrectPassword.Error()})
		log.Printf("%v#", err)
		return
	}

	// Generate new JWT Token
	token, err := middlewares.GenerateTokenPair(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": errors.ErrorInternalServerError.Error()})
		log.Printf("%v#", err)
		return
	}

	c.JSON(http.StatusOK, token)

}
