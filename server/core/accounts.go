package core

import "context"

type (
	Account struct {
		ID      ID     `json:"id"`
		Address string `json:"address"`
		Owner   ID     `json:"owner"`
	}

	AccountsStore interface {
		// Stores account obj and return account ID
		Insert(ctx context.Context, account *Account) (accountID ID, err error)

		// List returns Accounts for particular user from database
		ListByUser(ctx context.Context, userID ID) (result []*Account, err error)

		ListAll(ctx context.Context) ([]*Account, error)

		FindByID(ctx context.Context, accountID ID) (*Account, error)
	}

	AccountsService interface {
		Create(ctx context.Context, userID ID, account string) error
		Retrieve(ctx context.Context, userID, accountID ID) (*Account, error)
		ListByUser(ctx context.Context, userID ID) (result []*Account, err error)
	}
)
