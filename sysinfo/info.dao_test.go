package sysinfo

import (
	"echoStars/util"
	"testing"
)

// commonSetupDone is used to ensure that common setup is performed only once.
var commonSetupDone bool

// commonSetup performs the common setup required for all tests.
func commonSetup() {
	if !commonSetupDone {
		util.LoadConfig("../test")
		commonSetupDone = true
	}
}

// TestMain is executed before any test in the package.
func TestMain(m *testing.M) {
	// Perform common setup before running all tests
	commonSetup()

	// Run all tests in the package
	m.Run()
}

// TestNewInfoDao call NewInfoDao and check that it doesn't return nil
func TestNewInfoDao(t *testing.T) {
	_, error := NewInfoDao()

	if error != nil {
		t.Fatal("Error: the NewInfoDao should return a dao", error)
	}
}

func TestSaveInfo(t *testing.T) {
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
