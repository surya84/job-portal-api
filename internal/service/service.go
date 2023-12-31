package service

import (
	"context"
	"job-portal/internal/models"
	rediscache "job-portal/internal/redisCache"
	"job-portal/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

type NewService struct {
	rp  repository.Repository
	rdb rediscache.Cache
}

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service
type Service interface {
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	Authenticate(ctx context.Context, email string, password string) (jwt.RegisteredClaims, error)
	CreateJob(ctx context.Context, nj models.NewJobRequest, cId int) (models.NewJobResponse, error)
	ViewJob(ctx context.Context) ([]models.Job, error)
	GetJobInfoByID(ctx context.Context, jId int) (models.Job, error)
	ViewJobByCompanyId(ctx context.Context, cId int) ([]models.Job, error)
	CreateCompany(ctx context.Context, ni models.NewCompany) (models.Company, error)
	ViewCompany(ctx context.Context) ([]models.Company, error)
	GetCompanyInfoByID(ctx context.Context, uid int) (models.Company, error)

	ProcessJob(ctx context.Context, nj []models.ApplicationRequest) ([]models.ApplicationRequest, error)

	CheckEmail(ctx context.Context, passwordRequest models.UserRequest) (models.Response, error)

	CheckOtpResponse(ctx context.Context, otpVerification models.CheckOtp) (models.Response, error)
}

func NewServiceStore(s repository.Repository, r rediscache.Cache) Service {
	return &NewService{rp: s, rdb: r}
}

// func NewServiceRedis(r *redis.Client) Service {
// 	return &NewService{redis: r}
// }
