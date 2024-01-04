package monitor

import (
	"echoStars/notification"
	"echoStars/server"
	"echoStars/util"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var service MonitorService
var mockServerService *server.MockServerService
var mockNotificationService *notification.MockNotificationService

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

func createMonitorServiceInstance(t *testing.T) *gomock.Controller {
	util.LoadConfig("../dev")
	mockCtrl := gomock.NewController(t)
	mockServerService = server.NewMockServerService(mockCtrl)
	mockNotificationService = notification.NewMockNotificationService(mockCtrl)
	service = &MonitorServiceImpl{serverService: mockServerService, notificationService: mockNotificationService}

	return mockCtrl
}

func TestStartNoServers(t *testing.T) {
	mockCtrl := createMonitorServiceInstance(t)
	defer mockCtrl.Finish()
	expectedServers := make([]*server.Server, 0)
	mockServerService.EXPECT().GetAll().Return(expectedServers, nil)

	err := service.Start()

	assert.EqualError(t, err, "no servers to watch")
}

func TestStartErrorGettingServers(t *testing.T) {
	mockCtrl := createMonitorServiceInstance(t)
	defer mockCtrl.Finish()
	mockServerService.EXPECT().GetAll().Return(nil, fmt.Errorf("some error"))

	err := service.Start()

	assert.EqualError(t, err, "some error")
}

func TestStartServers(t *testing.T) {
	mockCtrl := createMonitorServiceInstance(t)
	defer mockCtrl.Finish()

	inputServers := getFakeServers()
	mockServerService.EXPECT().GetAll().Return(inputServers, nil)
	mockServerService.EXPECT().HealthCheck(inputServers[0].UrlHealth).Return(server.Up)
	mockServerService.EXPECT().HealthCheck(inputServers[1].UrlHealth).Times(2).Return(server.Up)
	mockServerService.EXPECT().HealthCheck(inputServers[2].UrlHealth).Return(server.Down)
	mockServerService.EXPECT().Upsert(inputServers[0]).Return(inputServers[0], nil)
	mockServerService.EXPECT().Upsert(inputServers[2]).Return(inputServers[2], nil)

	mockNotificationService.EXPECT().
		SendNotification("Server status changed", "STATUS CHANGE. Server www.downtoup.com is: UP!", "mailto@downtoup.com").
		Times(1).
		Return(nil)

	mockNotificationService.EXPECT().
		SendNotification("Server status changed", "STATUS CHANGE. Server www.uptodown.com is: DOWN!", "recipient.dev@recipient.com").
		Times(1).
		Return(nil)

	err := service.Start()

	// Allow some time for the goroutine to execute
	time.Sleep(time.Millisecond * time.Duration(1100))

	assert.NoError(t, err)
	assert.Equal(t, server.Up, inputServers[0].Status)
}

func getFakeServers() []*server.Server {
	ip := "127.0.0.1"
	mailTo1 := "mailto@downtoup.com"
	mailTo2 := "mailto@uptoup.com"
	servers := []*server.Server{
		{
			Hostname:  "www.downtoup.com",
			Ip:        &ip,
			UrlHealth: "http://downtoup.com/health",
			Status:    false,
			MailTo:    &mailTo1,
			Frequency: 700,
		},
		{
			Hostname:  "www.uptoup.com",
			Ip:        &ip,
			UrlHealth: "http://uptoup.com/health",
			Status:    true,
			MailTo:    &mailTo2,
			Frequency: 400,
		},
		{
			Hostname:  "www.uptodown.com",
			Ip:        &ip,
			UrlHealth: "http://uptodown.com/health",
			Status:    true,
			Frequency: 1000,
		},
	}
	return servers
}
