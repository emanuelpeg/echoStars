package sysinfo

type InfoService interface {
	CheckDb() (bool, error)
	SaveInfo() (*SysInfo, error)
	GetInfo() (*SysInfo, error)
}

type InfoServiceImpl struct {
	dao InfoDao
}

func NewInfoService() InfoService {
	daoImpl := NewInfoDao()
	return InfoServiceImpl{dao: daoImpl}
}

func (service InfoServiceImpl) CheckDb() (bool, error) {
	return service.dao.CheckDb()
}

func (service InfoServiceImpl) SaveInfo() (*SysInfo, error) {
	info := Info()
	error := service.dao.SaveInfo(info)
	return info, error
}

func (service InfoServiceImpl) GetInfo() (*SysInfo, error) {
	return service.dao.GetInfo()
}
