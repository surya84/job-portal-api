package service

import (
	"context"
	"job-portal/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func (r NewService) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	user, err := r.rs.CreateU(ctx, nu)
	return user, err
}
func (r NewService) Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims, error) {
	c, err := r.rs.AuthenticateUser(ctx, email, password)
	return c, err
}
