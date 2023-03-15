// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/server/proto/secrets (interfaces: SecretsManagementClient)

// Package mock_secrets is a generated GoMock package.
package mock_secrets

import (
	context "context"
	reflect "reflect"

	secrets "github.com/Nymfeparakit/gophkeeper/server/proto/secrets"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockSecretsManagementClient is a mock of SecretsManagementClient interface.
type MockSecretsManagementClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagementClientMockRecorder
}

// MockSecretsManagementClientMockRecorder is the mock recorder for MockSecretsManagementClient.
type MockSecretsManagementClientMockRecorder struct {
	mock *MockSecretsManagementClient
}

// NewMockSecretsManagementClient creates a new mock instance.
func NewMockSecretsManagementClient(ctrl *gomock.Controller) *MockSecretsManagementClient {
	mock := &MockSecretsManagementClient{ctrl: ctrl}
	mock.recorder = &MockSecretsManagementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsManagementClient) EXPECT() *MockSecretsManagementClientMockRecorder {
	return m.recorder
}

// AddBinaryInfo mocks base method.
func (m *MockSecretsManagementClient) AddBinaryInfo(arg0 context.Context, arg1 *secrets.BinaryInfo, arg2 ...grpc.CallOption) (*secrets.AddResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddBinaryInfo", varargs...)
	ret0, _ := ret[0].(*secrets.AddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddBinaryInfo indicates an expected call of AddBinaryInfo.
func (mr *MockSecretsManagementClientMockRecorder) AddBinaryInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBinaryInfo", reflect.TypeOf((*MockSecretsManagementClient)(nil).AddBinaryInfo), varargs...)
}

// AddCardInfo mocks base method.
func (m *MockSecretsManagementClient) AddCardInfo(arg0 context.Context, arg1 *secrets.CardInfo, arg2 ...grpc.CallOption) (*secrets.AddResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCardInfo", varargs...)
	ret0, _ := ret[0].(*secrets.AddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCardInfo indicates an expected call of AddCardInfo.
func (mr *MockSecretsManagementClientMockRecorder) AddCardInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCardInfo", reflect.TypeOf((*MockSecretsManagementClient)(nil).AddCardInfo), varargs...)
}

// AddCredentials mocks base method.
func (m *MockSecretsManagementClient) AddCredentials(arg0 context.Context, arg1 *secrets.Password, arg2 ...grpc.CallOption) (*secrets.AddResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddCredentials", varargs...)
	ret0, _ := ret[0].(*secrets.AddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCredentials indicates an expected call of AddCredentials.
func (mr *MockSecretsManagementClientMockRecorder) AddCredentials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCredentials", reflect.TypeOf((*MockSecretsManagementClient)(nil).AddCredentials), varargs...)
}

// AddTextInfo mocks base method.
func (m *MockSecretsManagementClient) AddTextInfo(arg0 context.Context, arg1 *secrets.TextInfo, arg2 ...grpc.CallOption) (*secrets.AddResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddTextInfo", varargs...)
	ret0, _ := ret[0].(*secrets.AddResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTextInfo indicates an expected call of AddTextInfo.
func (mr *MockSecretsManagementClientMockRecorder) AddTextInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTextInfo", reflect.TypeOf((*MockSecretsManagementClient)(nil).AddTextInfo), varargs...)
}

// GetBinaryByID mocks base method.
func (m *MockSecretsManagementClient) GetBinaryByID(arg0 context.Context, arg1 *secrets.GetSecretRequest, arg2 ...grpc.CallOption) (*secrets.GetBinaryResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBinaryByID", varargs...)
	ret0, _ := ret[0].(*secrets.GetBinaryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBinaryByID indicates an expected call of GetBinaryByID.
func (mr *MockSecretsManagementClientMockRecorder) GetBinaryByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinaryByID", reflect.TypeOf((*MockSecretsManagementClient)(nil).GetBinaryByID), varargs...)
}

// GetCardByID mocks base method.
func (m *MockSecretsManagementClient) GetCardByID(arg0 context.Context, arg1 *secrets.GetSecretRequest, arg2 ...grpc.CallOption) (*secrets.GetCardResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCardByID", varargs...)
	ret0, _ := ret[0].(*secrets.GetCardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCardByID indicates an expected call of GetCardByID.
func (mr *MockSecretsManagementClientMockRecorder) GetCardByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCardByID", reflect.TypeOf((*MockSecretsManagementClient)(nil).GetCardByID), varargs...)
}

// GetCredentialsByID mocks base method.
func (m *MockSecretsManagementClient) GetCredentialsByID(arg0 context.Context, arg1 *secrets.GetSecretRequest, arg2 ...grpc.CallOption) (*secrets.GetCredentialsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCredentialsByID", varargs...)
	ret0, _ := ret[0].(*secrets.GetCredentialsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredentialsByID indicates an expected call of GetCredentialsByID.
func (mr *MockSecretsManagementClientMockRecorder) GetCredentialsByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredentialsByID", reflect.TypeOf((*MockSecretsManagementClient)(nil).GetCredentialsByID), varargs...)
}

// GetTextByID mocks base method.
func (m *MockSecretsManagementClient) GetTextByID(arg0 context.Context, arg1 *secrets.GetSecretRequest, arg2 ...grpc.CallOption) (*secrets.GetTextResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTextByID", varargs...)
	ret0, _ := ret[0].(*secrets.GetTextResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTextByID indicates an expected call of GetTextByID.
func (mr *MockSecretsManagementClientMockRecorder) GetTextByID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTextByID", reflect.TypeOf((*MockSecretsManagementClient)(nil).GetTextByID), varargs...)
}

// ListSecrets mocks base method.
func (m *MockSecretsManagementClient) ListSecrets(arg0 context.Context, arg1 *secrets.EmptyRequest, arg2 ...grpc.CallOption) (*secrets.ListSecretResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecrets", varargs...)
	ret0, _ := ret[0].(*secrets.ListSecretResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockSecretsManagementClientMockRecorder) ListSecrets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockSecretsManagementClient)(nil).ListSecrets), varargs...)
}

// UpdateBinaryInfo mocks base method.
func (m *MockSecretsManagementClient) UpdateBinaryInfo(arg0 context.Context, arg1 *secrets.BinaryInfo, arg2 ...grpc.CallOption) (*secrets.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBinaryInfo", varargs...)
	ret0, _ := ret[0].(*secrets.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBinaryInfo indicates an expected call of UpdateBinaryInfo.
func (mr *MockSecretsManagementClientMockRecorder) UpdateBinaryInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBinaryInfo", reflect.TypeOf((*MockSecretsManagementClient)(nil).UpdateBinaryInfo), varargs...)
}

// UpdateCardInfo mocks base method.
func (m *MockSecretsManagementClient) UpdateCardInfo(arg0 context.Context, arg1 *secrets.CardInfo, arg2 ...grpc.CallOption) (*secrets.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCardInfo", varargs...)
	ret0, _ := ret[0].(*secrets.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCardInfo indicates an expected call of UpdateCardInfo.
func (mr *MockSecretsManagementClientMockRecorder) UpdateCardInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCardInfo", reflect.TypeOf((*MockSecretsManagementClient)(nil).UpdateCardInfo), varargs...)
}

// UpdateCredentials mocks base method.
func (m *MockSecretsManagementClient) UpdateCredentials(arg0 context.Context, arg1 *secrets.Password, arg2 ...grpc.CallOption) (*secrets.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateCredentials", varargs...)
	ret0, _ := ret[0].(*secrets.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCredentials indicates an expected call of UpdateCredentials.
func (mr *MockSecretsManagementClientMockRecorder) UpdateCredentials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCredentials", reflect.TypeOf((*MockSecretsManagementClient)(nil).UpdateCredentials), varargs...)
}

// UpdateTextInfo mocks base method.
func (m *MockSecretsManagementClient) UpdateTextInfo(arg0 context.Context, arg1 *secrets.TextInfo, arg2 ...grpc.CallOption) (*secrets.EmptyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTextInfo", varargs...)
	ret0, _ := ret[0].(*secrets.EmptyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTextInfo indicates an expected call of UpdateTextInfo.
func (mr *MockSecretsManagementClientMockRecorder) UpdateTextInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTextInfo", reflect.TypeOf((*MockSecretsManagementClient)(nil).UpdateTextInfo), varargs...)
}
