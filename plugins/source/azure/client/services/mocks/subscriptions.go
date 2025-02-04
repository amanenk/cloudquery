// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cloudquery/plugins/source/azure/client/services (interfaces: SubscriptionsClient,TenantsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armsubscriptions "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	gomock "github.com/golang/mock/gomock"
)

// MockSubscriptionsClient is a mock of SubscriptionsClient interface.
type MockSubscriptionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionsClientMockRecorder
}

// MockSubscriptionsClientMockRecorder is the mock recorder for MockSubscriptionsClient.
type MockSubscriptionsClientMockRecorder struct {
	mock *MockSubscriptionsClient
}

// NewMockSubscriptionsClient creates a new mock instance.
func NewMockSubscriptionsClient(ctrl *gomock.Controller) *MockSubscriptionsClient {
	mock := &MockSubscriptionsClient{ctrl: ctrl}
	mock.recorder = &MockSubscriptionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionsClient) EXPECT() *MockSubscriptionsClientMockRecorder {
	return m.recorder
}

// NewListLocationsPager mocks base method.
func (m *MockSubscriptionsClient) NewListLocationsPager(arg0 string, arg1 *armsubscriptions.ClientListLocationsOptions) *runtime.Pager[armsubscriptions.ClientListLocationsResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListLocationsPager", arg0, arg1)
	ret0, _ := ret[0].(*runtime.Pager[armsubscriptions.ClientListLocationsResponse])
	return ret0
}

// NewListLocationsPager indicates an expected call of NewListLocationsPager.
func (mr *MockSubscriptionsClientMockRecorder) NewListLocationsPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListLocationsPager", reflect.TypeOf((*MockSubscriptionsClient)(nil).NewListLocationsPager), arg0, arg1)
}

// NewListPager mocks base method.
func (m *MockSubscriptionsClient) NewListPager(arg0 *armsubscriptions.ClientListOptions) *runtime.Pager[armsubscriptions.ClientListResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*runtime.Pager[armsubscriptions.ClientListResponse])
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockSubscriptionsClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockSubscriptionsClient)(nil).NewListPager), arg0)
}

// MockTenantsClient is a mock of TenantsClient interface.
type MockTenantsClient struct {
	ctrl     *gomock.Controller
	recorder *MockTenantsClientMockRecorder
}

// MockTenantsClientMockRecorder is the mock recorder for MockTenantsClient.
type MockTenantsClientMockRecorder struct {
	mock *MockTenantsClient
}

// NewMockTenantsClient creates a new mock instance.
func NewMockTenantsClient(ctrl *gomock.Controller) *MockTenantsClient {
	mock := &MockTenantsClient{ctrl: ctrl}
	mock.recorder = &MockTenantsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTenantsClient) EXPECT() *MockTenantsClientMockRecorder {
	return m.recorder
}

// NewListPager mocks base method.
func (m *MockTenantsClient) NewListPager(arg0 *armsubscriptions.TenantsClientListOptions) *runtime.Pager[armsubscriptions.TenantsClientListResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*runtime.Pager[armsubscriptions.TenantsClientListResponse])
	return ret0
}

// NewListPager indicates an expected call of NewListPager.
func (mr *MockTenantsClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockTenantsClient)(nil).NewListPager), arg0)
}
