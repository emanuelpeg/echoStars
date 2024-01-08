package server

import (
	"echoStars/util"
	"fmt"
)

func Seed() error {
	fmt.Println("Server seed starts...")
	var servers []Server
	seedFileName := util.ApplicationConfig.GetString("database.seed.server.file.name")
	err := util.ReadJSONFile("server/"+seedFileName, &servers)
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
