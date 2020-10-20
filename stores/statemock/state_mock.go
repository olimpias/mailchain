// Code generated by MockGen. DO NOT EDIT.
// Source: state.go

// Package statemock is a generated GoMock package.
package statemock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mail "github.com/mailchain/mailchain/internal/mail"
	stores "github.com/mailchain/mailchain/stores"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// DeleteMessageRead mocks base method.
func (m *MockState) DeleteMessageRead(messageID mail.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessageRead", messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessageRead indicates an expected call of DeleteMessageRead.
func (mr *MockStateMockRecorder) DeleteMessageRead(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessageRead", reflect.TypeOf((*MockState)(nil).DeleteMessageRead), messageID)
}

// PutMessageRead mocks base method.
func (m *MockState) PutMessageRead(messageID mail.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMessageRead", messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutMessageRead indicates an expected call of PutMessageRead.
func (mr *MockStateMockRecorder) PutMessageRead(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMessageRead", reflect.TypeOf((*MockState)(nil).PutMessageRead), messageID)
}

// GetReadStatus mocks base method.
func (m *MockState) GetReadStatus(messageID mail.ID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadStatus", messageID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReadStatus indicates an expected call of GetReadStatus.
func (mr *MockStateMockRecorder) GetReadStatus(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadStatus", reflect.TypeOf((*MockState)(nil).GetReadStatus), messageID)
}

// PutMessage mocks base method.
func (m *MockState) PutMessage(protocol, network, address string, message stores.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMessage", protocol, network, address, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutMessage indicates an expected call of PutMessage.
func (mr *MockStateMockRecorder) PutMessage(protocol, network, address, message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMessage", reflect.TypeOf((*MockState)(nil).PutMessage), protocol, network, address, message)
}

// GetMessages mocks base method.
func (m *MockState) GetMessages(protocol, network, address string) ([]stores.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", protocol, network, address)
	ret0, _ := ret[0].([]stores.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessages indicates an expected call of GetMessages.
func (mr *MockStateMockRecorder) GetMessages(protocol, network, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockState)(nil).GetMessages), protocol, network, address)
}