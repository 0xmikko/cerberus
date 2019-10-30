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
	"github.com/MikaelLazarev/cerberus/server/config"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/sideshow/apns2"
)

type service struct {
	Client   *apns2.Client
	BundleID string
	store    core.NotificationStore
}

func New(ns core.NotificationStore) core.NotificationService {

	apnConfig, err := config.GetAPNConfig()
	if err != nil {
		return nil
	}

	client := apns2.NewTokenClient(&apnConfig.Token)

	return &service{
		Client:   client,
		BundleID: apnConfig.BundleID,
		store:    ns,
	}
}
