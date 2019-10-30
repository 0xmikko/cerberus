package notifications

import (
	"context"
	"fmt"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/payload"
	"log"
)

func (s *service) SendCode(ctx context.Context, userID core.ID, message string) error {

	notification := &apns2.Notification{}

	notificationToken, err := s.store.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	for token, _ := range notificationToken.IOSTokens {

		if token == "" {
			continue
		}

		notification.DeviceToken = token
		notification.Topic = s.BundleID
		notification.Payload = payload.NewPayload().AlertBody(message).Badge(1)

		res, err := s.Client.Push(notification)

		if err != nil {
			log.Fatal("Error:", err)
		}

		fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)

	}

	return nil
}
