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

// CreateCompany mocks base method.
func (m *MockService) CreateCompany(ctx context.Context, ni models.NewCompany, userId uint) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, ni, userId)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockServiceMockRecorder) CreateCompany(ctx, ni, userId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockService)(nil).CreateCompany), ctx, ni, userId)
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
func (m *MockService) GetCompanyInfoByID(uid int) (models.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyInfoByID", uid)
	ret0, _ := ret[0].(models.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyInfoByID indicates an expected call of GetCompanyInfoByID.
func (mr *MockServiceMockRecorder) GetCompanyInfoByID(uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyInfoByID", reflect.TypeOf((*MockService)(nil).GetCompanyInfoByID), uid)
}

// GetJobInfoByID mocks base method.
func (m *MockService) GetJobInfoByID(jId int) (models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobInfoByID", jId)
	ret0, _ := ret[0].(models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobInfoByID indicates an expected call of GetJobInfoByID.
func (mr *MockServiceMockRecorder) GetJobInfoByID(jId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobInfoByID", reflect.TypeOf((*MockService)(nil).GetJobInfoByID), jId)
}

// UserSignin mocks base method.
func (m *MockService) UserSignin(ctx context.Context, email, password string) (jwt.RegisteredClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignin", ctx, email, password)
	ret0, _ := ret[0].(jwt.RegisteredClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignin indicates an expected call of UserSignin.
func (mr *MockServiceMockRecorder) UserSignin(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignin", reflect.TypeOf((*MockService)(nil).UserSignin), ctx, email, password)
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
func (m *MockService) ViewJobByCompanyId(cId int) ([]models.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewJobByCompanyId", cId)
	ret0, _ := ret[0].([]models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewJobByCompanyId indicates an expected call of ViewJobByCompanyId.
func (mr *MockServiceMockRecorder) ViewJobByCompanyId(cId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewJobByCompanyId", reflect.TypeOf((*MockService)(nil).ViewJobByCompanyId), cId)
}
