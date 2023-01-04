// Code generated by MockGen. DO NOT EDIT.
// Source: sink.go

// Package hook is a generated GoMock package.
package hook

import (
	url "net/url"
	reflect "reflect"

	event "github.com/authgear/authgear-server/pkg/api/event"
	accesscontrol "github.com/authgear/authgear-server/pkg/util/accesscontrol"
	gomock "github.com/golang/mock/gomock"
)

// MockStandardAttributesServiceNoEvent is a mock of StandardAttributesServiceNoEvent interface.
type MockStandardAttributesServiceNoEvent struct {
	ctrl     *gomock.Controller
	recorder *MockStandardAttributesServiceNoEventMockRecorder
}

// MockStandardAttributesServiceNoEventMockRecorder is the mock recorder for MockStandardAttributesServiceNoEvent.
type MockStandardAttributesServiceNoEventMockRecorder struct {
	mock *MockStandardAttributesServiceNoEvent
}

// NewMockStandardAttributesServiceNoEvent creates a new mock instance.
func NewMockStandardAttributesServiceNoEvent(ctrl *gomock.Controller) *MockStandardAttributesServiceNoEvent {
	mock := &MockStandardAttributesServiceNoEvent{ctrl: ctrl}
	mock.recorder = &MockStandardAttributesServiceNoEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStandardAttributesServiceNoEvent) EXPECT() *MockStandardAttributesServiceNoEventMockRecorder {
	return m.recorder
}

