package notification

import (
	"echoStars/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmailNotificationService(t *testing.T) {
	//Example of how different configuration files can be used for testing (still in progress)
	util.LoadConfig("../test")
	emailNotificationService := NewEmailNotificationService()

	assert.NotNil(t, emailNotificationService, "Expected a non-nil NotificationService")
}
