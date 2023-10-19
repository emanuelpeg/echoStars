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

func GetServers() ([]Server, error) {
	serverDao := FactoryDao("boltdb")
	servers, err := serverDao.GetAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return servers, nil
}

func AddServer(server Server) error {
	serverDao := FactoryDao("boltdb")
	err := serverDao.Create(&server)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
