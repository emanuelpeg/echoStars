package monitor

import (
	"echoStars/server"
	"fmt"
	"net/http"
	"time"
)

var serverService server.ServerService

func init() {
	serverService, _ = server.NewServerService()
}

func Start() error {
	if serverService == nil {
		return fmt.Errorf("error creating server service")
	}

	servers, err := serverService.GetAll()
	if err != nil {
		fmt.Println("Error getting servers: ", err)
		return err
	}

	if len(servers) == 0 {
		return fmt.Errorf("no servers to watch")
	}

	watchServers(servers)

	return nil
}

func watchServers(servers []*server.Server) {
	for _, thisServer := range servers {
		go func(thisServer *server.Server) {
			for range time.Tick(time.Millisecond * time.Duration(thisServer.Frequency)) {
				start := time.Now()
				res, err := http.Get(thisServer.UrlHealth)
				newStatus := server.Down
				if err != nil {
					fmt.Println("SERVER URL GET ERR: ", err)
				} else {
					fmt.Println("SERVER URL GET RES: ", res.Request.URL, res.StatusCode, time.Since(start))
					newStatus = server.Up
				}
				if thisServer.Status != newStatus {
					thisServer.Status = newStatus
					_, err := serverService.Upsert(thisServer)
					if err != nil {
						fmt.Println("Err updating server status: ", err)
					}

				}
			}
		}(thisServer)
	}
}
