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
	"strings"
)

func (s service) Authenticate(ctx context.Context, userDTO *core.LoginDTO) (*core.User, error) {

	user, err := s.store.FindByUsername(strings.ToLower(userDTO.Email))
	if err != nil {
		return nil, errors.ErrorUserNotFound
	}

	if err != nil {
		return nil, errors.ErrorUserNotFound
	}

	return user, err
}
