// Code generated by MockGen. DO NOT EDIT.
// Source: level.go

// Package models is a generated GoMock package.
package models

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLevelStore is a mock of LevelStore interface
type MockLevelStore struct {
	ctrl     *gomock.Controller
	recorder *MockLevelStoreMockRecorder
}

// MockLevelStoreMockRecorder is the mock recorder for MockLevelStore
type MockLevelStoreMockRecorder struct {
	mock *MockLevelStore
}

// NewMockLevelStore creates a new mock instance
func NewMockLevelStore(ctrl *gomock.Controller) *MockLevelStore {
	mock := &MockLevelStore{ctrl: ctrl}
	mock.recorder = &MockLevelStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLevelStore) EXPECT() *MockLevelStoreMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockLevelStore) Get(cID int64, job Job) (*Level, error) {
	ret := m.ctrl.Call(m, "Get", cID, job)
	ret0, _ := ret[0].(*Level)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockLevelStoreMockRecorder) Get(cID, job interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLevelStore)(nil).Get), cID, job)
}

// Set mocks base method
func (m *MockLevelStore) Set(lvl *Level) error {
	ret := m.ctrl.Call(m, "Set", lvl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockLevelStoreMockRecorder) Set(lvl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockLevelStore)(nil).Set), lvl)
}
