// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/client/internal/services (interfaces: UpdateDeleteTextService)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	dto "github.com/Nymfeparakit/gophkeeper/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockUpdateDeleteTextService is a mock of UpdateDeleteTextService interface.
type MockUpdateDeleteTextService struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateDeleteTextServiceMockRecorder
}

// MockUpdateDeleteTextServiceMockRecorder is the mock recorder for MockUpdateDeleteTextService.
type MockUpdateDeleteTextServiceMockRecorder struct {
	mock *MockUpdateDeleteTextService
}

// NewMockUpdateDeleteTextService creates a new mock instance.
func NewMockUpdateDeleteTextService(ctrl *gomock.Controller) *MockUpdateDeleteTextService {
	mock := &MockUpdateDeleteTextService{ctrl: ctrl}
	mock.recorder = &MockUpdateDeleteTextServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateDeleteTextService) EXPECT() *MockUpdateDeleteTextServiceMockRecorder {
	return m.recorder
}

// DeleteLocalSecret mocks base method.
func (m *MockUpdateDeleteTextService) DeleteLocalSecret(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLocalSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLocalSecret indicates an expected call of DeleteLocalSecret.
func (mr *MockUpdateDeleteTextServiceMockRecorder) DeleteLocalSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLocalSecret", reflect.TypeOf((*MockUpdateDeleteTextService)(nil).DeleteLocalSecret), arg0, arg1)
}

// DeleteRemoteSecret mocks base method.
func (m *MockUpdateDeleteTextService) DeleteRemoteSecret(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRemoteSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRemoteSecret indicates an expected call of DeleteRemoteSecret.
func (mr *MockUpdateDeleteTextServiceMockRecorder) DeleteRemoteSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRemoteSecret", reflect.TypeOf((*MockUpdateDeleteTextService)(nil).DeleteRemoteSecret), arg0, arg1)
}

// UpdateLocalSecret mocks base method.
func (m *MockUpdateDeleteTextService) UpdateLocalSecret(arg0 context.Context, arg1 dto.TextInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLocalSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLocalSecret indicates an expected call of UpdateLocalSecret.
func (mr *MockUpdateDeleteTextServiceMockRecorder) UpdateLocalSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLocalSecret", reflect.TypeOf((*MockUpdateDeleteTextService)(nil).UpdateLocalSecret), arg0, arg1)
}

// UpdateRemoteSecret mocks base method.
func (m *MockUpdateDeleteTextService) UpdateRemoteSecret(arg0 context.Context, arg1 dto.TextInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRemoteSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRemoteSecret indicates an expected call of UpdateRemoteSecret.
func (mr *MockUpdateDeleteTextServiceMockRecorder) UpdateRemoteSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemoteSecret", reflect.TypeOf((*MockUpdateDeleteTextService)(nil).UpdateRemoteSecret), arg0, arg1)
}
