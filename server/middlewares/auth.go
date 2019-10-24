/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package middlewares

import (
	"fmt"
	config2 "github.com/MikaelLazarev/cerberus/server/config"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var SigningKey []byte
var userService core.UserService

func init() {
	securityConfig, err := config2.GetSecurityConfig()
	if err != nil {
		log.Fatal("Cant get Signing key")
	}
	SigningKey = []byte(securityConfig.SigningKey)
}

func AuthHandler(us core.UserService) gin.HandlerFunc {

	userService = us

	return func(c *gin.Context) {

		// Getting token from header
		token := c.Request.Header.Get("Authorization")

		// Check if toke in correct format
		// ie Bearer xx03xllasx
		const b = "Bearer "
		if !strings.Contains(token, b) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Your request is not authorized"})
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "An authorization token was not supplied"})
			return
		}
		// Validate token
		valid, err := ValidateToken(t[1])
		fmt.Println(t[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Invalid authorization token"})
			return
		}

		// Set userId Variable to gin.Context
		userID, ok := valid.Claims.(jwt.MapClaims)["user_id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Invalid authorization token"})
			return
		}

		// Set userId Variable to gin.Context
		admin, ok := valid.Claims.(jwt.MapClaims)["admin"].(bool)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Invalid authorization token"})
			return
		}

		c.Set("userId", core.ID(userID))
		c.Set("admin", admin)

		c.Next()
	}
}

func HasAdminRights() gin.HandlerFunc {
	return func(c *gin.Context) {
		admin := c.MustGet("admin").(bool)
		if admin {
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "You haven't admin access"})
	}
}
