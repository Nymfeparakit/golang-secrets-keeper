// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/client/internal/services (interfaces: UpdatePasswordsService)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
)

// MockUpdatePasswordsService is a mock of UpdatePasswordsService interface.
type MockUpdatePasswordsService struct {
	ctrl     *gomock.Controller
	recorder *MockUpdatePasswordsServiceMockRecorder
}

// MockUpdatePasswordsServiceMockRecorder is the mock recorder for MockUpdatePasswordsService.
type MockUpdatePasswordsServiceMockRecorder struct {
	mock *MockUpdatePasswordsService
}

// NewMockUpdatePasswordsService creates a new mock instance.
func NewMockUpdatePasswordsService(ctrl *gomock.Controller) *MockUpdatePasswordsService {
	mock := &MockUpdatePasswordsService{ctrl: ctrl}
	mock.recorder = &MockUpdatePasswordsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdatePasswordsService) EXPECT() *MockUpdatePasswordsServiceMockRecorder {
	return m.recorder
}