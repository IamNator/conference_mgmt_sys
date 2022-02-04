// Code generated by MockGen. DO NOT EDIT.
// Source: conference.go

// Package mock is a generated GoMock package.
package mock

import (
	model "conference/model"
	storage "conference/storage"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIConferenceRepository is a mock of IConferenceRepository interface.
type MockIConferenceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIConferenceRepositoryMockRecorder
}

// MockIConferenceRepositoryMockRecorder is the mock recorder for MockIConferenceRepository.
type MockIConferenceRepositoryMockRecorder struct {
	mock *MockIConferenceRepository
}

// NewMockIConferenceRepository creates a new mock instance.
func NewMockIConferenceRepository(ctrl *gomock.Controller) *MockIConferenceRepository {
	mock := &MockIConferenceRepository{ctrl: ctrl}
	mock.recorder = &MockIConferenceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIConferenceRepository) EXPECT() *MockIConferenceRepositoryMockRecorder {
	return m.recorder
}

// CreateConference mocks base method.
func (m *MockIConferenceRepository) CreateConference(conference model.Conference) (*model.Conference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateConference", conference)
	ret0, _ := ret[0].(*model.Conference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConference indicates an expected call of CreateConference.
func (mr *MockIConferenceRepositoryMockRecorder) CreateConference(conference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConference", reflect.TypeOf((*MockIConferenceRepository)(nil).CreateConference), conference)
}

// CreateParticipant mocks base method.
func (m *MockIConferenceRepository) CreateParticipant(participant model.Participant) (*model.Participant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParticipant", participant)
	ret0, _ := ret[0].(*model.Participant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateParticipant indicates an expected call of CreateParticipant.
func (mr *MockIConferenceRepositoryMockRecorder) CreateParticipant(participant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParticipant", reflect.TypeOf((*MockIConferenceRepository)(nil).CreateParticipant), participant)
}

// CreateSpeaker mocks base method.
func (m *MockIConferenceRepository) CreateSpeaker(speaker model.Speaker) (*model.Speaker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSpeaker", speaker)
	ret0, _ := ret[0].(*model.Speaker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSpeaker indicates an expected call of CreateSpeaker.
func (mr *MockIConferenceRepositoryMockRecorder) CreateSpeaker(speaker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSpeaker", reflect.TypeOf((*MockIConferenceRepository)(nil).CreateSpeaker), speaker)
}

// CreateTalk mocks base method.
func (m *MockIConferenceRepository) CreateTalk(talk model.Talk) (*model.Talk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTalk", talk)
	ret0, _ := ret[0].(*model.Talk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTalk indicates an expected call of CreateTalk.
func (mr *MockIConferenceRepositoryMockRecorder) CreateTalk(talk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTalk", reflect.TypeOf((*MockIConferenceRepository)(nil).CreateTalk), talk)
}

// DeleteParticipant mocks base method.
func (m *MockIConferenceRepository) DeleteParticipant(participantId, talkId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteParticipant", participantId, talkId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteParticipant indicates an expected call of DeleteParticipant.
func (mr *MockIConferenceRepositoryMockRecorder) DeleteParticipant(participantId, talkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteParticipant", reflect.TypeOf((*MockIConferenceRepository)(nil).DeleteParticipant), participantId, talkId)
}

// DeleteSpeaker mocks base method.
func (m *MockIConferenceRepository) DeleteSpeaker(speakerId, talkId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSpeaker", speakerId, talkId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSpeaker indicates an expected call of DeleteSpeaker.
func (mr *MockIConferenceRepositoryMockRecorder) DeleteSpeaker(speakerId, talkId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSpeaker", reflect.TypeOf((*MockIConferenceRepository)(nil).DeleteSpeaker), speakerId, talkId)
}

// GetAllConference mocks base method.
func (m *MockIConferenceRepository) GetAllConference(page, pageSize int) ([]model.Conference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllConference", page, pageSize)
	ret0, _ := ret[0].([]model.Conference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllConference indicates an expected call of GetAllConference.
func (mr *MockIConferenceRepositoryMockRecorder) GetAllConference(page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllConference", reflect.TypeOf((*MockIConferenceRepository)(nil).GetAllConference), page, pageSize)
}

// GetConference mocks base method.
func (m *MockIConferenceRepository) GetConference(conferenceId uint, page, pageSize int) ([]model.Conference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConference", conferenceId, page, pageSize)
	ret0, _ := ret[0].([]model.Conference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConference indicates an expected call of GetConference.
func (mr *MockIConferenceRepositoryMockRecorder) GetConference(conferenceId, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConference", reflect.TypeOf((*MockIConferenceRepository)(nil).GetConference), conferenceId, page, pageSize)
}

// GetEditHistory mocks base method.
func (m *MockIConferenceRepository) GetEditHistory(conferenceID uint, page, pageSize int) ([]model.EditHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEditHistory", conferenceID, page, pageSize)
	ret0, _ := ret[0].([]model.EditHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEditHistory indicates an expected call of GetEditHistory.
func (mr *MockIConferenceRepositoryMockRecorder) GetEditHistory(conferenceID, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEditHistory", reflect.TypeOf((*MockIConferenceRepository)(nil).GetEditHistory), conferenceID, page, pageSize)
}

// GetParticipants mocks base method.
func (m *MockIConferenceRepository) GetParticipants(TalkId uint, page, pageSize int) ([]model.Talk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParticipants", TalkId, page, pageSize)
	ret0, _ := ret[0].([]model.Talk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParticipants indicates an expected call of GetParticipants.
func (mr *MockIConferenceRepositoryMockRecorder) GetParticipants(TalkId, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParticipants", reflect.TypeOf((*MockIConferenceRepository)(nil).GetParticipants), TalkId, page, pageSize)
}

// GetSpeakers mocks base method.
func (m *MockIConferenceRepository) GetSpeakers(TalkId uint, page, pageSize int) ([]model.Talk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSpeakers", TalkId, page, pageSize)
	ret0, _ := ret[0].([]model.Talk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSpeakers indicates an expected call of GetSpeakers.
func (mr *MockIConferenceRepositoryMockRecorder) GetSpeakers(TalkId, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSpeakers", reflect.TypeOf((*MockIConferenceRepository)(nil).GetSpeakers), TalkId, page, pageSize)
}

// GetTalks mocks base method.
func (m *MockIConferenceRepository) GetTalks(conferenceId uint, page, pageSize int) ([]model.Talk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTalks", conferenceId, page, pageSize)
	ret0, _ := ret[0].([]model.Talk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTalks indicates an expected call of GetTalks.
func (mr *MockIConferenceRepositoryMockRecorder) GetTalks(conferenceId, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTalks", reflect.TypeOf((*MockIConferenceRepository)(nil).GetTalks), conferenceId, page, pageSize)
}

// SaveEditHistory mocks base method.
func (m *MockIConferenceRepository) SaveEditHistory(history model.EditHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveEditHistory", history)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveEditHistory indicates an expected call of SaveEditHistory.
func (mr *MockIConferenceRepositoryMockRecorder) SaveEditHistory(history interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveEditHistory", reflect.TypeOf((*MockIConferenceRepository)(nil).SaveEditHistory), history)
}

// UpdateConference mocks base method.
func (m *MockIConferenceRepository) UpdateConference(conference model.Conference) (*model.Conference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConference", conference)
	ret0, _ := ret[0].(*model.Conference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateConference indicates an expected call of UpdateConference.
func (mr *MockIConferenceRepositoryMockRecorder) UpdateConference(conference interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConference", reflect.TypeOf((*MockIConferenceRepository)(nil).UpdateConference), conference)
}

// UpdateTalk mocks base method.
func (m *MockIConferenceRepository) UpdateTalk(talk model.Talk) (*model.Talk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTalk", talk)
	ret0, _ := ret[0].(*model.Talk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTalk indicates an expected call of UpdateTalk.
func (mr *MockIConferenceRepositoryMockRecorder) UpdateTalk(talk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTalk", reflect.TypeOf((*MockIConferenceRepository)(nil).UpdateTalk), talk)
}

// WithTx mocks base method.
func (m *MockIConferenceRepository) WithTx(db *gorm.DB) storage.IConferenceRepository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithTx", db)
	ret0, _ := ret[0].(storage.IConferenceRepository)
	return ret0
}

// WithTx indicates an expected call of WithTx.
func (mr *MockIConferenceRepositoryMockRecorder) WithTx(db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithTx", reflect.TypeOf((*MockIConferenceRepository)(nil).WithTx), db)
}
