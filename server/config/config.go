/*
* Cerberus Wallet App
* Designed for CoinList ChainLink Hackathon
*
* https://github.com/MikaelLazarev/cerberus
*
* Copyright (c) 2019, Mikhail Lazarev
 */
package config

import "os"

// Define controller constants
const PaginationLimit int = 30
const PaginationOffset int = 10

const (
	DEV = iota
	TEST
	PROD
)

// Return environment
func getEnv() int {

	env := os.Getenv("ENV")
	switch env {
	case "TEST":
		return TEST
	case "PROD":
		return PROD
	}

	return DEV
}
