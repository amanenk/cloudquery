// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: MySQLServersClient,MySQLConfigurationsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	gomock "github.com/golang/mock/gomock"
)

// MockMySQLServersClient is a mock of MySQLServersClient interface.
type MockMySQLServersClient struct {
	ctrl     *gomock.Controller
	recorder *MockMySQLServersClientMockRecorder
}

// MockMySQLServersClientMockRecorder is the mock recorder for MockMySQLServersClient.
type MockMySQLServersClientMockRecorder struct {
	mock *MockMySQLServersClient
}

// NewMockMySQLServersClient creates a new mock instance.
func NewMockMySQLServersClient(ctrl *gomock.Controller) *MockMySQLServersClient {
	mock := &MockMySQLServersClient{ctrl: ctrl}
	mock.recorder = &MockMySQLServersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMySQLServersClient) EXPECT() *MockMySQLServersClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockMySQLServersClient) List(arg0 context.Context) (mysql.ServerListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(mysql.ServerListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMySQLServersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMySQLServersClient)(nil).List), arg0)
}

// MockMySQLConfigurationsClient is a mock of MySQLConfigurationsClient interface.
type MockMySQLConfigurationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMySQLConfigurationsClientMockRecorder
}

// MockMySQLConfigurationsClientMockRecorder is the mock recorder for MockMySQLConfigurationsClient.
type MockMySQLConfigurationsClientMockRecorder struct {
	mock *MockMySQLConfigurationsClient
}

// NewMockMySQLConfigurationsClient creates a new mock instance.
func NewMockMySQLConfigurationsClient(ctrl *gomock.Controller) *MockMySQLConfigurationsClient {
	mock := &MockMySQLConfigurationsClient{ctrl: ctrl}
	mock.recorder = &MockMySQLConfigurationsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMySQLConfigurationsClient) EXPECT() *MockMySQLConfigurationsClientMockRecorder {
	return m.recorder
}

// ListByServer mocks base method.
func (m *MockMySQLConfigurationsClient) ListByServer(arg0 context.Context, arg1, arg2 string) (mysql.ConfigurationListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(mysql.ConfigurationListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByServer indicates an expected call of ListByServer.
func (mr *MockMySQLConfigurationsClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockMySQLConfigurationsClient)(nil).ListByServer), arg0, arg1, arg2)
}
