// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/otterize/spifferize/src/spireclient/entries (interfaces: Registry)

// Package mock_entries is a generated GoMock package.
package mock_entries

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// RegisterK8SPodEntry mocks base method.
func (m *MockRegistry) RegisterK8SPodEntry(arg0 context.Context, arg1, arg2, arg3 string, arg4 int32, arg5 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterK8SPodEntry", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterK8SPodEntry indicates an expected call of RegisterK8SPodEntry.
func (mr *MockRegistryMockRecorder) RegisterK8SPodEntry(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterK8SPodEntry", reflect.TypeOf((*MockRegistry)(nil).RegisterK8SPodEntry), arg0, arg1, arg2, arg3, arg4, arg5)
}