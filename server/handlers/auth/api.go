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
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
)

var userService core.UserService

func RegisterController(r gin.IRouter, us core.UserService) {

	userService = us

	r.POST("/login/", LoginHandler)
	r.POST("/signup/", SignUpHandler)
	r.POST("/token/refresh/", RefreshTokenHandler)

}
