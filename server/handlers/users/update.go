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
	"github.com/gin-gonic/gin"
)

func UpdateHandler(c *gin.Context) {

	//userID := c.MustGet("userId").(core.ID)
	//if userID == "" {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't get updatedProfile id from request"})
	//	return
	//}
	//
	//var profileDTO core.ProfileDTO
	//err := c.BindJSON(&profileDTO)
	//if err != nil {
	//	log.Println(err)
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Can't get updatedProfile id from request"})
	//	return
	//}
	//
	//updatedProfile, err := userService.UpdateUserProfile(context.TODO(), userID, &profileDTO)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "User not found"})
	//	fmt.Println(err)
	//} else {
	//	c.JSON(http.StatusOK, updatedProfile)
	//}

}
