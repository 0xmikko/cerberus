package core

import "context"

const (
	Undefined = iota
	PartiallyConfirmed
	Cancelled
	Confirmed
)

type (
	Transaction struct {
		ID     ID   `json:"id"`
		From   ID   `json:"from"`
		To     ID   `json:"to"`
		Amount int  `json:"amount"`
		State  int  `json:"state"`
		Active bool `json:"active"`
	}

	TransactionsStore interface {
		// Stores transaction obj and return transaction ID
		Insert(ctx context.Context, transaction *Transaction) (transactionID ID, err error)

		// List returns Transactions for particular user from database
		ListByUser(ctx context.Context, userID ID) (result []*Transaction, err error)

		FindByID(ctx context.Context, transactionID ID) (*Transaction, error)

		Update(ctx context.Context, transaction *Transaction) error
	}

	TransactionsService interface {
		Retrieve(ctx context.Context, transactionID ID, userID ID) (*Transaction, error)

		List(ctx context.Context, userID ID) ([]*Transaction, error)
	}
)
