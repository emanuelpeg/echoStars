package computer

import (
	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
)

type ServerStatus bool

const (
	Up   ServerStatus = true
	Down ServerStatus = false
)

type Server struct {
	Hostname  string       `json:"hostname"`
	Ip        string       `json:"ip"`
	UrlHealth string       `json:"url"`
	Status    ServerStatus `json:"status"`
	MailTo    *string      `json:"mailTo"`
	Frequency uint64       `json:"frequency"`
}

func Init(e *echo.Echo) {
	initRoutes(e)
}

// GetServers godoc
// @Summary Lists registered servers.
// @Description List all registered servers.
// @Tags server
// @Accept */*
// @Produce json
// @Success 200 {object} []computer.server
// @Router /servers [get]
func GetServers() ([]Server, error) {
	serverDao := FactoryDao("boltdb")
	servers, err := serverDao.GetAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return servers, nil
}

// CreateServer godoc
// @Summary Creates a new server.
// @Description Creates new server for monitoring using a server struct as input.
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} computer.server
// @Router /servers [post]
func CreateServer(server Server) error {
	serverDao := FactoryDao("boltdb")
	if serverDao != nil {
		err := serverDao.Create(&server)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

// DeleteServer godoc
// @Summary delete a server.
// @Description delete  server from server id.
// @Tags server
// @Accept json
// @Produce json
// @Success 204 {object} computer.server
// @Router /servers [post]
func DeleteServer(server Server) error {
	serverDao := FactoryDao("boltdb")
	if serverDao != nil {
		err := serverDao.Delete(&server)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}
