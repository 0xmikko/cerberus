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
	"os"
)

type SecurityConfig struct {
	SigningKey string `json:"signing_key"`
	SentryDSN  string `json:"sentry_dsn"`
}

func GetSecurityConfig() (*SecurityConfig, error) {

	switch getEnv() {
	default:
		w := &SecurityConfig{}
		err := LoadConfigFromJSON("./config/dev_security.json", w)
		return w, err

	case PROD:
		signingKey := os.Getenv("SIGNING_KEY")
		sentryDSN := os.Getenv("SENTRY_DSN")
		w := &SecurityConfig{signingKey, sentryDSN}
		return w, nil

	}
}
