package notification

import (
	"encoding/json"
	"gopkg.in/gomail.v2"
	"log"
	"os"
)

type EmailNotificationConfig struct {
	SmtpHost     string `json:"smtp_host"`
	SmtpPort     int    `json:"smtp_port"`
	SmtpUsername string `json:"smtp_username"`
	SmtpPassword string `json:"smtp_password"`
}
type EmailNotification struct {
	config EmailNotificationConfig
}

func NewEmailNotification(config EmailNotificationConfig) *EmailNotification {
	return &EmailNotification{config: config}
}

func loadEmailNotificationConfig() (EmailNotificationConfig, error) {
	configFile, err := os.Open("notification/config.json")
	if err != nil {
		return EmailNotificationConfig{}, err
	}
	defer configFile.Close()

	var config EmailNotificationConfig
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		return EmailNotificationConfig{}, err
	}

	return config, nil
}

func (e *EmailNotification) SendNotification(subject, body, recipient string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "your_email@example.com")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.config.SmtpHost, e.config.SmtpPort, e.config.SmtpUsername, e.config.SmtpPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email notification:", err)
		return err
	}

	return nil
}
