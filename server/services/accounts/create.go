package accounts

import (
	"context"
	"fmt"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func (s *service) Create(ctx context.Context, userID core.ID, account string) error {

	_, err := s.store.Insert(ctx, &core.Account{
		ID:      core.ID(account),
		Address: account,
		Owner:   userID,
	})

	go func(account string) {
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

		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-logs:
				fmt.Println(vLog) // pointer to event log
			}
		}
	}(account)

	return err

}
