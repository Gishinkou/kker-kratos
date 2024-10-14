// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source=interface.go -destination=interface_mock.go -package=authserviceiface AuthService
//

// Package authserviceiface is a generated GoMock package.
package authserviceiface

import (
	context "context"
	reflect "reflect"

	verificationcode "github.com/Gishinkou/kker-kratos/backend/baseService/internal/domain/entity/verificationcode"
	gomock "go.uber.org/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// CreateVerificationCode mocks base method.
func (m *MockAuthService) CreateVerificationCode(ctx context.Context, bits, expireTime int64) (*verificationcode.VerificationCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVerificationCode", ctx, bits, expireTime)
	ret0, _ := ret[0].(*verificationcode.VerificationCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVerificationCode indicates an expected call of CreateVerificationCode.
func (mr *MockAuthServiceMockRecorder) CreateVerificationCode(ctx, bits, expireTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVerificationCode", reflect.TypeOf((*MockAuthService)(nil).CreateVerificationCode), ctx, bits, expireTime)
}

// ValidateVerificationCode mocks base method.
func (m *MockAuthService) ValidateVerificationCode(ctx context.Context, code *verificationcode.VerificationCode) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateVerificationCode", ctx, code)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateVerificationCode indicates an expected call of ValidateVerificationCode.
func (mr *MockAuthServiceMockRecorder) ValidateVerificationCode(ctx, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateVerificationCode", reflect.TypeOf((*MockAuthService)(nil).ValidateVerificationCode), ctx, code)
}
