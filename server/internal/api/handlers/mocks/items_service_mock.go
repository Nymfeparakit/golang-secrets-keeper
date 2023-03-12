// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers (interfaces: ItemsService)

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	context "context"
	reflect "reflect"

	dto "github.com/Nymfeparakit/gophkeeper/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockItemsService is a mock of ItemsService interface.
type MockItemsService struct {
	ctrl     *gomock.Controller
	recorder *MockItemsServiceMockRecorder
}

// MockItemsServiceMockRecorder is the mock recorder for MockItemsService.
type MockItemsServiceMockRecorder struct {
	mock *MockItemsService
}

// NewMockItemsService creates a new mock instance.
func NewMockItemsService(ctrl *gomock.Controller) *MockItemsService {
	mock := &MockItemsService{ctrl: ctrl}
	mock.recorder = &MockItemsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemsService) EXPECT() *MockItemsServiceMockRecorder {
	return m.recorder
}

// AddCardInfo mocks base method.
func (m *MockItemsService) AddCardInfo(arg0 context.Context, arg1 *dto.CardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCardInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCardInfo indicates an expected call of AddCardInfo.
func (mr *MockItemsServiceMockRecorder) AddCardInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCardInfo", reflect.TypeOf((*MockItemsService)(nil).AddCardInfo), arg0, arg1)
}

// AddCredentials mocks base method.
func (m *MockItemsService) AddCredentials(arg0 context.Context, arg1 *dto.LoginPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCredentials", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPassword indicates an expected call of AddPassword.
func (mr *MockItemsServiceMockRecorder) AddPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCredentials", reflect.TypeOf((*MockItemsService)(nil).AddCredentials), arg0, arg1)
}

// AddTextInfo mocks base method.
func (m *MockItemsService) AddTextInfo(arg0 context.Context, arg1 *dto.TextInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTextInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTextInfo indicates an expected call of AddTextInfo.
func (mr *MockItemsServiceMockRecorder) AddTextInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTextInfo", reflect.TypeOf((*MockItemsService)(nil).AddTextInfo), arg0, arg1)
}

// ListSecrets mocks base method.
func (m *MockItemsService) ListSecrets(arg0 context.Context, arg1 string) (dto.SecretsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0, arg1)
	ret0, _ := ret[0].(dto.SecretsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListItems indicates an expected call of ListItems.
func (mr *MockItemsServiceMockRecorder) ListItems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockItemsService)(nil).ListSecrets), arg0, arg1)
}
