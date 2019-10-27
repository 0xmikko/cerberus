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

	acc, err := s.accountService.Create(ctx, newID, "0xTESTTESTTESTTESTTESTTESTTESTTESTTESTTEST12")
	if err != nil {
		return nil, errors.ErrorCantCreateAccounts
	}

	transactions := []*core.TransactionItem{
		&core.TransactionItem{
			ID:        "242313123412414_1234_124",
			Owner:     newID,
			AccountID: acc,
			To:        "0x123456789012345678901234567890123456789012",
			Amount:    50000000000000,
			Deadline:  time.Now().Add(time.Hour),
			State:     core.Undefined,
			Active:    true,
		},
		&core.TransactionItem{
			ID:        "242313123412414_1234_124",
			Owner:     newID,
			AccountID: acc,
			To:        "0x123456789012345678901234567890123456789012",
			Amount:    60000000000000,
			Deadline:  time.Now().Add(15 * time.Minute),
			State:     core.Cancelled,
			Active:    true,
		},
		&core.TransactionItem{
			ID:        "242313123412414_1234_124",
			Owner:     newID,
			AccountID: acc,
			To:        "0x123456789012345678901234567890123456789012",
			Amount:    60000000000000,
			Deadline:  time.Now().Add(15 * time.Minute),
			State:     core.Confirmed,
			Active:    true,
		},
	}
	for _, t := range transactions {

		err = s.transactionsService.Create(ctx, t)
		if err != nil {
			return nil, errors.ErrorCantCreateAccounts
		}
	}

	return s.store.FindByID(newID)
}
