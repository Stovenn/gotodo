// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stovenn/gotodo/internal/core/ports (interfaces: TodoService)

// Package mockservice is a generated GoMock package.
package mockservice

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/stovenn/gotodo/internal/core/domain"
)

// MockTodoService is a mock of TodoService interface.
type MockTodoService struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceMockRecorder
}

// MockTodoServiceMockRecorder is the mock recorder for MockTodoService.
type MockTodoServiceMockRecorder struct {
	mock *MockTodoService
}

// NewMockTodoService creates a new mock instance.
func NewMockTodoService(ctrl *gomock.Controller) *MockTodoService {
	mock := &MockTodoService{ctrl: ctrl}
	mock.recorder = &MockTodoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoService) EXPECT() *MockTodoServiceMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockTodoService) CreateTodo(arg0 domain.TodoCreationRequest) (*domain.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", arg0)
	ret0, _ := ret[0].(*domain.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockTodoServiceMockRecorder) CreateTodo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockTodoService)(nil).CreateTodo), arg0)
}

// DeleteAllTodos mocks base method.
func (m *MockTodoService) DeleteAllTodos() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllTodos")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllTodos indicates an expected call of DeleteAllTodos.
func (mr *MockTodoServiceMockRecorder) DeleteAllTodos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllTodos", reflect.TypeOf((*MockTodoService)(nil).DeleteAllTodos))
}

// DeleteTodo mocks base method.
func (m *MockTodoService) DeleteTodo(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockTodoServiceMockRecorder) DeleteTodo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockTodoService)(nil).DeleteTodo), arg0)
}

// DisplayAllTodos mocks base method.
func (m *MockTodoService) DisplayAllTodos() ([]*domain.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisplayAllTodos")
	ret0, _ := ret[0].([]*domain.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisplayAllTodos indicates an expected call of DisplayAllTodos.
func (mr *MockTodoServiceMockRecorder) DisplayAllTodos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisplayAllTodos", reflect.TypeOf((*MockTodoService)(nil).DisplayAllTodos))
}

// DisplayTodo mocks base method.
func (m *MockTodoService) DisplayTodo(arg0 string) (*domain.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisplayTodo", arg0)
	ret0, _ := ret[0].(*domain.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisplayTodo indicates an expected call of DisplayTodo.
func (mr *MockTodoServiceMockRecorder) DisplayTodo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisplayTodo", reflect.TypeOf((*MockTodoService)(nil).DisplayTodo), arg0)
}

// PartiallyUpdateTodo mocks base method.
func (m *MockTodoService) PartiallyUpdateTodo(arg0 string, arg1 domain.TodoPartialUpdateRequest) (*domain.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PartiallyUpdateTodo", arg0, arg1)
	ret0, _ := ret[0].(*domain.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PartiallyUpdateTodo indicates an expected call of PartiallyUpdateTodo.
func (mr *MockTodoServiceMockRecorder) PartiallyUpdateTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PartiallyUpdateTodo", reflect.TypeOf((*MockTodoService)(nil).PartiallyUpdateTodo), arg0, arg1)
}

// UpdateTodo mocks base method.
func (m *MockTodoService) UpdateTodo(arg0 string, arg1 domain.TodoUpdateRequest) (*domain.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", arg0, arg1)
	ret0, _ := ret[0].(*domain.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockTodoServiceMockRecorder) UpdateTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockTodoService)(nil).UpdateTodo), arg0, arg1)
}