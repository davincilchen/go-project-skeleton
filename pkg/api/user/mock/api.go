// Code generated by MockGen. DO NOT EDIT.
// Source: api.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	user "github.com/alanchchen/go-project-skeleton/pkg/api/user"
	gomock "github.com/golang/mock/gomock"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockServiceClient is a mock of ServiceClient interface
type MockServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockServiceClientMockRecorder
}

// MockServiceClientMockRecorder is the mock recorder for MockServiceClient
type MockServiceClientMockRecorder struct {
	mock *MockServiceClient
}

// NewMockServiceClient creates a new mock instance
func NewMockServiceClient(ctrl *gomock.Controller) *MockServiceClient {
	mock := &MockServiceClient{ctrl: ctrl}
	mock.recorder = &MockServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceClient) EXPECT() *MockServiceClientMockRecorder {
	return m.recorder
}

// AddUser mocks base method
func (m *MockServiceClient) AddUser(ctx context.Context, in *user.AddUserRequest, opts ...grpc.CallOption) (*user.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddUser", varargs...)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser
func (mr *MockServiceClientMockRecorder) AddUser(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockServiceClient)(nil).AddUser), varargs...)
}

// FindUserById mocks base method
func (m *MockServiceClient) FindUserById(ctx context.Context, in *user.FindUserByIdRequest, opts ...grpc.CallOption) (*user.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindUserById", varargs...)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById
func (mr *MockServiceClientMockRecorder) FindUserById(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockServiceClient)(nil).FindUserById), varargs...)
}

// FindUserByName mocks base method
func (m *MockServiceClient) FindUserByName(ctx context.Context, in *user.FindUserByNameRequest, opts ...grpc.CallOption) (*user.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindUserByName", varargs...)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByName indicates an expected call of FindUserByName
func (mr *MockServiceClientMockRecorder) FindUserByName(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByName", reflect.TypeOf((*MockServiceClient)(nil).FindUserByName), varargs...)
}

// ListUsers mocks base method
func (m *MockServiceClient) ListUsers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*user.Users, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListUsers", varargs...)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockServiceClientMockRecorder) ListUsers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockServiceClient)(nil).ListUsers), varargs...)
}

// MockServiceServer is a mock of ServiceServer interface
type MockServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockServiceServerMockRecorder
}

// MockServiceServerMockRecorder is the mock recorder for MockServiceServer
type MockServiceServerMockRecorder struct {
	mock *MockServiceServer
}

// NewMockServiceServer creates a new mock instance
func NewMockServiceServer(ctrl *gomock.Controller) *MockServiceServer {
	mock := &MockServiceServer{ctrl: ctrl}
	mock.recorder = &MockServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServiceServer) EXPECT() *MockServiceServerMockRecorder {
	return m.recorder
}

// AddUser mocks base method
func (m *MockServiceServer) AddUser(arg0 context.Context, arg1 *user.AddUserRequest) (*user.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser
func (mr *MockServiceServerMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockServiceServer)(nil).AddUser), arg0, arg1)
}

// FindUserById mocks base method
func (m *MockServiceServer) FindUserById(arg0 context.Context, arg1 *user.FindUserByIdRequest) (*user.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", arg0, arg1)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById
func (mr *MockServiceServerMockRecorder) FindUserById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockServiceServer)(nil).FindUserById), arg0, arg1)
}

// FindUserByName mocks base method
func (m *MockServiceServer) FindUserByName(arg0 context.Context, arg1 *user.FindUserByNameRequest) (*user.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByName", arg0, arg1)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByName indicates an expected call of FindUserByName
func (mr *MockServiceServerMockRecorder) FindUserByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByName", reflect.TypeOf((*MockServiceServer)(nil).FindUserByName), arg0, arg1)
}

// ListUsers mocks base method
func (m *MockServiceServer) ListUsers(arg0 context.Context, arg1 *empty.Empty) (*user.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].(*user.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockServiceServerMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockServiceServer)(nil).ListUsers), arg0, arg1)
}