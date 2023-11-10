package server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewServerDao call NewServerDao and check that it doesn't return nil
func TestNewServerDao(t *testing.T) {
	_, err := NewServerDao("../database/config.json")

	if err != nil {
		t.Fatal("Error: the NewServerDao should return a dao", err)
	}
}

func TestCreateServer(t *testing.T) {
	var dao, err = NewServerDao("../database/config_test.json")

	if err != nil {
		t.Fatal("Error: the NewServerDao should return a dao", err)
	}

	email := "some-user@testserver.com"
	var server = Server{
		Hostname:  "test",
		Ip:        "127.0.10.0",
		UrlHealth: "http//otherhost:8080/health",
		Status:    true,
		MailTo:    &email,
		Frequency: 21,
	}

	var newServer = &Server{}
	newServer, err = dao.Create(&server)
	if err != nil {
		t.Fatal("Error: Create Server ", err)
	}

	assert.Equal(t, &server, newServer, "Test Server should equal expected just new Created server")
	_, err = dao.Delete(&server.UrlHealth)
}
func TestGetAllServer(t *testing.T) {
	var dao, err = NewServerDao("../database/config_test.json")
	if err != nil {
		t.Fatal("Error: the NewServerDao should return a dao", err)
	}

	// todo delete all before create or run on isolate test.
	var servers [3]*Server
	servers[0], err = dao.Create(getServerByUrl("http//hostone:8080/health"))
	servers[1], err = dao.Create(getServerByUrl("http//hosttwo:8080/health"))
	servers[2], err = dao.Create(getServerByUrl("http//hostthree:8080/health"))
	if err != nil {
		t.Fatal("Error: Create Server ", err)
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

	return &Server{
		Hostname:  "test",
		Ip:        "127.0.10.0",
		UrlHealth: urlHealth,
		Status:    true,
		MailTo:    &email,
		Frequency: 21,
	}
}
