package sysinfo

import (
	"echoStars/database"
	"testing"
)

func beforeEeach() {
	database.ConfigFileName = "../database/config_test.json"
}

// TestNewInfoDao call NewInfoDao and check that it doesn't return nil
func TestNewInfoDao(t *testing.T) {
	beforeEeach()
	_, error := NewInfoDao()

	if error != nil {
		t.Fatal("Error: the NewInfoDao should return a dao", error)
	}
}

func TestSaveInfo(t *testing.T) {
	beforeEeach()
	var dao, error = NewInfoDao()

	if error != nil {
		t.Fatal("Error: the NewInfoDao should return a dao", error)
	}

	var info = getSysInfoExample()

	error = dao.SaveInfo(&info)
	if error != nil {
		t.Fatal("Error: Save Info ", error)
	}

	infoSaved, error := dao.GetInfo()
	if error != nil {
		t.Fatal("Error: Get Info ", error)
	}

	if info != *infoSaved {
		t.Fatal("Error: The Info saved should be same than the info ", infoSaved, info)
	}
}
