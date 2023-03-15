// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/client/internal/services (interfaces: LocalSecretsStorage)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	dto "github.com/Nymfeparakit/gophkeeper/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockLocalSecretsStorage is a mock of LocalSecretsStorage interface.
type MockLocalSecretsStorage struct {
	ctrl     *gomock.Controller
	recorder *MockLocalSecretsStorageMockRecorder
}

// MockLocalSecretsStorageMockRecorder is the mock recorder for MockLocalSecretsStorage.
type MockLocalSecretsStorageMockRecorder struct {
	mock *MockLocalSecretsStorage
}

// NewMockLocalSecretsStorage creates a new mock instance.
func NewMockLocalSecretsStorage(ctrl *gomock.Controller) *MockLocalSecretsStorage {
	mock := &MockLocalSecretsStorage{ctrl: ctrl}
	mock.recorder = &MockLocalSecretsStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLocalSecretsStorage) EXPECT() *MockLocalSecretsStorageMockRecorder {
	return m.recorder
}

// AddBinaryInfo mocks base method.
func (m *MockLocalSecretsStorage) AddBinaryInfo(arg0 context.Context, arg1 *dto.BinaryInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBinaryInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddBinaryInfo indicates an expected call of AddBinaryInfo.
func (mr *MockLocalSecretsStorageMockRecorder) AddBinaryInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinaryInfo", reflect.TypeOf((*MockLocalSecretsStorage)(nil).AddBinaryInfo), arg0, arg1)
}

// AddCardInfo mocks base method.
func (m *MockLocalSecretsStorage) AddCardInfo(arg0 context.Context, arg1 *dto.CardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCardInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCardInfo indicates an expected call of AddCardInfo.
func (mr *MockLocalSecretsStorageMockRecorder) AddCardInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCardInfo", reflect.TypeOf((*MockLocalSecretsStorage)(nil).AddCardInfo), arg0, arg1)
}

// AddCredentials mocks base method.
func (m *MockLocalSecretsStorage) AddCredentials(arg0 context.Context, arg1 *dto.LoginPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCredentials", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCredentials indicates an expected call of AddCredentials.
func (mr *MockLocalSecretsStorageMockRecorder) AddCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCredentials", reflect.TypeOf((*MockLocalSecretsStorage)(nil).AddCredentials), arg0, arg1)
}

// AddSecrets mocks base method.
func (m *MockLocalSecretsStorage) AddSecrets(arg0 context.Context, arg1 dto.SecretsList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSecrets", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSecrets indicates an expected call of AddSecrets.
func (mr *MockLocalSecretsStorageMockRecorder) AddSecrets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSecrets", reflect.TypeOf((*MockLocalSecretsStorage)(nil).AddSecrets), arg0, arg1)
}

// AddTextInfo mocks base method.
func (m *MockLocalSecretsStorage) AddTextInfo(arg0 context.Context, arg1 *dto.TextInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTextInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTextInfo indicates an expected call of AddTextInfo.
func (mr *MockLocalSecretsStorageMockRecorder) AddTextInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTextInfo", reflect.TypeOf((*MockLocalSecretsStorage)(nil).AddTextInfo), arg0, arg1)
}

// GetBinaryById mocks base method.
func (m *MockLocalSecretsStorage) GetBinaryById(arg0 context.Context, arg1, arg2 string) (dto.BinaryInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinaryById", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.BinaryInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBinaryById indicates an expected call of GetBinaryById.
func (mr *MockLocalSecretsStorageMockRecorder) GetBinaryById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinaryById", reflect.TypeOf((*MockLocalSecretsStorage)(nil).GetBinaryById), arg0, arg1, arg2)
}

// GetCardById mocks base method.
func (m *MockLocalSecretsStorage) GetCardById(arg0 context.Context, arg1, arg2 string) (dto.CardInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCardById", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.CardInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCardById indicates an expected call of GetCardById.
func (mr *MockLocalSecretsStorageMockRecorder) GetCardById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardById", reflect.TypeOf((*MockLocalSecretsStorage)(nil).GetCardById), arg0, arg1, arg2)
}

// GetCredentialsById mocks base method.
func (m *MockLocalSecretsStorage) GetCredentialsById(arg0 context.Context, arg1, arg2 string) (dto.LoginPassword, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsById", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.LoginPassword)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsById indicates an expected call of GetCredentialsById.
func (mr *MockLocalSecretsStorageMockRecorder) GetCredentialsById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsById", reflect.TypeOf((*MockLocalSecretsStorage)(nil).GetCredentialsById), arg0, arg1, arg2)
}

// GetTextById mocks base method.
func (m *MockLocalSecretsStorage) GetTextById(arg0 context.Context, arg1, arg2 string) (dto.TextInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTextById", arg0, arg1, arg2)
	ret0, _ := ret[0].(dto.TextInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTextById indicates an expected call of GetTextById.
func (mr *MockLocalSecretsStorageMockRecorder) GetTextById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTextById", reflect.TypeOf((*MockLocalSecretsStorage)(nil).GetTextById), arg0, arg1, arg2)
}

// ListSecrets mocks base method.
func (m *MockLocalSecretsStorage) ListSecrets(arg0 context.Context, arg1 string) (dto.SecretsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0, arg1)
	ret0, _ := ret[0].(dto.SecretsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockLocalSecretsStorageMockRecorder) ListSecrets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockLocalSecretsStorage)(nil).ListSecrets), arg0, arg1)
}

// UpdateBinaryInfo mocks base method.
func (m *MockLocalSecretsStorage) UpdateBinaryInfo(arg0 context.Context, arg1 *dto.BinaryInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBinaryInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBinaryInfo indicates an expected call of UpdateBinaryInfo.
func (mr *MockLocalSecretsStorageMockRecorder) UpdateBinaryInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBinaryInfo", reflect.TypeOf((*MockLocalSecretsStorage)(nil).UpdateBinaryInfo), arg0, arg1)
}

// UpdateCardInfo mocks base method.
func (m *MockLocalSecretsStorage) UpdateCardInfo(arg0 context.Context, arg1 *dto.CardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCardInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCardInfo indicates an expected call of UpdateCardInfo.
func (mr *MockLocalSecretsStorageMockRecorder) UpdateCardInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCardInfo", reflect.TypeOf((*MockLocalSecretsStorage)(nil).UpdateCardInfo), arg0, arg1)
}

// UpdateCredentials mocks base method.
func (m *MockLocalSecretsStorage) UpdateCredentials(arg0 context.Context, arg1 *dto.LoginPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCredentials", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCredentials indicates an expected call of UpdateCredentials.
func (mr *MockLocalSecretsStorageMockRecorder) UpdateCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCredentials", reflect.TypeOf((*MockLocalSecretsStorage)(nil).UpdateCredentials), arg0, arg1)
}

// UpdateTextInfo mocks base method.
func (m *MockLocalSecretsStorage) UpdateTextInfo(arg0 context.Context, arg1 *dto.TextInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTextInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTextInfo indicates an expected call of UpdateTextInfo.
func (mr *MockLocalSecretsStorageMockRecorder) UpdateTextInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTextInfo", reflect.TypeOf((*MockLocalSecretsStorage)(nil).UpdateTextInfo), arg0, arg1)
}