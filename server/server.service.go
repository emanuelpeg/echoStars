package server

import (
	"echoStars/util"
	"fmt"
)

type ServerService interface {
	Create(server *Server) (*Server, error)
	GetAll() ([]*Server, error)
	Delete(hostname *string) (bool, error)
}

type ServerServiceImpl struct {
	dao ServerDaoInterface
}

func NewServerService() (ServerService, error) {
	daoImpl, error := NewServerDao()
	if error != nil {
		return nil, error
	}
	return ServerServiceImpl{dao: daoImpl}, nil
}

func (service ServerServiceImpl) Create(server *Server) (*Server, error) {
	newServer, err := service.dao.Create(server)
	if err != nil {
		return nil, err
	}
	return newServer, nil
}

func (service ServerServiceImpl) GetAll() ([]*Server, error) {
	return service.dao.GetAll()
}

func (service ServerServiceImpl) Delete(url *string) (bool, error) {
	return service.dao.Delete(url)
}

func Seed() error {
	fmt.Println("Seed Starts...")
	var servers []Server
	err := util.ReadJSONFile("server/seed.json", &servers)
	if err != nil {
		return err
	}

	service, _ := NewServerService()
	for _, server := range servers {
		_, err := service.Create(&server)
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
