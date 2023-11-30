package sysinfo

type InfoService interface {
	CheckDb() (bool, error)
	SaveInfo() (*SysInfo, error)
	GetInfo() (*SysInfo, error)
}

type InfoServiceImpl struct {
	dao  InfoDao
	util InfoUtil
}

func NewInfoService() (InfoService, error) {
	daoImpl, error := NewInfoDao()
	if error != nil {
		return nil, error
	}
	return InfoServiceImpl{dao: daoImpl, util: InfoUtilImpl{}}, nil
}

func (service InfoServiceImpl) CheckDb() (bool, error) {
	return service.dao.CheckDb()
}

func (service InfoServiceImpl) SaveInfo() (*SysInfo, error) {
	info := service.util.Info()
	error := service.dao.SaveInfo(info)
	return info, error
}

func (service InfoServiceImpl) GetInfo() (*SysInfo, error) {
	return service.dao.GetInfo()
}
