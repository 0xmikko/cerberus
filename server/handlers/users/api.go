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
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
)

var userService core.UserService

func RegisterController(r gin.IRouter, us core.UserService) {

	userService = us
	c := r.Group("/user")
	c.GET("/", ProfileHandler)
	c.PUT("/", UpdateHandler)

}
