package sysinfo

import (
	"go.uber.org/mock/gomock"
	"testing"
)

func TestServiceSaveInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	daoImpl := NewMockInfoDao(mockCtrl)

	daoImpl.EXPECT().SaveInfo(gomock.Any()).Return(nil)

	infoExample := getSysInfoExample()

	util := NewMockInfoUtil(mockCtrl)
	util.EXPECT().Info().Return(&infoExample)

	// Create a service without function NewInfoService because it is a test,
	// and it should use the mock dao.
	var service = InfoServiceImpl{dao: daoImpl, util: util}

	info, error := service.SaveInfo()
	if error != nil {
		t.Fatal("Error: Save Info ", error)
	}

	if *info != infoExample {
		t.Fatal("Error: The Info saved should be same than the info ", infoExample, info)
	}
}

func TestServiceGetInfo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	daoImpl := NewMockInfoDao(mockCtrl)

	var info = getSysInfoExample()

	daoImpl.EXPECT().GetInfo().Return(&info, nil)

	// Create a service without function NewInfoService because it is a test,
	// and it should use the mock dao.
	var service = InfoServiceImpl{dao: daoImpl}

	infoFromService, error := service.GetInfo()
	if error != nil {
		t.Fatal("Error: Save Info ", error)
	}

	if info != *infoFromService {
		t.Fatal("Error: The Info saved should be same than the info ", infoFromService, info)
	}
}
