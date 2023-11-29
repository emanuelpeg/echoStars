package server

type ServerService interface {
	Upsert(server *Server) (*Server, error)
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

func (service ServerServiceImpl) Upsert(server *Server) (*Server, error) {
	newServer, err := service.dao.Upsert(server)
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
