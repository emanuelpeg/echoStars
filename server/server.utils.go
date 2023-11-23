package server

import (
	"echoStars/util"
	"fmt"
)

func Seed() error {
	fmt.Println("Seed Starts...")
	var servers []Server
	err := util.ReadJSONFile("server/seed.json", &servers)
	if err != nil {
		return err
	}

	service, _ := NewServerService()
	for _, server := range servers {
		_, err := service.Upsert(&server)
		if err != nil {
			return err
		}
	}

	createdServers, _ := service.GetAll()
	for _, server := range createdServers {
		fmt.Println("Created server: ", server)
	}

	return nil
}
