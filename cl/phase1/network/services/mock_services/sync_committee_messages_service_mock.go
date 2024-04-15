// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/cl/phase1/network/services (interfaces: SyncCommitteeMessagesService)
//
// Generated by this command:
//
//	mockgen -destination=./mock_services/sync_committee_messages_service_mock.go -package=mock_services . SyncCommitteeMessagesService
//

// Package mock_services is a generated GoMock package.
package mock_services

import (
	context "context"
	reflect "reflect"

	cltypes "github.com/ledgerwatch/erigon/cl/cltypes"
	gomock "go.uber.org/mock/gomock"
)

// MockSyncCommitteeMessagesService is a mock of SyncCommitteeMessagesService interface.
type MockSyncCommitteeMessagesService struct {
	ctrl     *gomock.Controller
	recorder *MockSyncCommitteeMessagesServiceMockRecorder
}

// MockSyncCommitteeMessagesServiceMockRecorder is the mock recorder for MockSyncCommitteeMessagesService.
type MockSyncCommitteeMessagesServiceMockRecorder struct {
	mock *MockSyncCommitteeMessagesService
}

// NewMockSyncCommitteeMessagesService creates a new mock instance.
func NewMockSyncCommitteeMessagesService(ctrl *gomock.Controller) *MockSyncCommitteeMessagesService {
	mock := &MockSyncCommitteeMessagesService{ctrl: ctrl}
	mock.recorder = &MockSyncCommitteeMessagesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncCommitteeMessagesService) EXPECT() *MockSyncCommitteeMessagesServiceMockRecorder {
	return m.recorder
}

// ProcessMessage mocks base method.
func (m *MockSyncCommitteeMessagesService) ProcessMessage(arg0 context.Context, arg1 *uint64, arg2 *cltypes.SyncCommitteeMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMessage", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMessage indicates an expected call of ProcessMessage.
func (mr *MockSyncCommitteeMessagesServiceMockRecorder) ProcessMessage(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMessage", reflect.TypeOf((*MockSyncCommitteeMessagesService)(nil).ProcessMessage), arg0, arg1, arg2)
}