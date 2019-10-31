/*
 * Cerberus Wallet App
 * Designed for CoinList ChainLink Hackathon
 *
 * https://github.com/MikaelLazarev/cerberus
 *
 * Copyright (c) 2019, Mikhail Lazarev
 */
package accounts

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/config"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type service struct {
	store               core.AccountsStore
	transactionsService core.TransactionsService
}

var contractAbi abi.ABI

func New(accountStore core.AccountsStore, ts core.TransactionsService) core.AccountsService {

	accounts, err := accountStore.ListAll(context.TODO())
	if err != nil {
		log.Fatal("Cant get full list of connected accounts")
	}

	configType := config.GetConfigType()

	var abiStr []byte

	if configType == config.PROD {
		abiStr = []byte(os.Getenv("ABI"))

	} else {
		abiStr, err = ioutil.ReadFile("./abi.json")
		if err != nil {
			log.Fatal("Cant read file ./abi.json")
		}
	}

	contractAbi, err = abi.JSON(strings.NewReader(string(abiStr)))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("AAA%#v", contractAbi)

	for _, acc := range accounts {
		go monitor(acc.Address, accountStore, ts)
	}

	cs := &service{accountStore, ts}
	return cs
}
