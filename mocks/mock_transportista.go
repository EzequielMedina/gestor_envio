// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/ports/transportista.go
//
// Generated by this command:
//
//	mockgen -source=internal/core/ports/transportista.go -destination=mocks/mock_transportista.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	domain "main.go/internal/core/domain"
)

// MockTransportistaService is a mock of TransportistaService interface.
type MockTransportistaService struct {
	ctrl     *gomock.Controller
	recorder *MockTransportistaServiceMockRecorder
	isgomock struct{}
}

// MockTransportistaServiceMockRecorder is the mock recorder for MockTransportistaService.
type MockTransportistaServiceMockRecorder struct {
	mock *MockTransportistaService
}

// NewMockTransportistaService creates a new mock instance.
func NewMockTransportistaService(ctrl *gomock.Controller) *MockTransportistaService {
	mock := &MockTransportistaService{ctrl: ctrl}
	mock.recorder = &MockTransportistaServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransportistaService) EXPECT() *MockTransportistaServiceMockRecorder {
	return m.recorder
}

// ObtenerTransportistaByEmail mocks base method.
func (m *MockTransportistaService) ObtenerTransportistaByEmail(email string) (domain.Transportista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObtenerTransportistaByEmail", email)
	ret0, _ := ret[0].(domain.Transportista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObtenerTransportistaByEmail indicates an expected call of ObtenerTransportistaByEmail.
func (mr *MockTransportistaServiceMockRecorder) ObtenerTransportistaByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObtenerTransportistaByEmail", reflect.TypeOf((*MockTransportistaService)(nil).ObtenerTransportistaByEmail), email)
}

// MockTransportistaRepository is a mock of TransportistaRepository interface.
type MockTransportistaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransportistaRepositoryMockRecorder
	isgomock struct{}
}

// MockTransportistaRepositoryMockRecorder is the mock recorder for MockTransportistaRepository.
type MockTransportistaRepositoryMockRecorder struct {
	mock *MockTransportistaRepository
}

// NewMockTransportistaRepository creates a new mock instance.
func NewMockTransportistaRepository(ctrl *gomock.Controller) *MockTransportistaRepository {
	mock := &MockTransportistaRepository{ctrl: ctrl}
	mock.recorder = &MockTransportistaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransportistaRepository) EXPECT() *MockTransportistaRepositoryMockRecorder {
	return m.recorder
}

// ObtenerTransportistaByEmail mocks base method.
func (m *MockTransportistaRepository) ObtenerTransportistaByEmail(email string) (domain.Transportista, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObtenerTransportistaByEmail", email)
	ret0, _ := ret[0].(domain.Transportista)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObtenerTransportistaByEmail indicates an expected call of ObtenerTransportistaByEmail.
func (mr *MockTransportistaRepositoryMockRecorder) ObtenerTransportistaByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObtenerTransportistaByEmail", reflect.TypeOf((*MockTransportistaRepository)(nil).ObtenerTransportistaByEmail), email)
}
