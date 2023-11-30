package notification

import (
	"echoStars/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

// commonSetupDone is used to ensure that common setup is performed only once.
var commonSetupDone bool

// commonSetup performs the common setup required for all tests.
func commonSetup() {
	if !commonSetupDone {
		util.LoadConfig("../test")
		commonSetupDone = true
	}
}

// TestMain is executed before any test in the package.
func TestMain(m *testing.M) {
	// Perform common setup before running all tests
	commonSetup()

	// Run all tests in the package
	m.Run()
}

func TestNewEmailNotificationService(t *testing.T) {
	emailNotificationService := NewEmailNotificationService()

	assert.NotNil(t, emailNotificationService, "Expected a non-nil NotificationService")
}

func TestLoadEmailNotificationConfig(t *testing.T) {
	config, err := loadEmailNotificationConfig()

	assert.NoError(t, err, "Expected no error loading email notification config")
	assert.NotEmpty(t, config.SmtpHost, "SMTP host should not be empty")
	assert.NotEqual(t, 0, config.SmtpPort, "SMTP port should not be zero")
	assert.NotEmpty(t, config.SmtpUsername, "SMTP username should not be empty")
	assert.NotEmpty(t, config.SmtpPassword, "SMTP password should not be empty")
}

func TestEmailServiceSendNotification(t *testing.T) {
	emailService := NewEmailNotificationService()

	subject := "Test Subject"
	body := "<p>This is a test email body</p>"
	recipient := "test@example.com"

	err := emailService.SendNotification(subject, body, recipient)

	assert.NoError(t, err, "Expected no error sending email notification")
}
