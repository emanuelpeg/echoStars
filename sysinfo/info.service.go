package sysinfo

type InfoService interface {
	CheckDb() (bool, error)
	SaveInfo(*SysInfo) error
	GetInfo() (*SysInfo, error)
}

type InfoServiceImpl struct {
	dao InfoDao
}

func NewInfoService() (InfoService, error) {
	daoImpl, error := NewInfoDao()
	if error != nil {
		return nil, error
	}
	return InfoServiceImpl{dao: daoImpl}, nil
}

func (service InfoServiceImpl) CheckDb() (bool, error) {
	return service.dao.CheckDb()
}

func (service InfoServiceImpl) SaveInfo(info *SysInfo) error {
	error := service.dao.SaveInfo(info)
	return error
}

func (service InfoServiceImpl) GetInfo() (*SysInfo, error) {
	return service.dao.GetInfo()
}
