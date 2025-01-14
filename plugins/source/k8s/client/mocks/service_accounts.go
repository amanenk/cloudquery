// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/k8s/client (interfaces: ServiceAccountsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockServiceAccountsClient is a mock of ServiceAccountsClient interface.
type MockServiceAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountsClientMockRecorder
}

// MockServiceAccountsClientMockRecorder is the mock recorder for MockServiceAccountsClient.
type MockServiceAccountsClientMockRecorder struct {
	mock *MockServiceAccountsClient
}

// NewMockServiceAccountsClient creates a new mock instance.
func NewMockServiceAccountsClient(ctrl *gomock.Controller) *MockServiceAccountsClient {
	mock := &MockServiceAccountsClient{ctrl: ctrl}
	mock.recorder = &MockServiceAccountsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountsClient) EXPECT() *MockServiceAccountsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockServiceAccountsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ServiceAccountList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ServiceAccountList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceAccountsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServiceAccountsClient)(nil).List), arg0, arg1)
}
