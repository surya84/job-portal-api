package service

import (
	"context"
	"job-portal/internal/models"
	"job-portal/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service
type Service interface {
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	UserSignin(ctx context.Context, email, password string) (jwt.RegisteredClaims, error)
	CreateJob(ctx context.Context, nj models.NewJob, cId int) (models.Job, error)
	ViewJob() ([]models.Job, error)
	GetJobInfoByID(jId int) (models.Job, error)
	ViewJobByCompanyId(cId int) ([]models.Job, error)
	CreateCompany(ctx context.Context, ni models.NewCompany, userId uint) (models.Company, error)
	ViewCompany() ([]models.Company, error)
	GetCompanyInfoByID(uid int) (models.Company, error)
}

type NewService struct {
	Service
	rs repository.Repository
}

func NewServiceStore(s repository.Repository) NewService {
	return NewService{rs: s}
}
