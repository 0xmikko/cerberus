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

func GetConfigType() int {
	configEnv := os.Getenv("ENV")

	switch configEnv {
	case "TEST":
		return TEST

	case "PROD":
		return PROD

	}
	return DEV
}
