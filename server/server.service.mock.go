// Code generated by MockGen. DO NOT EDIT.
// Source: C:\projects\echoStars\server\server.service.go
//
// Generated by this command:
//
//	mockgen -source=C:\projects\echoStars\server\server.service.go -destination=C:\projects\echoStars\server\server.service.mock.go -package=server
//

// Package server is a generated GoMock package.
package server

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockServerService is a mock of ServerService interface.
type MockServerService struct {
	ctrl     *gomock.Controller
	recorder *MockServerServiceMockRecorder
}

// MockServerServiceMockRecorder is the mock recorder for MockServerService.
type MockServerServiceMockRecorder struct {
	mock *MockServerService
}

// NewMockServerService creates a new mock instance.
func NewMockServerService(ctrl *gomock.Controller) *MockServerService {
	mock := &MockServerService{ctrl: ctrl}
	mock.recorder = &MockServerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerService) EXPECT() *MockServerServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockServerService) Delete(hostname *string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", hostname)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockServerServiceMockRecorder) Delete(hostname any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockServerService)(nil).Delete), hostname)
}

// GetAll mocks base method.
func (m *MockServerService) GetAll() ([]*Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockServerServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockServerService)(nil).GetAll))
}

// HealthCheck mocks base method.
func (m *MockServerService) HealthCheck(urlHealth string) ServerStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", urlHealth)
	ret0, _ := ret[0].(ServerStatus)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockServerServiceMockRecorder) HealthCheck(urlHealth any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockServerService)(nil).HealthCheck), urlHealth)
}

// Upsert mocks base method.
func (m *MockServerService) Upsert(server *Server) (*Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", server)
	ret0, _ := ret[0].(*Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *MockServerServiceMockRecorder) Upsert(server any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockServerService)(nil).Upsert), server)
}
