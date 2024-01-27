// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/infraenv (interfaces: API)

// Package infraenv is a generated GoMock package.
package infraenv

import (
	context "context"
	reflect "reflect"

	strfmt "github.com/go-openapi/strfmt"
	gomock "github.com/golang/mock/gomock"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// DeleteOrphanInfraEnvs mocks base method.
func (m *MockAPI) DeleteOrphanInfraEnvs(arg0 context.Context, arg1 int, arg2 strfmt.DateTime) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrphanInfraEnvs", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrphanInfraEnvs indicates an expected call of DeleteOrphanInfraEnvs.
func (mr *MockAPIMockRecorder) DeleteOrphanInfraEnvs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrphanInfraEnvs", reflect.TypeOf((*MockAPI)(nil).DeleteOrphanInfraEnvs), arg0, arg1, arg2)
}

// DeregisterInfraEnv mocks base method.
func (m *MockAPI) DeregisterInfraEnv(arg0 context.Context, arg1 strfmt.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeregisterInfraEnv", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeregisterInfraEnv indicates an expected call of DeregisterInfraEnv.
func (mr *MockAPIMockRecorder) DeregisterInfraEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeregisterInfraEnv", reflect.TypeOf((*MockAPI)(nil).DeregisterInfraEnv), arg0, arg1)
}
