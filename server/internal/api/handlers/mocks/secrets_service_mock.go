// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/server/internal/api/handlers (interfaces: SecretsService)

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	context "context"
	reflect "reflect"

	dto "github.com/Nymfeparakit/gophkeeper/dto"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretsService is a mock of SecretsService interface.
type MockSecretsService struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsServiceMockRecorder
}

// MockSecretsServiceMockRecorder is the mock recorder for MockSecretsService.
type MockSecretsServiceMockRecorder struct {
	mock *MockSecretsService
}

// NewMockSecretsService creates a new mock instance.
func NewMockSecretsService(ctrl *gomock.Controller) *MockSecretsService {
	mock := &MockSecretsService{ctrl: ctrl}
	mock.recorder = &MockSecretsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsService) EXPECT() *MockSecretsServiceMockRecorder {
	return m.recorder
}

// AddBinaryInfo mocks base method.
func (m *MockSecretsService) AddBinaryInfo(arg0 context.Context, arg1 *dto.BinaryInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddBinaryInfo", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBinaryInfo indicates an expected call of AddBinaryInfo.
func (mr *MockSecretsServiceMockRecorder) AddBinaryInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinaryInfo", reflect.TypeOf((*MockSecretsService)(nil).AddBinaryInfo), arg0, arg1)
}

// AddCardInfo mocks base method.
func (m *MockSecretsService) AddCardInfo(arg0 context.Context, arg1 *dto.CardInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCardInfo", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCardInfo indicates an expected call of AddCardInfo.
func (mr *MockSecretsServiceMockRecorder) AddCardInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCardInfo", reflect.TypeOf((*MockSecretsService)(nil).AddCardInfo), arg0, arg1)
}

// AddCredentials mocks base method.
func (m *MockSecretsService) AddCredentials(arg0 context.Context, arg1 *dto.LoginPassword) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCredentials", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCredentials indicates an expected call of AddCredentials.
func (mr *MockSecretsServiceMockRecorder) AddCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCredentials", reflect.TypeOf((*MockSecretsService)(nil).AddCredentials), arg0, arg1)
}

// AddTextInfo mocks base method.
func (m *MockSecretsService) AddTextInfo(arg0 context.Context, arg1 *dto.TextInfo) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTextInfo", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTextInfo indicates an expected call of AddTextInfo.
func (mr *MockSecretsServiceMockRecorder) AddTextInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTextInfo", reflect.TypeOf((*MockSecretsService)(nil).AddTextInfo), arg0, arg1)
}

// DeleteBinaryInfo mocks base method.
func (m *MockSecretsService) DeleteBinaryInfo(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBinaryInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBinaryInfo indicates an expected call of DeleteBinaryInfo.
func (mr *MockSecretsServiceMockRecorder) DeleteBinaryInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBinaryInfo", reflect.TypeOf((*MockSecretsService)(nil).DeleteBinaryInfo), arg0, arg1)
}

// DeleteCardInfo mocks base method.
func (m *MockSecretsService) DeleteCardInfo(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCardInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCardInfo indicates an expected call of DeleteCardInfo.
func (mr *MockSecretsServiceMockRecorder) DeleteCardInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCardInfo", reflect.TypeOf((*MockSecretsService)(nil).DeleteCardInfo), arg0, arg1)
}

// DeleteCredentials mocks base method.
func (m *MockSecretsService) DeleteCredentials(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCredentials", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCredentials indicates an expected call of DeleteCredentials.
func (mr *MockSecretsServiceMockRecorder) DeleteCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCredentials", reflect.TypeOf((*MockSecretsService)(nil).DeleteCredentials), arg0, arg1)
}

// DeleteTextInfo mocks base method.
func (m *MockSecretsService) DeleteTextInfo(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTextInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTextInfo indicates an expected call of DeleteTextInfo.
func (mr *MockSecretsServiceMockRecorder) DeleteTextInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTextInfo", reflect.TypeOf((*MockSecretsService)(nil).DeleteTextInfo), arg0, arg1)
}

