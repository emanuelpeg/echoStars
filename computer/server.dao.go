package computer

import "log"

type ServerDao interface {
	GetAll() ([]Server, error)
	Create(server *Server) error
}

func FactoryDao(e string) ServerDao {
	var i ServerDao
	switch e {
	case "boltdb":
		i = ServerImplBoltDb{}
	default:
		log.Fatalf("Engine %s not implemented", e)
		return nil
	}

	return i
}
