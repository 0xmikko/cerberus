package core

import (
	"context"
	"time"
)

const (
	Undefined = iota
	Cancelled
	Confirmed
)

type (
	TransactionItem struct {
		ID        ID        `json:"id"`
		Owner     ID        `json:"owner"`
		AccountID ID        `json:"from"`
		To        ID        `json:"to"`
		Amount    int64     `json:"amount"`
		Deadline  time.Time `json:"deadline"`
		State     int       `json:"state"`
		Active    bool      `json:"active"`
	}

	TransactionsStore interface {
		// Stores transaction obj and return transaction ID
		Insert(ctx context.Context, transaction *TransactionItem) (transactionID ID, err error)

		// List returns Transactions for particular user from database
		ListByUser(ctx context.Context, userID ID) (result []*TransactionItem, err error)

		FindByID(ctx context.Context, transactionID ID) (*TransactionItem, error)

		Update(ctx context.Context, transaction *TransactionItem) error
	}

	TransactionsService interface {
		Create(ctx context.Context, newTransaction *TransactionItem) error

		Confirm(ctx context.Context, transactionID, userID ID, status bool) error

		Retrieve(ctx context.Context, transactionID ID, userID ID) (*TransactionItem, error)

		GetState(ctx context.Context, transactionID ID) (bool, error)

		List(ctx context.Context, userID ID) ([]*TransactionItem, error)
	}
)
