// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MongoDBNavigator/go-backend/domain/database/repository (interfaces: CollectionWriter)

// Package mock is a generated GoMock package.
package mock

import (
	value "github.com/MongoDBNavigator/go-backend/domain/database/value"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCollectionWriter is a mock of CollectionWriter interface
type MockCollectionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionWriterMockRecorder
}

// MockCollectionWriterMockRecorder is the mock recorder for MockCollectionWriter
type MockCollectionWriterMockRecorder struct {
	mock *MockCollectionWriter
}

// NewMockCollectionWriter creates a new mock instance
func NewMockCollectionWriter(ctrl *gomock.Controller) *MockCollectionWriter {
	mock := &MockCollectionWriter{ctrl: ctrl}
	mock.recorder = &MockCollectionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollectionWriter) EXPECT() *MockCollectionWriterMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCollectionWriter) Create(arg0 value.DBName, arg1 value.CollName) error {
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockCollectionWriterMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCollectionWriter)(nil).Create), arg0, arg1)
}

// Delete mocks base method
func (m *MockCollectionWriter) Delete(arg0 value.DBName, arg1 value.CollName) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCollectionWriterMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCollectionWriter)(nil).Delete), arg0, arg1)
}
