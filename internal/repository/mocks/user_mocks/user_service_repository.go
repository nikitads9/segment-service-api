// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nikitads9/segment-service-api/internal/repository/user (interfaces: Repository)

// Package user_mocks is a generated GoMock package.
package user_mocks

import (
	bytes "bytes"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/nikitads9/segment-service-api/internal/model"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddToSegment mocks base method.
func (m *MockRepository) AddToSegment(arg0 context.Context, arg1 string, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToSegment", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToSegment indicates an expected call of AddToSegment.
func (mr *MockRepositoryMockRecorder) AddToSegment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToSegment", reflect.TypeOf((*MockRepository)(nil).AddToSegment), arg0, arg1, arg2)
}

// AddUser mocks base method.
func (m *MockRepository) AddUser(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockRepositoryMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockRepository)(nil).AddUser), arg0, arg1)
}

// GetSegmentId mocks base method.
func (m *MockRepository) GetSegmentId(arg0 context.Context, arg1 string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegmentId", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegmentId indicates an expected call of GetSegmentId.
func (mr *MockRepositoryMockRecorder) GetSegmentId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegmentId", reflect.TypeOf((*MockRepository)(nil).GetSegmentId), arg0, arg1)
}

// GetSegments mocks base method.
func (m *MockRepository) GetSegments(arg0 context.Context, arg1 int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegments", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegments indicates an expected call of GetSegments.
func (mr *MockRepositoryMockRecorder) GetSegments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegments", reflect.TypeOf((*MockRepository)(nil).GetSegments), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockRepository) GetUser(arg0 context.Context, arg1 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockRepositoryMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockRepository)(nil).GetUser), arg0, arg1)
}

// GetUserHistoryCsv mocks base method.
func (m *MockRepository) GetUserHistoryCsv(arg0 context.Context, arg1 int64) (bytes.Buffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserHistoryCsv", arg0, arg1)
	ret0, _ := ret[0].(bytes.Buffer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserHistoryCsv indicates an expected call of GetUserHistoryCsv.
func (mr *MockRepositoryMockRecorder) GetUserHistoryCsv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserHistoryCsv", reflect.TypeOf((*MockRepository)(nil).GetUserHistoryCsv), arg0, arg1)
}

// RemoveFromSegment mocks base method.
func (m *MockRepository) RemoveFromSegment(arg0 context.Context, arg1 string, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromSegment", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromSegment indicates an expected call of RemoveFromSegment.
func (mr *MockRepositoryMockRecorder) RemoveFromSegment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromSegment", reflect.TypeOf((*MockRepository)(nil).RemoveFromSegment), arg0, arg1, arg2)
}

// RemoveUser mocks base method.
func (m *MockRepository) RemoveUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockRepositoryMockRecorder) RemoveUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockRepository)(nil).RemoveUser), arg0, arg1)
}

// SetExpireTime mocks base method.
func (m *MockRepository) SetExpireTime(arg0 context.Context, arg1 *model.SetExpireTimeInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetExpireTime", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetExpireTime indicates an expected call of SetExpireTime.
func (mr *MockRepositoryMockRecorder) SetExpireTime(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetExpireTime", reflect.TypeOf((*MockRepository)(nil).SetExpireTime), arg0, arg1)
}
