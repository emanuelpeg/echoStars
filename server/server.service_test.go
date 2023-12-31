package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestServiceUpsertServer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	daoImpl := NewMockServerDao(mockCtrl)

	// Create a service without function NewServerService because it is a test,
	// and it should use the mock dao.
	var service = ServerServiceImpl{dao: daoImpl}
	var server = getFakeServer("http://server-a.com")
	var serverReturn = getFakeServer("http://server-a.com")
	daoImpl.EXPECT().Upsert(server).Return(serverReturn, nil)

	if created, err := service.Upsert(server); err == nil {
		assert.Equal(t, server, created, "Test Server should equal expected just new server")
	} else {
		t.Fatal("Error: creating Server ", err)
	}

}

func TestServiceGetAll(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	daoImpl := NewMockServerDao(mockCtrl)

	var servers []*Server
	servers = append(servers, getFakeServer("http://server-a.com"))
	servers = append(servers, getFakeServer("http://server-b.com"))
	servers = append(servers, getFakeServer("http://server-c.com"))

	var service = ServerServiceImpl{dao: daoImpl}

	daoImpl.EXPECT().GetAll().Return(servers, nil)

	if allServersList, err := service.GetAll(); err == nil {
		assert.Equal(t, len(servers), len(allServersList), "The count of items does not match")
	} else {
		t.Fatal("Error: trying to get all list of servers", err)
	}

}

func getFakeServer(urlHealth string) *Server {
	mailTo := "someone@company.com"
	ip := "127.0.0.1"
	return &Server{
		Hostname:  "test",
		Ip:        &ip,
		UrlHealth: urlHealth,
		Status:    false,
		MailTo:    &mailTo,
		Frequency: 31,
	}
}
