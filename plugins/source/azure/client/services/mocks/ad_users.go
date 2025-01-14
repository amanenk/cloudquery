// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: ADUsersClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	graphrbac "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	gomock "github.com/golang/mock/gomock"
)

// MockADUsersClient is a mock of ADUsersClient interface.
type MockADUsersClient struct {
	ctrl     *gomock.Controller
	recorder *MockADUsersClientMockRecorder
}

// MockADUsersClientMockRecorder is the mock recorder for MockADUsersClient.
type MockADUsersClientMockRecorder struct {
	mock *MockADUsersClient
}

// NewMockADUsersClient creates a new mock instance.
func NewMockADUsersClient(ctrl *gomock.Controller) *MockADUsersClient {
	mock := &MockADUsersClient{ctrl: ctrl}
	mock.recorder = &MockADUsersClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockADUsersClient) EXPECT() *MockADUsersClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockADUsersClient) List(arg0 context.Context, arg1, arg2 string) (graphrbac.UserListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(graphrbac.UserListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockADUsersClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockADUsersClient)(nil).List), arg0, arg1, arg2)
}
