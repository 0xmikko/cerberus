/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */

package notifications

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
)

func (s *service) RemoveIOSToken(ctx context.Context, userID core.ID, token string) error {

	//notificationToken, err := s.store.FindByID(ctx, userID)
	//if err != nil {
	//	return err
	//}
	//
	//newTokens := make([]string, len(notificationToken.IOSTokens))
	//
	//for _, t := range newTokens {
	//	if t != token {
	//		newTokens = append(newTokens, token)
	//	}
	//}
	//
	//notificationToken.IOSTokens = newTokens
	//
	//err = s.store.Update(ctx, notificationToken)
	//
	//return err
	return nil
}
