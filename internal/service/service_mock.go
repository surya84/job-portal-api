// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source=service.go -destination=service_mock.go -package=service
//
// Package service is a generated GoMock package.
package service

import (
	context "context"
	models "job-portal/internal/models"
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockService) Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, email, password)
	ret0, _ := ret[0].(jwt.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockServiceMockRecorder) Authenticate(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockService)(nil).Authenticate), ctx, email, password)
}

// CreateCompany mocks base method.
func (m *MockService) CreateCompany(ctx context.Context, ni models.NewCompany) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, ni)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockServiceMockRecorder) CreateCompany(ctx, ni any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockService)(nil).CreateCompany), ctx, ni)
}

// CreateJob mocks base method.
func (m *MockService) CreateJob(ctx context.Context, nj models.NewJob, cId int) (models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateJob", ctx, nj, cId)
	ret0, _ := ret[0].(models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJob indicates an expected call of CreateJob.
func (mr *MockServiceMockRecorder) CreateJob(ctx, nj, cId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateJob", reflect.TypeOf((*MockService)(nil).CreateJob), ctx, nj, cId)
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, nu)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(ctx, nu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), ctx, nu)
}

// GetCompanyInfoByID mocks base method.
func (m *MockService) GetCompanyInfoByID(ctx context.Context, uid int) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyInfoByID", ctx, uid)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyInfoByID indicates an expected call of GetCompanyInfoByID.
func (mr *MockServiceMockRecorder) GetCompanyInfoByID(ctx, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyInfoByID", reflect.TypeOf((*MockService)(nil).GetCompanyInfoByID), ctx, uid)
}

// GetJobInfoByID mocks base method.
func (m *MockService) GetJobInfoByID(ctx context.Context, jId int) (models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobInfoByID", ctx, jId)
	ret0, _ := ret[0].(models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobInfoByID indicates an expected call of GetJobInfoByID.
func (mr *MockServiceMockRecorder) GetJobInfoByID(ctx, jId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobInfoByID", reflect.TypeOf((*MockService)(nil).GetJobInfoByID), ctx, jId)
}

// ProcessJob mocks base method.
func (m *MockService) ProcessJob(ctx context.Context, id int, nj models.NewJob) (models.NewJob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessJob", ctx, id, nj)
	ret0, _ := ret[0].(models.NewJob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessJob indicates an expected call of ProcessJob.
func (mr *MockServiceMockRecorder) ProcessJob(ctx, id, nj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessJob", reflect.TypeOf((*MockService)(nil).ProcessJob), ctx, id, nj)
}

// ViewCompany mocks base method.
func (m *MockService) ViewCompany(ctx context.Context) ([]models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCompany", ctx)
	ret0, _ := ret[0].([]models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewCompany indicates an expected call of ViewCompany.
func (mr *MockServiceMockRecorder) ViewCompany(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCompany", reflect.TypeOf((*MockService)(nil).ViewCompany), ctx)
}

// ViewJob mocks base method.
func (m *MockService) ViewJob(ctx context.Context) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJob", ctx)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewJob indicates an expected call of ViewJob.
func (mr *MockServiceMockRecorder) ViewJob(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJob", reflect.TypeOf((*MockService)(nil).ViewJob), ctx)
}

// ViewJobByCompanyId mocks base method.
func (m *MockService) ViewJobByCompanyId(ctx context.Context, cId int) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJobByCompanyId", ctx, cId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewJobByCompanyId indicates an expected call of ViewJobByCompanyId.
func (mr *MockServiceMockRecorder) ViewJobByCompanyId(ctx, cId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJobByCompanyId", reflect.TypeOf((*MockService)(nil).ViewJobByCompanyId), ctx, cId)
}
