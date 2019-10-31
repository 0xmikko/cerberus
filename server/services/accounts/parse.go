package accounts

import (
	"context"
	"fmt"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

func parseEvent(vLog types.Log, store core.AccountsStore, transactionsService core.TransactionsService) {
	event := struct {
		Request [32]byte       `json:"request"`
		To      common.Address `json:"to"`
		Amount  *big.Int
	}{}

	log.Println(vLog.Data)
	err := contractAbi.Unpack(&event, "NewPaymentRegistered", vLog.Data)
	if err == nil {

		fmt.Println(string(event.Request[:]))      // foo
		fmt.Println(string(event.To.Hex()))        // bar
		fmt.Println(string(event.Amount.String())) // bar

		accountID := core.ID(vLog.Address.Hex())
		account, err := store.FindByID(context.TODO(), accountID)
		if err != nil {
			return
		}

		err = transactionsService.Create(context.TODO(), &core.TransactionItem{
			ID:        core.ID(event.To.Hex()),
			Owner:     account.Owner,
			AccountID: accountID,
			To:        core.ID(event.To.Hex()),
			Amount:    event.Amount.Int64(),
			Deadline:  time.Time{},
			State:     core.Undefined,
			Active:    false,
		})

	}
}
