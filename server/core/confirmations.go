package core

type (
	Confirmation struct {
		ID       string `json:"id"`
		WalletID string `json:"wallet_id"`
		UserID   string `json:"user_id"`
	}
)
