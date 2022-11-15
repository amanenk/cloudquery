// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/storage/v1 (interfaces: CSIStorageCapacitiesGetter,CSIStorageCapacityInterface)

// Package v1 is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/storage/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/storage/v1"
	v12 "k8s.io/client-go/kubernetes/typed/storage/v1"
)

// MockCSIStorageCapacitiesGetter is a mock of CSIStorageCapacitiesGetter interface.
type MockCSIStorageCapacitiesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCSIStorageCapacitiesGetterMockRecorder
}

// MockCSIStorageCapacitiesGetterMockRecorder is the mock recorder for MockCSIStorageCapacitiesGetter.
type MockCSIStorageCapacitiesGetterMockRecorder struct {
	mock *MockCSIStorageCapacitiesGetter
}

// NewMockCSIStorageCapacitiesGetter creates a new mock instance.
func NewMockCSIStorageCapacitiesGetter(ctrl *gomock.Controller) *MockCSIStorageCapacitiesGetter {
	mock := &MockCSIStorageCapacitiesGetter{ctrl: ctrl}
	mock.recorder = &MockCSIStorageCapacitiesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCSIStorageCapacitiesGetter) EXPECT() *MockCSIStorageCapacitiesGetterMockRecorder {
	return m.recorder
}

// CSIStorageCapacities mocks base method.
func (m *MockCSIStorageCapacitiesGetter) CSIStorageCapacities(arg0 string) v12.CSIStorageCapacityInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CSIStorageCapacities", arg0)
	ret0, _ := ret[0].(v12.CSIStorageCapacityInterface)
	return ret0
}

// CSIStorageCapacities indicates an expected call of CSIStorageCapacities.
func (mr *MockCSIStorageCapacitiesGetterMockRecorder) CSIStorageCapacities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CSIStorageCapacities", reflect.TypeOf((*MockCSIStorageCapacitiesGetter)(nil).CSIStorageCapacities), arg0)
}

// MockCSIStorageCapacityInterface is a mock of CSIStorageCapacityInterface interface.
type MockCSIStorageCapacityInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCSIStorageCapacityInterfaceMockRecorder
}

// MockCSIStorageCapacityInterfaceMockRecorder is the mock recorder for MockCSIStorageCapacityInterface.
type MockCSIStorageCapacityInterfaceMockRecorder struct {
	mock *MockCSIStorageCapacityInterface
}

// NewMockCSIStorageCapacityInterface creates a new mock instance.
func NewMockCSIStorageCapacityInterface(ctrl *gomock.Controller) *MockCSIStorageCapacityInterface {
	mock := &MockCSIStorageCapacityInterface{ctrl: ctrl}
	mock.recorder = &MockCSIStorageCapacityInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCSIStorageCapacityInterface) EXPECT() *MockCSIStorageCapacityInterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockCSIStorageCapacityInterface) Apply(arg0 context.Context, arg1 *v11.CSIStorageCapacityApplyConfiguration, arg2 v10.ApplyOptions) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockCSIStorageCapacityInterface) Create(arg0 context.Context, arg1 *v1.CSIStorageCapacity, arg2 v10.CreateOptions) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Create(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockCSIStorageCapacityInterface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockCSIStorageCapacityInterface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockCSIStorageCapacityInterface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockCSIStorageCapacityInterface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.CSIStorageCapacityList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.CSIStorageCapacityList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockCSIStorageCapacityInterface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 interface{}, arg5 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockCSIStorageCapacityInterface) Update(arg0 context.Context, arg1 *v1.CSIStorageCapacity, arg2 v10.UpdateOptions) (*v1.CSIStorageCapacity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CSIStorageCapacity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockCSIStorageCapacityInterface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockCSIStorageCapacityInterfaceMockRecorder) Watch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockCSIStorageCapacityInterface)(nil).Watch), arg0, arg1)
}
