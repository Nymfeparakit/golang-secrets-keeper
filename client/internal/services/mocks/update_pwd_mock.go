// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/client/internal/services (interfaces: UpdateDeletePasswordService)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	dto "github.com/Nymfeparakit/gophkeeper/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockUpdateDeletePasswordService is a mock of UpdateDeletePasswordService interface.
type MockUpdateDeletePasswordService struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateDeletePasswordServiceMockRecorder
}

// MockUpdateDeletePasswordServiceMockRecorder is the mock recorder for MockUpdateDeletePasswordService.
type MockUpdateDeletePasswordServiceMockRecorder struct {
	mock *MockUpdateDeletePasswordService
}

// NewMockUpdateDeletePasswordService creates a new mock instance.
func NewMockUpdateDeletePasswordService(ctrl *gomock.Controller) *MockUpdateDeletePasswordService {
	mock := &MockUpdateDeletePasswordService{ctrl: ctrl}
	mock.recorder = &MockUpdateDeletePasswordServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateDeletePasswordService) EXPECT() *MockUpdateDeletePasswordServiceMockRecorder {
	return m.recorder
}

// DeleteLocalSecret mocks base method.
func (m *MockUpdateDeletePasswordService) DeleteLocalSecret(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLocalSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLocalSecret indicates an expected call of DeleteLocalSecret.
func (mr *MockUpdateDeletePasswordServiceMockRecorder) DeleteLocalSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLocalSecret", reflect.TypeOf((*MockUpdateDeletePasswordService)(nil).DeleteLocalSecret), arg0, arg1)
}

// DeleteRemoteSecret mocks base method.
func (m *MockUpdateDeletePasswordService) DeleteRemoteSecret(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRemoteSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRemoteSecret indicates an expected call of DeleteRemoteSecret.
func (mr *MockUpdateDeletePasswordServiceMockRecorder) DeleteRemoteSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRemoteSecret", reflect.TypeOf((*MockUpdateDeletePasswordService)(nil).DeleteRemoteSecret), arg0, arg1)
}

// UpdateLocalSecret mocks base method.
func (m *MockUpdateDeletePasswordService) UpdateLocalSecret(arg0 context.Context, arg1 dto.LoginPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLocalSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLocalSecret indicates an expected call of UpdateLocalSecret.
func (mr *MockUpdateDeletePasswordServiceMockRecorder) UpdateLocalSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLocalSecret", reflect.TypeOf((*MockUpdateDeletePasswordService)(nil).UpdateLocalSecret), arg0, arg1)
}

// UpdateRemoteSecret mocks base method.
func (m *MockUpdateDeletePasswordService) UpdateRemoteSecret(arg0 context.Context, arg1 dto.LoginPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRemoteSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRemoteSecret indicates an expected call of UpdateRemoteSecret.
func (mr *MockUpdateDeletePasswordServiceMockRecorder) UpdateRemoteSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemoteSecret", reflect.TypeOf((*MockUpdateDeletePasswordService)(nil).UpdateRemoteSecret), arg0, arg1)
}
