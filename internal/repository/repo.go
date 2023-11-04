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

//go:generate mockgen -source=repo.go -destination=repository_mock.go -package=repository

type Repository interface {
	CreateU(ctx context.Context, nu models.NewUser) (models.User, error)
	AuthenticateUser(ctx context.Context, email, password string) (jwt.RegisteredClaims, error)
	CreateJ(ctx context.Context, nj models.NewJob, cId int) (models.Job, error)
	ViewJobs(ctx context.Context) ([]models.Job, error)
	GetJobById(jId int) (models.Job, error)
	ViewJobById(cId int) ([]models.Job, error)
	CreateC(ctx context.Context, nc models.NewCompany, userId uint) (models.Company, error)
	ViewCompanies(ctx context.Context) ([]models.Company, error)
	GetCompanyByID(uid int) (models.Company, error)
}
type RepoStore struct {
	Repository
}

func NewRepoStore(r Repository) RepoStore {
	return RepoStore{Repository: r}
}

// NewService is the constructor for the Conn struct.
func NewRepo(db *gorm.DB) (*Conn, error) {

	// We check if the database instance is nil, which would indicate an issue.
	if db == nil {
		return nil, errors.New("please provide a valid connection")
	}

	// We initialize our service with the passed database instance.
	s := &Conn{db: db}
	return s, nil
}