// UpdateStandardAttributes mocks base method.
func (m *MockStandardAttributesServiceNoEvent) UpdateStandardAttributes(role accesscontrol.Role, userID string, stdAttrs map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStandardAttributes", role, userID, stdAttrs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStandardAttributes indicates an expected call of UpdateStandardAttributes.
func (mr *MockStandardAttributesServiceNoEventMockRecorder) UpdateStandardAttributes(role, userID, stdAttrs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStandardAttributes", reflect.TypeOf((*MockStandardAttributesServiceNoEvent)(nil).UpdateStandardAttributes), role, userID, stdAttrs)
}

// MockCustomAttributesServiceNoEvent is a mock of CustomAttributesServiceNoEvent interface.
type MockCustomAttributesServiceNoEvent struct {
	ctrl     *gomock.Controller
	recorder *MockCustomAttributesServiceNoEventMockRecorder
}

// MockCustomAttributesServiceNoEventMockRecorder is the mock recorder for MockCustomAttributesServiceNoEvent.
type MockCustomAttributesServiceNoEventMockRecorder struct {
	mock *MockCustomAttributesServiceNoEvent
}

// NewMockCustomAttributesServiceNoEvent creates a new mock instance.
func NewMockCustomAttributesServiceNoEvent(ctrl *gomock.Controller) *MockCustomAttributesServiceNoEvent {
	mock := &MockCustomAttributesServiceNoEvent{ctrl: ctrl}
	mock.recorder = &MockCustomAttributesServiceNoEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomAttributesServiceNoEvent) EXPECT() *MockCustomAttributesServiceNoEventMockRecorder {
	return m.recorder
}

// UpdateAllCustomAttributes mocks base method.
func (m *MockCustomAttributesServiceNoEvent) UpdateAllCustomAttributes(role accesscontrol.Role, userID string, reprForm map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAllCustomAttributes", role, userID, reprForm)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAllCustomAttributes indicates an expected call of UpdateAllCustomAttributes.
func (mr *MockCustomAttributesServiceNoEventMockRecorder) UpdateAllCustomAttributes(role, userID, reprForm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAllCustomAttributes", reflect.TypeOf((*MockCustomAttributesServiceNoEvent)(nil).UpdateAllCustomAttributes), role, userID, reprForm)
}

// MockWebHook is a mock of WebHook interface.
type MockWebHook struct {
	ctrl     *gomock.Controller
	recorder *MockWebHookMockRecorder
}

// MockWebHookMockRecorder is the mock recorder for MockWebHook.
type MockWebHookMockRecorder struct {
	mock *MockWebHook
}

// NewMockWebHook creates a new mock instance.
func NewMockWebHook(ctrl *gomock.Controller) *MockWebHook {
	mock := &MockWebHook{ctrl: ctrl}
	mock.recorder = &MockWebHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebHook) EXPECT() *MockWebHookMockRecorder {
	return m.recorder
}

// DeliverBlockingEvent mocks base method.
func (m *MockWebHook) DeliverBlockingEvent(u *url.URL, e *event.Event) (*event.HookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverBlockingEvent", u, e)
	ret0, _ := ret[0].(*event.HookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeliverBlockingEvent indicates an expected call of DeliverBlockingEvent.
func (mr *MockWebHookMockRecorder) DeliverBlockingEvent(u, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverBlockingEvent", reflect.TypeOf((*MockWebHook)(nil).DeliverBlockingEvent), u, e)
}

// DeliverNonBlockingEvent mocks base method.
func (m *MockWebHook) DeliverNonBlockingEvent(u *url.URL, e *event.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverNonBlockingEvent", u, e)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverNonBlockingEvent indicates an expected call of DeliverNonBlockingEvent.
func (mr *MockWebHookMockRecorder) DeliverNonBlockingEvent(u, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverNonBlockingEvent", reflect.TypeOf((*MockWebHook)(nil).DeliverNonBlockingEvent), u, e)
}

// SupportURL mocks base method.
func (m *MockWebHook) SupportURL(u *url.URL) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportURL", u)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SupportURL indicates an expected call of SupportURL.
func (mr *MockWebHookMockRecorder) SupportURL(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportURL", reflect.TypeOf((*MockWebHook)(nil).SupportURL), u)
}

// MockDenoHook is a mock of DenoHook interface.
type MockDenoHook struct {
	ctrl     *gomock.Controller
	recorder *MockDenoHookMockRecorder
}

// MockDenoHookMockRecorder is the mock recorder for MockDenoHook.
type MockDenoHookMockRecorder struct {
	mock *MockDenoHook
}

// NewMockDenoHook creates a new mock instance.
func NewMockDenoHook(ctrl *gomock.Controller) *MockDenoHook {
	mock := &MockDenoHook{ctrl: ctrl}
	mock.recorder = &MockDenoHookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDenoHook) EXPECT() *MockDenoHookMockRecorder {
	return m.recorder
}

// DeliverBlockingEvent mocks base method.
func (m *MockDenoHook) DeliverBlockingEvent(u *url.URL, e *event.Event) (*event.HookResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverBlockingEvent", u, e)
	ret0, _ := ret[0].(*event.HookResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeliverBlockingEvent indicates an expected call of DeliverBlockingEvent.
func (mr *MockDenoHookMockRecorder) DeliverBlockingEvent(u, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverBlockingEvent", reflect.TypeOf((*MockDenoHook)(nil).DeliverBlockingEvent), u, e)
}

// DeliverNonBlockingEvent mocks base method.
func (m *MockDenoHook) DeliverNonBlockingEvent(u *url.URL, e *event.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverNonBlockingEvent", u, e)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeliverNonBlockingEvent indicates an expected call of DeliverNonBlockingEvent.
func (mr *MockDenoHookMockRecorder) DeliverNonBlockingEvent(u, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverNonBlockingEvent", reflect.TypeOf((*MockDenoHook)(nil).DeliverNonBlockingEvent), u, e)
}

// SupportURL mocks base method.
func (m *MockDenoHook) SupportURL(u *url.URL) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportURL", u)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SupportURL indicates an expected call of SupportURL.
func (mr *MockDenoHookMockRecorder) SupportURL(u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportURL", reflect.TypeOf((*MockDenoHook)(nil).SupportURL), u)
}