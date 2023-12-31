// Code generated by MockGen. DO NOT EDIT.
// Source: /home/emanuel/Projects/goserver/echoStars/sysinfo/info.service.go
//
// Generated by this command:
//
//	mockgen -source=/home/emanuel/Projects/goserver/echoStars/sysinfo/info.service.go -destination=/home/emanuel/Projects/goserver/echoStars/sysinfo/info.service.mock.go
//
// Package mock_sysinfo is a generated GoMock package.
package sysinfo

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInfoService is a mock of InfoService interface.
type MockInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockInfoServiceMockRecorder
}

// MockInfoServiceMockRecorder is the mock recorder for MockInfoService.
type MockInfoServiceMockRecorder struct {
	mock *MockInfoService
}

// NewMockInfoService creates a new mock instance.
func NewMockInfoService(ctrl *gomock.Controller) *MockInfoService {
	mock := &MockInfoService{ctrl: ctrl}
	mock.recorder = &MockInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfoService) EXPECT() *MockInfoServiceMockRecorder {
	return m.recorder
}

// CheckDb mocks base method.
func (m *MockInfoService) CheckDb() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckDb")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDb indicates an expected call of CheckDb.
func (mr *MockInfoServiceMockRecorder) CheckDb() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDb", reflect.TypeOf((*MockInfoService)(nil).CheckDb))
}

// GetInfo mocks base method.
func (m *MockInfoService) GetInfo() (*SysInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(*SysInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInfo indicates an expected call of GetInfo.
func (mr *MockInfoServiceMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockInfoService)(nil).GetInfo))
}

// SaveInfo mocks base method.
func (m *MockInfoService) SaveInfo() (*SysInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveInfo")
	ret0, _ := ret[0].(*SysInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveInfo indicates an expected call of SaveInfo.
func (mr *MockInfoServiceMockRecorder) SaveInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveInfo", reflect.TypeOf((*MockInfoService)(nil).SaveInfo))
}
