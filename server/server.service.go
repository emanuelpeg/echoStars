package server

type ServerService interface {
	Upsert(server *Server) (*Server, error)
	GetAll() ([]*Server, error)
	Delete(hostname *string) (bool, error)
}

type ServerServiceImpl struct {
	dao ServerDao
}

func NewServerService() (ServerService, error) {
	daoImpl, err := NewServerDao()
	if err != nil {
		return nil, err
	}
	return ServerServiceImpl{dao: daoImpl}, nil
}

func (service ServerServiceImpl) Upsert(server *Server) (*Server, error) {
	var err error
	if newServer, err := service.dao.Upsert(server); err == nil {
		return newServer, nil
	}
	return nil, err
}

func (service ServerServiceImpl) GetAll() ([]*Server, error) {
	return service.dao.GetAll()
}

func (service ServerServiceImpl) Delete(url *string) (bool, error) {
	return service.dao.Delete(url)
}
