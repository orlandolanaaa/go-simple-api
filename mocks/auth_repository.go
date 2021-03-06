// Code generated by MockGen. DO NOT EDIT.
// Source: internal/modules/auth/auth_repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	auth "be_entry_task/internal/modules/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthRepository is a mock of AuthRepository interface.
type MockAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepositoryMockRecorder
}

// MockAuthRepositoryMockRecorder is the mock recorder for MockAuthRepository.
type MockAuthRepositoryMockRecorder struct {
	mock *MockAuthRepository
}

// NewMockAuthRepository creates a new mock instance.
func NewMockAuthRepository(ctrl *gomock.Controller) *MockAuthRepository {
	mock := &MockAuthRepository{ctrl: ctrl}
	mock.recorder = &MockAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepository) EXPECT() *MockAuthRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthRepository) Create(arg0 auth.UserToken) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthRepository)(nil).Create), arg0)
}

// SearchWithToken mocks base method.
func (m *MockAuthRepository) SearchWithToken(token string) (auth.UserToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithToken", token)
	ret0, _ := ret[0].(auth.UserToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithToken indicates an expected call of SearchWithToken.
func (mr *MockAuthRepositoryMockRecorder) SearchWithToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithToken", reflect.TypeOf((*MockAuthRepository)(nil).SearchWithToken), token)
}
