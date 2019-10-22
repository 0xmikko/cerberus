package core


type (
	Confirmation struct {
		ID                string `json:"id"`
		WalletID          string `json:"wallet_id"`
		UserID            string `json:"user_id"`
		State             int    `json:"state"`
		ConfirmationsSent int16  `json:"confirmations_sent"`
		Active            bool   `json:"active"`
	}
)

const (
	Undefined = iota
	PartiallyConfirmed
	Cancelled
	Confirmed
)
