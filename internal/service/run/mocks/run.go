// Code generated by MockGen. DO NOT EDIT.
// Source: run.go
//
// Generated by this command:
//
//	mockgen -source=run.go -destination=/Users/gda/Documents/github/qase-tms/qasectl/internal/service/run/mocks/run.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// Mockclient is a mock of client interface.
type Mockclient struct {
	ctrl     *gomock.Controller
	recorder *MockclientMockRecorder
}

// MockclientMockRecorder is the mock recorder for Mockclient.
type MockclientMockRecorder struct {
	mock *Mockclient
}

// NewMockclient creates a new mock instance.
func NewMockclient(ctrl *gomock.Controller) *Mockclient {
	mock := &Mockclient{ctrl: ctrl}
	mock.recorder = &MockclientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockclient) EXPECT() *MockclientMockRecorder {
	return m.recorder
}

// CompleteRun mocks base method.
func (m *Mockclient) CompleteRun(ctx context.Context, projectCode string, runId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteRun", ctx, projectCode, runId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompleteRun indicates an expected call of CompleteRun.
func (mr *MockclientMockRecorder) CompleteRun(ctx, projectCode, runId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteRun", reflect.TypeOf((*Mockclient)(nil).CompleteRun), ctx, projectCode, runId)
}

// CreateRun mocks base method.
func (m *Mockclient) CreateRun(ctx context.Context, projectCode, title, description, envSlug string, mileID, planID int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRun", ctx, projectCode, title, description, envSlug, mileID, planID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRun indicates an expected call of CreateRun.
func (mr *MockclientMockRecorder) CreateRun(ctx, projectCode, title, description, envSlug, mileID, planID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRun", reflect.TypeOf((*Mockclient)(nil).CreateRun), ctx, projectCode, title, description, envSlug, mileID, planID)
}
