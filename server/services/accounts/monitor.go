package accounts

import (
	"context"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func monitor(account string, store core.AccountsStore, ts core.TransactionsService) {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")

	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(account)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Monitoring ", account)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			parseEvent(vLog, store, ts) // pointer to event log
		}
	}
}
