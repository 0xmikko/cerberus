/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package auth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/MikaelLazarev/cerberus/server/handlers/errors"
	"github.com/MikaelLazarev/cerberus/server/middlewares"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {

	validDTO := &core.LoginDTO{
		Email:    "registered@dbs.com",
		Password: "ilovedbs",
	}

	validUser, err := Services.UserService.SignUp(context.TODO(), validDTO)
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		name          string
		loginDTO      *core.LoginDTO
		expectedCode  int
		errorResponse map[string]string
		user          *core.User
	}{
		// LOGIN REQUEST TESTS
		{
			name: "#1. Wrong credentials",
			loginDTO: &core.LoginDTO{
				Email:    "user1",
				Password: "jkjkj",
			},
			expectedCode:  http.StatusBadRequest,
			errorResponse: map[string]string{"error": errors.ErrorWrongRequest.Error()},
		},
		{
			name: "#2. Wrong user/password",
			loginDTO: &core.LoginDTO{
				Email:    "user1@dbs.com",
				Password: "jkjkj",
			},
			expectedCode:  http.StatusForbidden,
			errorResponse: map[string]string{"error": errors.ErrorUserNotFoundOrIncorrectPassword.Error()},
		},
		{
			name:         "#3. Correct login",
			loginDTO:     validDTO,
			expectedCode: http.StatusOK,
			user:         validUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			jsonUserDTO, err := json.Marshal(tt.loginDTO)
			if err != nil {
				log.Fatal("Cant Marshal loginDTO")
			}
			req, _ := http.NewRequest("POST", "/auth/login/", bytes.NewReader(jsonUserDTO))
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)

			if w.Code == http.StatusOK {
				type response struct {
					Access  string `json:"access"`
					Refresh string `json:"refresh"`
				}

				var tokenData response
				err := json.Unmarshal(w.Body.Bytes(), &tokenData)
				assert.NoError(t, err)

				valid, err := middlewares.ValidateToken(tokenData.Access)
				assert.NoError(t, err)

				// Set userId Variable to gin.Context
				userID, ok := valid.Claims.(jwt.MapClaims)["user_id"].(string)
				assert.Equal(t, true, ok)
				assert.Equal(t, tt.user.ID, core.ID(userID))

			} else {

				expectedResponseJSON, err := json.Marshal(tt.errorResponse)
				if err != nil {
					log.Fatal(err)
				}
				assert.Equal(t, string(expectedResponseJSON), w.Body.String())
			}

		})
	}
}
