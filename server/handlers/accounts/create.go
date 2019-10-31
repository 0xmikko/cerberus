/*
* Cerberus Wallet App
* Designed for CoinList ChainLink Hackathon
*
* https://github.com/MikaelLazarev/cerberus
*
* Copyright (c) 2019, Mikhail Lazarev
 */
package accounts

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CreateAccountDTO struct {
	Account string `json:"account"`
}

func CreateHandler(c *gin.Context) {

	var accountDTO CreateAccountDTO

	userID := c.MustGet("userId").(core.ID)
	err := c.BindJSON(&accountDTO)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Wrong parameter"})
		return
	}

	err = accountService.Create(context.TODO(), userID, accountDTO.Account)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accountsList, err := accountService.ListByUser(context.TODO(), userID)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, accountsList)

}
