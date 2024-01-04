package monitor

import (
	"echoStars/notification"
	"echoStars/server"
	"echoStars/util"
	"fmt"
	"time"
)

type MonitorService interface {
	Start() error
}

type MonitorServiceImpl struct {
	serverService       server.ServerService
	notificationService notification.NotificationService
}

func NewMonitorService() (MonitorService, error) {
	serverServiceImpl, error := server.NewServerService()
	notificationServiceImpl := notification.NewEmailNotificationService()
	if error != nil {
		return nil, error
	}
	return &MonitorServiceImpl{serverService: serverServiceImpl, notificationService: notificationServiceImpl}, nil
}

func (service *MonitorServiceImpl) Start() error {
	servers, err := service.serverService.GetAll()
	if err != nil {
		fmt.Println("Error getting servers: ", err)
		return err
	}

	if len(servers) == 0 {
		return fmt.Errorf("no servers to watch")
	}

	service.watchServers(servers)

	return nil
}

func (service *MonitorServiceImpl) watchServers(servers []*server.Server) {
	for _, thisServer := range servers {
		go func(thisServer *server.Server) {
			for range time.Tick(time.Millisecond * time.Duration(thisServer.Frequency)) {
				start := time.Now()
				newStatus := service.serverService.HealthCheck(thisServer.UrlHealth)
				if !newStatus {
					fmt.Println("Server URL ERR", thisServer.UrlHealth, time.Since(start))
				} else {
					fmt.Println("Server URL OK: ", thisServer.UrlHealth, time.Since(start))
				}
				if thisServer.Status != newStatus {
					err := service.updateServerStatus(thisServer)
					if err == nil {
						service.notifyStatusChange(thisServer)
					}

				}
			}
		}(thisServer)
	}
}

func (service *MonitorServiceImpl) updateServerStatus(thisServer *server.Server) error {
	thisServer.Status = !thisServer.Status
	_, err := service.serverService.Upsert(thisServer)
	if err != nil {
		fmt.Println("Err updating server status: ", err)
	}
	return err
}

func (service *MonitorServiceImpl) notifyStatusChange(thisServer *server.Server) error {
	prettyStatus := "DOWN"
	if thisServer.Status {
		prettyStatus = "UP"
	}
	serverStatusChangeMsg := fmt.Sprintf("STATUS CHANGE. Server %s is: %v!", thisServer.Hostname, prettyStatus)
	fmt.Println(serverStatusChangeMsg)
	var recipient string
	if thisServer.MailTo != nil {
		recipient = *thisServer.MailTo
	} else {
		recipient = util.ApplicationConfig.GetString("notification.recipient")
	}
	if recipient != "" {
		err := service.notificationService.SendNotification("Server status changed", serverStatusChangeMsg, recipient)
		if err != nil {
			fmt.Println("Err sending status change notification: ", err)
		} else {
			fmt.Println("Notification sent to: ", recipient)
		}
	} else {
		return fmt.Errorf("no recipients found")
	}

	return nil
}
