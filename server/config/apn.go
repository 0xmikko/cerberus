/*
* Cerberus Wallet App
* Designed for CoinList ChainLink Hackathon
*
* https://github.com/MikaelLazarev/cerberus
*
* Copyright (c) 2019, Mikhail Lazarev
 */
package config

import (
	"errors"
	"github.com/sideshow/apns2/token"
	"log"
	"os"
)

type APNConfig struct {
	Token    token.Token `json:"-"`
	KeyID    string      `json:"key_id"`
	TeamID   string      `json:"team_id"`
	BundleID string      `json:"bundle_id"`
}

// Get Developers config from JSON files stored in config folder
func GetAPNConfig() (*APNConfig, error) {
	switch getEnv() {
	case DEV:
		return getAPNConfigFromJSON("./config/dev_apn.json")
	case TEST:
		return getAPNConfigFromJSON("./config/dev_apn.json")
	case PROD:
		return getAPNProdConfig()

	}
	return nil, errors.New("Wrong configuration type providedd")
}

func getAPNConfigFromJSON(fname string) (*APNConfig, error) {
	apnConfig := APNConfig{}
	err := LoadConfigFromJSON(fname, &apnConfig)
	if err != nil {
		log.Fatal("APN configuration  error:", err)
	}

	apnConfig.Token.AuthKey, err = token.AuthKeyFromFile("./config/AuthKey.p8")
	if err != nil {
		log.Fatal("APN configuration  error:", err)
	}

	apnConfig.Token.TeamID = apnConfig.TeamID
	apnConfig.Token.KeyID = apnConfig.KeyID

	return &apnConfig, err
}

func getAPNProdConfig() (*APNConfig, error) {

	log.Println("APN Prod configuration...")
	authKey, err := token.AuthKeyFromBytes([]byte("APN_AUTHKEY"))
	if err != nil {
		log.Fatal("token error:", err)
	}

	apnConfig := &APNConfig{
		Token: token.Token{
			AuthKey: authKey,
			KeyID:   os.Getenv("APN_KEYID"),
			TeamID:  os.Getenv("APN_TEAMID"),
		},
		KeyID:    os.Getenv("APN_KEYID"),
		TeamID:   os.Getenv("APN_TEAMID"),
		BundleID: os.Getenv("BUNDLE_ID"),
	}

	return apnConfig, nil
}
