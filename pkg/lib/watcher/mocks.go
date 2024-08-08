// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package watcher is a generated GoMock package.
package watcher

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockWatcher is a mock of Watcher interface.
type MockWatcher struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherMockRecorder
}

// MockWatcherMockRecorder is the mock recorder for MockWatcher.
type MockWatcherMockRecorder struct {
	mock *MockWatcher
}

// NewMockWatcher creates a new mock instance.
func NewMockWatcher(ctrl *gomock.Controller) *MockWatcher {
	mock := &MockWatcher{ctrl: ctrl}
	mock.recorder = &MockWatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcher) EXPECT() *MockWatcherMockRecorder {
	return m.recorder
}

// Checkpoint mocks base method.
func (m *MockWatcher) Checkpoint(ctx context.Context, eventSeqNum uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkpoint", ctx, eventSeqNum)
	ret0, _ := ret[0].(error)
	return ret0
}

// Checkpoint indicates an expected call of Checkpoint.
func (mr *MockWatcherMockRecorder) Checkpoint(ctx, eventSeqNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkpoint", reflect.TypeOf((*MockWatcher)(nil).Checkpoint), ctx, eventSeqNum)
}

// ID mocks base method.
func (m *MockWatcher) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockWatcherMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockWatcher)(nil).ID))
}

// SeekToOffset mocks base method.
func (m *MockWatcher) SeekToOffset(ctx context.Context, eventSeqNum uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeekToOffset", ctx, eventSeqNum)
	ret0, _ := ret[0].(error)
	return ret0
}

// SeekToOffset indicates an expected call of SeekToOffset.
func (mr *MockWatcherMockRecorder) SeekToOffset(ctx, eventSeqNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeekToOffset", reflect.TypeOf((*MockWatcher)(nil).SeekToOffset), ctx, eventSeqNum)
}

// Stats mocks base method.
func (m *MockWatcher) Stats() Stats {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(Stats)
	return ret0
}

// Stats indicates an expected call of Stats.
func (mr *MockWatcherMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockWatcher)(nil).Stats))
}

// Stop mocks base method.
func (m *MockWatcher) Stop(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", ctx)
}

// Stop indicates an expected call of Stop.
func (mr *MockWatcherMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWatcher)(nil).Stop), ctx)
}

// MockEventHandler is a mock of EventHandler interface.
type MockEventHandler struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerMockRecorder
}

// MockEventHandlerMockRecorder is the mock recorder for MockEventHandler.
type MockEventHandlerMockRecorder struct {
	mock *MockEventHandler
}

// NewMockEventHandler creates a new mock instance.
func NewMockEventHandler(ctrl *gomock.Controller) *MockEventHandler {
	mock := &MockEventHandler{ctrl: ctrl}
	mock.recorder = &MockEventHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventHandler) EXPECT() *MockEventHandlerMockRecorder {
	return m.recorder
}

// HandleEvent mocks base method.
func (m *MockEventHandler) HandleEvent(ctx context.Context, event Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleEvent", ctx, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleEvent indicates an expected call of HandleEvent.
func (mr *MockEventHandlerMockRecorder) HandleEvent(ctx, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleEvent", reflect.TypeOf((*MockEventHandler)(nil).HandleEvent), ctx, event)
}

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// GetWatcher mocks base method.
func (m *MockRegistry) GetWatcher(watcherID string) (Watcher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWatcher", watcherID)
	ret0, _ := ret[0].(Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWatcher indicates an expected call of GetWatcher.
func (mr *MockRegistryMockRecorder) GetWatcher(watcherID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWatcher", reflect.TypeOf((*MockRegistry)(nil).GetWatcher), watcherID)
}

// Stop mocks base method.
func (m *MockRegistry) Stop(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockRegistryMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockRegistry)(nil).Stop), ctx)
}

// Watch mocks base method.
func (m *MockRegistry) Watch(ctx context.Context, watcherID string, handler EventHandler, opts ...WatchOption) (Watcher, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, watcherID, handler}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Watch", varargs...)
	ret0, _ := ret[0].(Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockRegistryMockRecorder) Watch(ctx, watcherID, handler interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, watcherID, handler}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockRegistry)(nil).Watch), varargs...)
}

