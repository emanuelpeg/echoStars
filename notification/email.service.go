package notification

import (
	"echoStars/util"
	"gopkg.in/gomail.v2"
	"log"
)

type EmailServiceConfig struct {
	SmtpHost     string
	SmtpPort     int
	SmtpUsername string
	SmtpPassword string
}
type EmailService struct {
	config EmailServiceConfig
}

func newEmailService(config EmailServiceConfig) *EmailService {
	return &EmailService{config: config}
}

func loadEmailNotificationConfig() (EmailServiceConfig, error) {
	config := EmailServiceConfig{
		SmtpHost:     util.ApplicationConfig.GetString("notification.smtp.host"),
		SmtpPort:     util.ApplicationConfig.GetInt("notification.smtp.port"),
		SmtpUsername: util.ApplicationConfig.GetString("notification.smtp.username"),
		SmtpPassword: util.ApplicationConfig.GetString("notification.smtp.password"),
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
