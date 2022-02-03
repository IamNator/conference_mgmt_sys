// Code generated by MockGen. DO NOT EDIT.
// Source: conference.go

// Package mock is a generated GoMock package.
package mock

import (
	model "conference/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIConferenceService is a mock of IConferenceService interface.
type MockIConferenceService struct {
	ctrl     *gomock.Controller
	recorder *MockIConferenceServiceMockRecorder
}

// MockIConferenceServiceMockRecorder is the mock recorder for MockIConferenceService.
type MockIConferenceServiceMockRecorder struct {
	mock *MockIConferenceService
}

// NewMockIConferenceService creates a new mock instance.
func NewMockIConferenceService(ctrl *gomock.Controller) *MockIConferenceService {
	mock := &MockIConferenceService{ctrl: ctrl}
	mock.recorder = &MockIConferenceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIConferenceService) EXPECT() *MockIConferenceServiceMockRecorder {
	return m.recorder
}

// CreateConference mocks base method.
func (m *MockIConferenceService) CreateConference(conference model.CreateConferenceReq) (*model.Conference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConference", conference)
	ret0, _ := ret[0].(*model.Conference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConference indicates an expected call of CreateConference.
func (mr *MockIConferenceServiceMockRecorder) CreateConference(conference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConference", reflect.TypeOf((*MockIConferenceService)(nil).CreateConference), conference)
}
