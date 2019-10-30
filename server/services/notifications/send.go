package notifications

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MikaelLazarev/cerberus/server/core"
	"github.com/gin-gonic/gin"
	"github.com/sideshow/apns2"
	"log"
)

func (s *service) SendCode(ctx context.Context, userID core.ID, message string) error {

	notification := &apns2.Notification{}

	notificationToken, err := s.store.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	payload := gin.H{
		"aps": gin.H{
			"alert": message,
		}}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	for token, _ := range notificationToken.IOSTokens {

		if token == "" {
			continue
		}

		notification.DeviceToken = token
		notification.Topic = s.BundleID
		notification.Payload = jsonPayload

		res, err := s.Client.Push(notification)

		if err != nil {
			log.Fatal("Error:", err)
		}

		fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)

	}

	return nil
}
