package server

import (
	"echoStars/util"
	"testing"

	"github.com/stretchr/testify/assert"
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

// TestNewServerDao call NewServerDao and check that it doesn't return nil
func TestNewServerDao(t *testing.T) {
	_, err := NewServerDao()

	if err != nil {
		t.Fatal("Error: the NewServerDao should return a dao", err)
	}
}

func TestCreateServer(t *testing.T) {
	var dao, err = NewServerDao()

	if err != nil {
		t.Fatal("Error: the NewServerDao should return a dao", err)
	}

	email := "some-user@testserver.com"
	ip := "127.0.10.0"
	var server = Server{
		Hostname:  "test",
		Ip:        &ip,
		UrlHealth: "http//otherhost:8080/health",
		Status:    true,
		MailTo:    &email,
		Frequency: 21,
	}

	var newServer = &Server{}
	newServer, err = dao.Upsert(&server)
	if err != nil {
		t.Fatal("Error: Create Server ", err)
	}

	assert.Equal(t, &server, newServer, "Test Server should equal expected just new Created server")
	_, err = dao.Delete(&server.UrlHealth)
}

func TestGetAllServer(t *testing.T) {
	dao, err := NewServerDao()
	if err != nil {
		t.Fatal("Error: the NewServerDao should return a dao", err)
	}

	// todo delete all before create or run on isolate test.
	var servers [3]*Server
	servers[0], err = dao.Upsert(getServerByUrl("http//hostone:8080/health"))
	servers[1], err = dao.Upsert(getServerByUrl("http//hosttwo:8080/health"))
	servers[2], err = dao.Upsert(getServerByUrl("http//hostthree:8080/health"))
	if err != nil {
		t.Fatal("Error: Create Server %w", err)
	}

	all, err := dao.GetAll()
	if err != nil {
		return
	}
	if err != nil {
		t.Fatal("Error: Get all Servers ", err)
	}
	assert.Equal(t, len(servers), len(all), "The count of items does not match")

}

func getServerByUrl(urlHealth string) *Server {
	email := "some-user@testserver.com"
	ip := "127.0.10.0"

	return &Server{
		Hostname:  "test",
		Ip:        &ip,
		UrlHealth: urlHealth,
		Status:    true,
		MailTo:    &email,
		Frequency: 21,
	}
}
