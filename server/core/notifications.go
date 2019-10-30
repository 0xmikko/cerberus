package core

import "context"

type (
	NotificationToken struct {
		UserID        ID              `json:"user_id"`
		IOSTokens     map[string]bool `json:"ios_tokens"`
		AndroidTokens map[string]bool `json:"ios_tokens"`
	}

	NotificationStore interface {
		// Stores account obj and return account ID
		FindByID(ctx context.Context, userID ID) (*NotificationToken, error)

		Update(ctx context.Context, tokens *NotificationToken) error
		AddIOSToken(ctx context.Context, userID ID, token string) error
	}

	NotificationService interface {
		AddIOSToken(ctx context.Context, userID ID, token string) error

		RemoveIOSToken(ctx context.Context, userID ID, token string) error

		SendCode(ctx context.Context, userID ID, code string) error
	}
)
