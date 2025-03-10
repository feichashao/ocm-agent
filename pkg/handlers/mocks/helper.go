// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/ocm-agent/pkg/handlers (interfaces: OCMClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/openshift/ocm-agent-operator/api/v1alpha1"
)

// MockOCMClient is a mock of OCMClient interface.
type MockOCMClient struct {
	ctrl     *gomock.Controller
	recorder *MockOCMClientMockRecorder
}

// MockOCMClientMockRecorder is the mock recorder for MockOCMClient.
type MockOCMClientMockRecorder struct {
	mock *MockOCMClient
}

// NewMockOCMClient creates a new mock instance.
func NewMockOCMClient(ctrl *gomock.Controller) *MockOCMClient {
	mock := &MockOCMClient{ctrl: ctrl}
	mock.recorder = &MockOCMClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOCMClient) EXPECT() *MockOCMClientMockRecorder {
	return m.recorder
}

// SendServiceLog mocks base method.
func (m *MockOCMClient) SendServiceLog(arg0, arg1, arg2, arg3 string, arg4 v1alpha1.NotificationSeverity, arg5 string, arg6 []v1alpha1.NotificationReferenceType, arg7 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendServiceLog", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendServiceLog indicates an expected call of SendServiceLog.
func (mr *MockOCMClientMockRecorder) SendServiceLog(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendServiceLog", reflect.TypeOf((*MockOCMClient)(nil).SendServiceLog), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}
