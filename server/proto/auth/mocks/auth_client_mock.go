// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nymfeparakit/gophkeeper/server/proto/auth (interfaces: AuthManagementClient)

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	auth "github.com/Nymfeparakit/gophkeeper/server/proto/auth"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAuthManagementClient is a mock of AuthManagementClient interface.
type MockAuthManagementClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthManagementClientMockRecorder
}

// MockAuthManagementClientMockRecorder is the mock recorder for MockAuthManagementClient.
type MockAuthManagementClientMockRecorder struct {
	mock *MockAuthManagementClient
}

// NewMockAuthManagementClient creates a new mock instance.
func NewMockAuthManagementClient(ctrl *gomock.Controller) *MockAuthManagementClient {
	mock := &MockAuthManagementClient{ctrl: ctrl}
	mock.recorder = &MockAuthManagementClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthManagementClient) EXPECT() *MockAuthManagementClientMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthManagementClient) Login(arg0 context.Context, arg1 *auth.LoginRequest, arg2 ...grpc.CallOption) (*auth.LoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Login", varargs...)
	ret0, _ := ret[0].(*auth.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthManagementClientMockRecorder) Login(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthManagementClient)(nil).Login), varargs...)
}

// SignUp mocks base method.
func (m *MockAuthManagementClient) SignUp(arg0 context.Context, arg1 *auth.SignUpRequest, arg2 ...grpc.CallOption) (*auth.SignUpResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignUp", varargs...)
	ret0, _ := ret[0].(*auth.SignUpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthManagementClientMockRecorder) SignUp(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthManagementClient)(nil).SignUp), varargs...)
}
