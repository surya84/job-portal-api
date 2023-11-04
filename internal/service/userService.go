package service

import (
	"context"
	"job-portal/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func (r NewService) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	userData, err := r.rs.CreateU(ctx, nu)
	if err != nil {
		return models.User{}, err
	}
	return userData, err
}
func (r NewService) UserSignin(ctx context.Context, email, password string) (jwt.RegisteredClaims, error) {
	AuthenticateUser, err := r.rs.AuthenticateUser(ctx, email, password)
	if err != nil {
		return jwt.RegisteredClaims{}, nil
	}
	return AuthenticateUser, err
}
