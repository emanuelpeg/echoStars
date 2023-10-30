package notification

import (
	"github.com/labstack/gommon/log"
)

type NotificationFactory interface {
	SendNotification(subject, body, recipient string) error
}

func NewEmailNotificationFactory() NotificationFactory {
	config, err := loadEmailNotificationConfig()
	if err != nil {
		log.Fatal("Error loading email notification configuration:", err)
	}
	return NewEmailNotification(config)
}
