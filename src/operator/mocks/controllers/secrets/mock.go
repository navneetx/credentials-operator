// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/otterize/spire-integration-operator/src/controllers (interfaces: SecretsManager)

// Package mock_secrets is a generated GoMock package.
package mock_secrets

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	secretstypes "github.com/otterize/spire-integration-operator/src/controllers/secrets/types"
	v1 "k8s.io/api/core/v1"
)

// MockSecretsManager is a mock of SecretsManager interface.
type MockSecretsManager struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerMockRecorder
}

// MockSecretsManagerMockRecorder is the mock recorder for MockSecretsManager.
type MockSecretsManagerMockRecorder struct {
	mock *MockSecretsManager
}

// NewMockSecretsManager creates a new mock instance.
func NewMockSecretsManager(ctrl *gomock.Controller) *MockSecretsManager {
	mock := &MockSecretsManager{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsManager) EXPECT() *MockSecretsManagerMockRecorder {
	return m.recorder
}

// EnsureTLSSecret mocks base method.
func (m *MockSecretsManager) EnsureTLSSecret(arg0 context.Context, arg1 secretstypes.SecretConfig, arg2 *v1.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureTLSSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureTLSSecret indicates an expected call of EnsureTLSSecret.
func (mr *MockSecretsManagerMockRecorder) EnsureTLSSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureTLSSecret", reflect.TypeOf((*MockSecretsManager)(nil).EnsureTLSSecret), arg0, arg1, arg2)
}

// RefreshTLSSecrets mocks base method.
func (m *MockSecretsManager) RefreshTLSSecrets(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTLSSecrets", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshTLSSecrets indicates an expected call of RefreshTLSSecrets.
func (mr *MockSecretsManagerMockRecorder) RefreshTLSSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTLSSecrets", reflect.TypeOf((*MockSecretsManager)(nil).RefreshTLSSecrets), arg0)
}
