// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/repositories.go

// Package repositories_mocks is a generated GoMock package.
package repositories_mocks

import (
	context "context"
	repositories "go-scheduler-api/internal/repositories"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockJobRepository is a mock of JobRepository interface.
type MockJobRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJobRepositoryMockRecorder
}

// MockJobRepositoryMockRecorder is the mock recorder for MockJobRepository.
type MockJobRepositoryMockRecorder struct {
	mock *MockJobRepository
}

// NewMockJobRepository creates a new mock instance.
func NewMockJobRepository(ctrl *gomock.Controller) *MockJobRepository {
	mock := &MockJobRepository{ctrl: ctrl}
	mock.recorder = &MockJobRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobRepository) EXPECT() *MockJobRepositoryMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockJobRepository) Count(ctx context.Context, query repositories.JobQueryParams) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, query)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockJobRepositoryMockRecorder) Count(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockJobRepository)(nil).Count), ctx, query)
}

// Create mocks base method.
func (m *MockJobRepository) Create(ctx context.Context, job repositories.Job) (*repositories.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, job)
	ret0, _ := ret[0].(*repositories.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockJobRepositoryMockRecorder) Create(ctx, job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockJobRepository)(nil).Create), ctx, job)
}

// FindAll mocks base method.
func (m *MockJobRepository) FindAll(ctx context.Context, query repositories.JobQueryParams) ([]repositories.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, query)
	ret0, _ := ret[0].([]repositories.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockJobRepositoryMockRecorder) FindAll(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockJobRepository)(nil).FindAll), ctx, query)
}

// FindById mocks base method.
func (m *MockJobRepository) FindById(ctx context.Context, id uuid.UUID) (*repositories.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", ctx, id)
	ret0, _ := ret[0].(*repositories.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockJobRepositoryMockRecorder) FindById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockJobRepository)(nil).FindById), ctx, id)
}

// MockJobScheduleRepository is a mock of JobScheduleRepository interface.
type MockJobScheduleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockJobScheduleRepositoryMockRecorder
}

// MockJobScheduleRepositoryMockRecorder is the mock recorder for MockJobScheduleRepository.
type MockJobScheduleRepositoryMockRecorder struct {
	mock *MockJobScheduleRepository
}

// NewMockJobScheduleRepository creates a new mock instance.
func NewMockJobScheduleRepository(ctrl *gomock.Controller) *MockJobScheduleRepository {
	mock := &MockJobScheduleRepository{ctrl: ctrl}
	mock.recorder = &MockJobScheduleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobScheduleRepository) EXPECT() *MockJobScheduleRepositoryMockRecorder {
	return m.recorder
}

// CreateMany mocks base method.
func (m *MockJobScheduleRepository) CreateMany(ctx context.Context, schedules ...repositories.JobSchedule) ([]repositories.JobSchedule, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range schedules {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMany", varargs...)
	ret0, _ := ret[0].([]repositories.JobSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMany indicates an expected call of CreateMany.
func (mr *MockJobScheduleRepositoryMockRecorder) CreateMany(ctx interface{}, schedules ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, schedules...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMany", reflect.TypeOf((*MockJobScheduleRepository)(nil).CreateMany), varargs...)
}

// FindByIds mocks base method.
func (m *MockJobScheduleRepository) FindByIds(ctx context.Context, id ...uuid.UUID) ([]repositories.JobSchedule, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range id {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByIds", varargs...)
	ret0, _ := ret[0].([]repositories.JobSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIds indicates an expected call of FindByIds.
func (mr *MockJobScheduleRepositoryMockRecorder) FindByIds(ctx interface{}, id ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, id...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIds", reflect.TypeOf((*MockJobScheduleRepository)(nil).FindByIds), varargs...)
}

// FindByJobId mocks base method.
func (m *MockJobScheduleRepository) FindByJobId(ctx context.Context, jobId uuid.UUID) ([]repositories.JobSchedule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByJobId", ctx, jobId)
	ret0, _ := ret[0].([]repositories.JobSchedule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByJobId indicates an expected call of FindByJobId.
func (mr *MockJobScheduleRepositoryMockRecorder) FindByJobId(ctx, jobId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByJobId", reflect.TypeOf((*MockJobScheduleRepository)(nil).FindByJobId), ctx, jobId)
}
