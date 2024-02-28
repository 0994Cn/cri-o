// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cri-o/cri-o/internal/config/seccomp/seccompociartifact (interfaces: Impl)

// Package seccompociartifactmock is a generated GoMock package.
package seccompociartifactmock

import (
	context "context"
	reflect "reflect"

	ociartifact "github.com/cri-o/cri-o/internal/config/ociartifact"
	gomock "github.com/golang/mock/gomock"
)

// MockImpl is a mock of Impl interface.
type MockImpl struct {
	ctrl     *gomock.Controller
	recorder *MockImplMockRecorder
}

// MockImplMockRecorder is the mock recorder for MockImpl.
type MockImplMockRecorder struct {
	mock *MockImpl
}

// NewMockImpl creates a new mock instance.
func NewMockImpl(ctrl *gomock.Controller) *MockImpl {
	mock := &MockImpl{ctrl: ctrl}
	mock.recorder = &MockImplMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImpl) EXPECT() *MockImplMockRecorder {
	return m.recorder
}

// Pull mocks base method.
func (m *MockImpl) Pull(arg0 context.Context, arg1 string, arg2 *ociartifact.PullOptions) (*ociartifact.Artifact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull", arg0, arg1, arg2)
	ret0, _ := ret[0].(*ociartifact.Artifact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pull indicates an expected call of Pull.
func (mr *MockImplMockRecorder) Pull(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockImpl)(nil).Pull), arg0, arg1, arg2)
}
