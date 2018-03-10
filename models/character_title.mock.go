// Code generated by MockGen. DO NOT EDIT.
// Source: character_title.go

// Package models is a generated GoMock package.
package models

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCharacterTitleStore is a mock of CharacterTitleStore interface
type MockCharacterTitleStore struct {
	ctrl     *gomock.Controller
	recorder *MockCharacterTitleStoreMockRecorder
}

// MockCharacterTitleStoreMockRecorder is the mock recorder for MockCharacterTitleStore
type MockCharacterTitleStoreMockRecorder struct {
	mock *MockCharacterTitleStore
}

// NewMockCharacterTitleStore creates a new mock instance
func NewMockCharacterTitleStore(ctrl *gomock.Controller) *MockCharacterTitleStore {
	mock := &MockCharacterTitleStore{ctrl: ctrl}
	mock.recorder = &MockCharacterTitleStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCharacterTitleStore) EXPECT() *MockCharacterTitleStoreMockRecorder {
	return m.recorder
}

// GetOrCreate mocks base method
func (m *MockCharacterTitleStore) GetOrCreate(title string) (*CharacterTitle, error) {
	ret := m.ctrl.Call(m, "GetOrCreate", title)
	ret0, _ := ret[0].(*CharacterTitle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrCreate indicates an expected call of GetOrCreate
func (mr *MockCharacterTitleStoreMockRecorder) GetOrCreate(title interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrCreate", reflect.TypeOf((*MockCharacterTitleStore)(nil).GetOrCreate), title)
}