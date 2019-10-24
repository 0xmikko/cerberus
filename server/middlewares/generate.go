package middlewares

import (
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateTokenPair(user *core.User) (map[string]string, error) {

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["admin"] = user.Admin
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["user_id"] = user.ID
	rtClaims["admin"] = user.Admin
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rtString, err := refreshToken.SignedString(SigningKey)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access":  tokenString,
		"refresh": rtString,
	}, err
}
