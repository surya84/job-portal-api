package repository

import (
	"context"
	"errors"
	"job-portal/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Conn struct {

	// db is an instance of the SQLite database.
	db *gorm.DB
}

//go:generate mockgen -source=repo.go -destination=repo_mock.go -package=repository
type Repository interface {
	CreateU(ctx context.Context, nu models.NewUser) (models.User, error)
	AuthenticateUser(ctx context.Context, email string, password string) (jwt.RegisteredClaims, error)
	CreateJ(ctx context.Context, nj models.NewJobRequest, cId int) (models.NewJobResponse, error)
	ViewJobs() ([]models.Job, error)
	GetJobById(jId int) (models.Job, error)
	ViewJobById(cId int) ([]models.Job, error)
	CreateC(ctx context.Context, nc models.NewCompany) (models.Company, error)
	ViewCompanies() ([]models.Company, error)
	GetCompanyByID(uid int) (models.Company, error)

	GetJobProcessData(id int) (models.Job, error)

	CheckUserData(ctx context.Context, email string) bool

	//GetDataFromRedis(jid uint)
}

func NewRepo(db *gorm.DB) (Repository, error) {
	if db == nil {
		return nil, errors.New("db cannot be nil")
	}
	return &Conn{db: db}, nil
}
