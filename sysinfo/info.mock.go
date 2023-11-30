// Code generated by MockGen. DO NOT EDIT.
// Source: /home/emanuel/Projects/goserver/echoStars/sysinfo/info.go
//
// Generated by this command:
//
//	mockgen -source=/home/emanuel/Projects/goserver/echoStars/sysinfo/info.go -destination=/home/emanuel/Projects/goserver/echoStars/sysinfo/info.mock.go
//
// Package mock_sysinfo is a generated GoMock package.
package sysinfo

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockInfoUtil is a mock of InfoUtil interface.
type MockInfoUtil struct {
	ctrl     *gomock.Controller
	recorder *MockInfoUtilMockRecorder
}

// MockInfoUtilMockRecorder is the mock recorder for MockInfoUtil.
type MockInfoUtilMockRecorder struct {
	mock *MockInfoUtil
}

// NewMockInfoUtil creates a new mock instance.
func NewMockInfoUtil(ctrl *gomock.Controller) *MockInfoUtil {
	mock := &MockInfoUtil{ctrl: ctrl}
	mock.recorder = &MockInfoUtilMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfoUtil) EXPECT() *MockInfoUtilMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockInfoUtil) Info() *SysInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(*SysInfo)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockInfoUtilMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockInfoUtil)(nil).Info))
}

func getSysInfoExample() SysInfo {
	return SysInfo{
		Hostname:     "test",
		Platform:     "testOs",
		Uptime:       200,
		RAM:          2000,
		RAMAvailable: 1000,
		RAMFree:      500,
		Disk:         100,
	}
}
