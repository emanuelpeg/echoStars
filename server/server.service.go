package server

type ServerService interface {
	Create(server *Server) (*Server, error)
	GetAll() ([]*Server, error)
	Delete(hostname *string) (bool, error)
}

type ServerServiceImpl struct {
	dao ServerDaoInterface
}

func NewServerService() ServerService {
	return ServerServiceImpl{dao: NewServerDao()}
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

func (service ServerServiceImpl) Delete(hostname *string) (bool, error) {
	return service.dao.Delete(hostname)
}
