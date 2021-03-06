// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MongoDBNavigator/go-backend/domain/database/repository (interfaces: ValidationWriter)

// Package mock is a generated GoMock package.
package mock

import (
	value "github.com/MongoDBNavigator/go-backend/domain/database/value"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockValidationWriter is a mock of ValidationWriter interface
type MockValidationWriter struct {
	ctrl     *gomock.Controller
	recorder *MockValidationWriterMockRecorder
}

// MockValidationWriterMockRecorder is the mock recorder for MockValidationWriter
type MockValidationWriterMockRecorder struct {
	mock *MockValidationWriter
}

// NewMockValidationWriter creates a new mock instance
func NewMockValidationWriter(ctrl *gomock.Controller) *MockValidationWriter {
	mock := &MockValidationWriter{ctrl: ctrl}
	mock.recorder = &MockValidationWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidationWriter) EXPECT() *MockValidationWriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockValidationWriter) Write(arg0 value.DBName, arg1 value.CollName) error {
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockValidationWriterMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockValidationWriter)(nil).Write), arg0, arg1)
}
