package notification

import (
	"github.com/labstack/gommon/log"
)

type NotificationService interface {
	SendNotification(subject, body, recipient string) error
}

func NewEmailNotificationService() NotificationService {
	config, err := loadEmailNotificationConfig()
	if err != nil {
		log.Fatal("Error loading email notification configuration:", err)
	}
	return newEmailService(config)
}

func NewSMSNotificationService() NotificationService {
	config, err := loadSMSServiceConfig()
	if err != nil {
		log.Fatal("Error loading email notification configuration:", err)
	}
	return newSMSService(config)
}
