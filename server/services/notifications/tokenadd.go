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
	"log"
)

func (s *service) AddIOSToken(ctx context.Context, userID core.ID, token string) error {
	notificationToken, err := s.store.FindByID(ctx, userID)

	log.Printf("%#v", notificationToken)

	if err != nil && err.Error() == "mongo: no documents in result" {
		err = nil

		initMap := make(map[string]bool)
		initMap[""] = true

		notificationToken = &core.NotificationToken{
			UserID:        userID,
			IOSTokens:     initMap,
			AndroidTokens: initMap,
		}

		err = s.store.Update(ctx, notificationToken)
		if err != nil {
			return nil
		}

	}

	if err != nil {
		return err
	}

	return s.store.AddIOSToken(ctx, userID, token)

}
