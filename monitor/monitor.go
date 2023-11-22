package monitor

import (
	"echoStars/server"
	"fmt"
	"net/http"
	"time"
)

func Start() error {
	serverService, err := server.NewServerService()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	servers, err := serverService.GetAll()
	if err != nil {
		fmt.Println("Error", err)
		return err
	}

	for _, server := range servers {
		go func(url string, freq uint64) {
			for range time.Tick(time.Millisecond * time.Duration(freq)) {
				start := time.Now()
				res, err := http.Get(url)
				if err != nil {
					fmt.Println("ERR: ", err)
				} else {
					fmt.Println("RES: ", res.Request.URL, res.StatusCode, time.Since(start))
				}
			}
		}(server.UrlHealth, server.Frequency)
	}

	return nil
}
