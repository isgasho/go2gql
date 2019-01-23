// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/EGT-Ukraine/go2gql/tests/dataloader/generated/clients/apis (interfaces: ItemsServiceClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	apis "github.com/EGT-Ukraine/go2gql/tests/dataloader/generated/clients/apis"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockItemsServiceClient is a mock of ItemsServiceClient interface
type MockItemsServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockItemsServiceClientMockRecorder
}

// MockItemsServiceClientMockRecorder is the mock recorder for MockItemsServiceClient
type MockItemsServiceClientMockRecorder struct {
	mock *MockItemsServiceClient
}

// NewMockItemsServiceClient creates a new mock instance
func NewMockItemsServiceClient(ctrl *gomock.Controller) *MockItemsServiceClient {
	mock := &MockItemsServiceClient{ctrl: ctrl}
	mock.recorder = &MockItemsServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockItemsServiceClient) EXPECT() *MockItemsServiceClientMockRecorder {
	return m.recorder
}

// GetOne mocks base method
func (m *MockItemsServiceClient) GetOne(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (*apis.Item, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOne", varargs...)
	ret0, _ := ret[0].(*apis.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOne indicates an expected call of GetOne
func (mr *MockItemsServiceClientMockRecorder) GetOne(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOne", reflect.TypeOf((*MockItemsServiceClient)(nil).GetOne), varargs...)
}

// List mocks base method
func (m *MockItemsServiceClient) List(arg0 context.Context, arg1 *empty.Empty, arg2 ...grpc.CallOption) (*apis.ItemListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*apis.ItemListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockItemsServiceClientMockRecorder) List(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockItemsServiceClient)(nil).List), varargs...)
}