package notification

import "log"

type SMSNotification struct{}

func (s *SMSNotification) SendNotification(message, recipient string) error {
	// Implement the SMS notification logic here
	log.Printf("Sending SMS to %s\nMessage: %s\n", recipient, message)
	return nil
}
