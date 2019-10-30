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
var notificationService core.NotificationService

func RegisterController(r gin.IRouter, us core.UserService, ns core.NotificationService) {

	userService = us
	notificationService = ns

	c := r.Group("/user")
	c.GET("/", ProfileHandler)
	c.PUT("/", UpdateHandler)
	c.POST("/register/ios/", RegisterIOSHandler)

}
