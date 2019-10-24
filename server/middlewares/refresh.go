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
	"context"
	"fmt"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func RefreshToken(refreshToken string) (map[string]string, error) {
	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return SigningKey, nil
	})

	if err != nil {
		return nil, errors.New("Refresh token problem")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		userId, ok := claims["user_id"].(string)

		if !ok {
			return nil, errors.New("Invalid token!")
		}

		user, err := userService.GetUserById(context.TODO(), core.ID(userId))
		if err != nil {
			return nil, err
		}
		return GenerateTokenPair(user)
	}
	return nil, errors.New("Refresh token problem")
}
