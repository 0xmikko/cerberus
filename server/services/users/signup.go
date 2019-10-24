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
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/MikaelLazarev/cerberus/server/services/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

func (s service) SignUp(ctx context.Context, userDTO *core.SignUpDTO) (*core.User, error) {

	userDTO.Email = strings.ToLower(userDTO.Email)
	isExists, err := s.store.IsExists(userDTO.Email)

	if isExists {
		return nil, errors.ErrorUserWithUsernameExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), 1)
	if err != nil {
		return nil, err
	}

	newUser := &core.User{
		Username:  userDTO.Email,
		Email:     userDTO.Email,
		Password:  hashedPassword,
		LastLogin: time.Time{},
		Active:    false,
	}

	newID, err := s.store.Insert(newUser)
	if err != nil {
		return nil, err
	}

	return s.store.FindByID(newID)
}
