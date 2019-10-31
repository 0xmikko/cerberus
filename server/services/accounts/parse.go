package accounts

import (
	"context"
	"encoding/hex"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"strings"
	"time"
)

func parseEvent(vLog types.Log, store core.AccountsStore, transactionsService core.TransactionsService) {
	event := struct {
		Request [32]byte       `json:"request"`
		To      common.Address `json:"to"`
		Amount  *big.Int       `json:"amount"`
	}{}

	log.Println("Pure data:", vLog.Data)
	err := contractAbi.Unpack(&event, "NewPaymentRegistered", vLog.Data)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("EVENT CATCHED")
	log.Println(string(event.Request[:]))      // foo
	log.Println(string(event.To.Hex()))        // bar
	log.Println(string(event.Amount.String())) // bar

	newByte32 := make([]byte, 32)
	for num, b := range event.Request {
		newByte32[num] = b
	}

	requestID := core.ID(strings.ToLower(hex.EncodeToString(newByte32)))
	log.Println("RID:", requestID)

	accountID := core.ID(vLog.Address.Hex())
	account, err := store.FindByID(context.TODO(), accountID)
	if err != nil {
		return
	}

	err = transactionsService.Create(context.TODO(), &core.TransactionItem{
		ID:        requestID,
		Owner:     account.Owner,
		AccountID: accountID,
		To:        core.ID(event.To.Hex()),
		Amount:    event.Amount.Int64(),
		Deadline:  time.Time{},
		State:     core.Undefined,
		Active:    false,
	})

}