// MockEventStore is a mock of EventStore interface.
type MockEventStore struct {
	ctrl     *gomock.Controller
	recorder *MockEventStoreMockRecorder
}

// MockEventStoreMockRecorder is the mock recorder for MockEventStore.
type MockEventStoreMockRecorder struct {
	mock *MockEventStore
}

// NewMockEventStore creates a new mock instance.
func NewMockEventStore(ctrl *gomock.Controller) *MockEventStore {
	mock := &MockEventStore{ctrl: ctrl}
	mock.recorder = &MockEventStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventStore) EXPECT() *MockEventStoreMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockEventStore) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockEventStoreMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEventStore)(nil).Close), ctx)
}

// GetCheckpoint mocks base method.
func (m *MockEventStore) GetCheckpoint(ctx context.Context, watcherID string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckpoint", ctx, watcherID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckpoint indicates an expected call of GetCheckpoint.
func (mr *MockEventStoreMockRecorder) GetCheckpoint(ctx, watcherID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckpoint", reflect.TypeOf((*MockEventStore)(nil).GetCheckpoint), ctx, watcherID)
}

// GetEvents mocks base method.
func (m *MockEventStore) GetEvents(ctx context.Context, params GetEventsRequest) (*GetEventsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvents", ctx, params)
	ret0, _ := ret[0].(*GetEventsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents.
func (mr *MockEventStoreMockRecorder) GetEvents(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockEventStore)(nil).GetEvents), ctx, params)
}

// GetLatestEventNum mocks base method.
func (m *MockEventStore) GetLatestEventNum(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestEventNum", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestEventNum indicates an expected call of GetLatestEventNum.
func (mr *MockEventStoreMockRecorder) GetLatestEventNum(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestEventNum", reflect.TypeOf((*MockEventStore)(nil).GetLatestEventNum), ctx)
}

// StoreCheckpoint mocks base method.
func (m *MockEventStore) StoreCheckpoint(ctx context.Context, watcherID string, eventSeqNum uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreCheckpoint", ctx, watcherID, eventSeqNum)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreCheckpoint indicates an expected call of StoreCheckpoint.
func (mr *MockEventStoreMockRecorder) StoreCheckpoint(ctx, watcherID, eventSeqNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreCheckpoint", reflect.TypeOf((*MockEventStore)(nil).StoreCheckpoint), ctx, watcherID, eventSeqNum)
}

// StoreEvent mocks base method.
func (m *MockEventStore) StoreEvent(ctx context.Context, operation Operation, objectType string, object interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreEvent", ctx, operation, objectType, object)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreEvent indicates an expected call of StoreEvent.
func (mr *MockEventStoreMockRecorder) StoreEvent(ctx, operation, objectType, object interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreEvent", reflect.TypeOf((*MockEventStore)(nil).StoreEvent), ctx, operation, objectType, object)
}

// MockSerializer is a mock of Serializer interface.
type MockSerializer struct {
	ctrl     *gomock.Controller
	recorder *MockSerializerMockRecorder
}

// MockSerializerMockRecorder is the mock recorder for MockSerializer.
type MockSerializerMockRecorder struct {
	mock *MockSerializer
}

// NewMockSerializer creates a new mock instance.
func NewMockSerializer(ctrl *gomock.Controller) *MockSerializer {
	mock := &MockSerializer{ctrl: ctrl}
	mock.recorder = &MockSerializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSerializer) EXPECT() *MockSerializerMockRecorder {
	return m.recorder
}

// Marshal mocks base method.
func (m *MockSerializer) Marshal(event Event) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal", event)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal.
func (mr *MockSerializerMockRecorder) Marshal(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockSerializer)(nil).Marshal), event)
}

// Unmarshal mocks base method.
func (m *MockSerializer) Unmarshal(data []byte, event *Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", data, event)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unmarshal indicates an expected call of Unmarshal.
func (mr *MockSerializerMockRecorder) Unmarshal(data, event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockSerializer)(nil).Unmarshal), data, event)
}