// Code generated by MockGen. DO NOT EDIT.
// Source: api.pb.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	greeter "github.com/alanchchen/go-project-skeleton/pkg/api/greeter"
	gomock "github.com/golang/mock/gomock"
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

// SayHello mocks base method
func (m *MockServiceClient) SayHello(ctx context.Context, in *greeter.HelloRequest, opts ...grpc.CallOption) (*greeter.HelloReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SayHello", varargs...)
	ret0, _ := ret[0].(*greeter.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello
func (mr *MockServiceClientMockRecorder) SayHello(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockServiceClient)(nil).SayHello), varargs...)
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

// SayHello mocks base method
func (m *MockServiceServer) SayHello(arg0 context.Context, arg1 *greeter.HelloRequest) (*greeter.HelloReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SayHello", arg0, arg1)
	ret0, _ := ret[0].(*greeter.HelloReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SayHello indicates an expected call of SayHello
func (mr *MockServiceServerMockRecorder) SayHello(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockServiceServer)(nil).SayHello), arg0, arg1)
}