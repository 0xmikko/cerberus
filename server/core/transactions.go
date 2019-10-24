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
		ID                ID    `json:"id"`
		UserID            ID    `json:"user_id"`
		State             int   `json:"state"`
		ConfirmationsSent int16 `json:"confirmations_sent"`
		Active            bool  `json:"active"`
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