// GetBinaryById mocks base method.
func (m *MockSecretsService) GetBinaryById(arg0 context.Context, arg1, arg2 string) (*dto.BinaryInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinaryById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.BinaryInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBinaryById indicates an expected call of GetBinaryById.
func (mr *MockSecretsServiceMockRecorder) GetBinaryById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinaryById", reflect.TypeOf((*MockSecretsService)(nil).GetBinaryById), arg0, arg1, arg2)
}

// GetCardById mocks base method.
func (m *MockSecretsService) GetCardById(arg0 context.Context, arg1, arg2 string) (*dto.CardInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCardById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.CardInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCardById indicates an expected call of GetCardById.
func (mr *MockSecretsServiceMockRecorder) GetCardById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardById", reflect.TypeOf((*MockSecretsService)(nil).GetCardById), arg0, arg1, arg2)
}

// GetCredentialsById mocks base method.
func (m *MockSecretsService) GetCredentialsById(arg0 context.Context, arg1, arg2 string) (*dto.LoginPassword, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCredentialsById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.LoginPassword)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsById indicates an expected call of GetCredentialsById.
func (mr *MockSecretsServiceMockRecorder) GetCredentialsById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsById", reflect.TypeOf((*MockSecretsService)(nil).GetCredentialsById), arg0, arg1, arg2)
}

// GetTextById mocks base method.
func (m *MockSecretsService) GetTextById(arg0 context.Context, arg1, arg2 string) (*dto.TextInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTextById", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.TextInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTextById indicates an expected call of GetTextById.
func (mr *MockSecretsServiceMockRecorder) GetTextById(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTextById", reflect.TypeOf((*MockSecretsService)(nil).GetTextById), arg0, arg1, arg2)
}

// ListSecrets mocks base method.
func (m *MockSecretsService) ListSecrets(arg0 context.Context, arg1 string) (dto.SecretsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0, arg1)
	ret0, _ := ret[0].(dto.SecretsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretsServiceMockRecorder) ListSecrets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretsService)(nil).ListSecrets), arg0, arg1)
}

// UpdateBinaryInfo mocks base method.
func (m *MockSecretsService) UpdateBinaryInfo(arg0 context.Context, arg1 *dto.BinaryInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBinaryInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBinaryInfo indicates an expected call of UpdateBinaryInfo.
func (mr *MockSecretsServiceMockRecorder) UpdateBinaryInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBinaryInfo", reflect.TypeOf((*MockSecretsService)(nil).UpdateBinaryInfo), arg0, arg1)
}

// UpdateCardInfo mocks base method.
func (m *MockSecretsService) UpdateCardInfo(arg0 context.Context, arg1 *dto.CardInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCardInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCardInfo indicates an expected call of UpdateCardInfo.
func (mr *MockSecretsServiceMockRecorder) UpdateCardInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCardInfo", reflect.TypeOf((*MockSecretsService)(nil).UpdateCardInfo), arg0, arg1)
}

// UpdateCredentials mocks base method.
func (m *MockSecretsService) UpdateCredentials(arg0 context.Context, arg1 *dto.LoginPassword) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCredentials", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCredentials indicates an expected call of UpdateCredentials.
func (mr *MockSecretsServiceMockRecorder) UpdateCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCredentials", reflect.TypeOf((*MockSecretsService)(nil).UpdateCredentials), arg0, arg1)
}

// UpdateTextInfo mocks base method.
func (m *MockSecretsService) UpdateTextInfo(arg0 context.Context, arg1 *dto.TextInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTextInfo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTextInfo indicates an expected call of UpdateTextInfo.
func (mr *MockSecretsServiceMockRecorder) UpdateTextInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTextInfo", reflect.TypeOf((*MockSecretsService)(nil).UpdateTextInfo), arg0, arg1)
}
