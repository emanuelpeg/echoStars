package computer

import "log"

type ServerDaoInterface interface {
	GetAll() ([]Server, error)
	Create(server *Server) error
	Delete(server *Server) error
}

func FactoryDao(dbEngine string) ServerDaoInterface {
	var i ServerDaoInterface
	switch dbEngine {
	case "boltdb":
		i = ServerDaoBoltDb{}
	default:
		log.Fatalf("Engine %s not implemented", dbEngine)
		return nil
	}

	return i
}
