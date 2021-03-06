package models

import (
	gomock "github.com/golang/mock/gomock"
)

var _ DataStore = &MockDataStore{}

// MockDataStore is a more convenient DataStore used for mocking.
type MockDataStore struct {
	CharacterStore          *MockCharacterStore
	CharacterTombstoneStore *MockCharacterTombstoneStore
	CharacterTitleStore     *MockCharacterTitleStore
	LevelStore              *MockLevelStore
}

// NewMockDataStore creates a new DataStore, full of mock implementations of data stores.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	return &MockDataStore{
		CharacterStore:          NewMockCharacterStore(ctrl),
		CharacterTombstoneStore: NewMockCharacterTombstoneStore(ctrl),
		CharacterTitleStore:     NewMockCharacterTitleStore(ctrl),
		LevelStore:              NewMockLevelStore(ctrl),
	}
}

// Characters implements the DataStore interface.
func (ds *MockDataStore) Characters() CharacterStore {
	return ds.CharacterStore
}

// CharacterTombstones implements the DataStore interface.
func (ds *MockDataStore) CharacterTombstones() CharacterTombstoneStore {
	return ds.CharacterTombstoneStore
}

// CharacterTitles implements the DataStore interface.
func (ds *MockDataStore) CharacterTitles() CharacterTitleStore {
	return ds.CharacterTitleStore
}

// Levels implements the DataStore interface.
func (ds *MockDataStore) Levels() LevelStore {
	return ds.LevelStore
}
