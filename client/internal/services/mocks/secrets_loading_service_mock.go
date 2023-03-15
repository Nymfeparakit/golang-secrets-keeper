// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/client/internal/services (interfaces: SecretsLoadingService)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSecretsLoadingService is a mock of SecretsLoadingService interface.
type MockSecretsLoadingService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsLoadingServiceMockRecorder
}

// MockSecretsLoadingServiceMockRecorder is the mock recorder for MockSecretsLoadingService.
type MockSecretsLoadingServiceMockRecorder struct {
	mock *MockSecretsLoadingService
}

// NewMockSecretsLoadingService creates a new mock instance.
func NewMockSecretsLoadingService(ctrl *gomock.Controller) *MockSecretsLoadingService {
	mock := &MockSecretsLoadingService{ctrl: ctrl}
	mock.recorder = &MockSecretsLoadingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsLoadingService) EXPECT() *MockSecretsLoadingServiceMockRecorder {
	return m.recorder
}

// LoadSecrets mocks base method.
func (m *MockSecretsLoadingService) LoadSecrets(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadSecrets", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadSecrets indicates an expected call of LoadSecrets.
func (mr *MockSecretsLoadingServiceMockRecorder) LoadSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadSecrets", reflect.TypeOf((*MockSecretsLoadingService)(nil).LoadSecrets), arg0)
}