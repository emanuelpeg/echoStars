package notification

import "log"

type SMSService struct{}

func (s *SMSService) SendNotification(message, recipient string) error {
	// Implement the SMS notification logic here
	log.Printf("Sending SMS to %s\nMessage: %s\n", recipient, message)
	return nil
}
