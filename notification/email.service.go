package notification

import (
	"echoStars/util"
	"gopkg.in/gomail.v2"
	"log"
)

type EmailServiceConfig struct {
	SmtpHost     string `json:"smtp_host"`
	SmtpPort     int    `json:"smtp_port"`
	SmtpUsername string `json:"smtp_username"`
	SmtpPassword string `json:"smtp_password"`
}
type EmailService struct {
	config EmailServiceConfig
}

func newEmailService(config EmailServiceConfig) *EmailService {
	return &EmailService{config: config}
}

func loadEmailNotificationConfig() (EmailServiceConfig, error) {
	var config EmailServiceConfig

	err := util.ReadJSONFile("notification/config.local.json", &config)
	if err != nil {
		err = util.ReadJSONFile("notification/config.json", &config)
		if err != nil {
			return EmailServiceConfig{}, err
		}
	}

	return config, nil
}
func (e *EmailService) SendNotification(subject, body, recipient string) error {
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
