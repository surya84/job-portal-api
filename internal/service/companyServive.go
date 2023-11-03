package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateCompany(ctx context.Context, ni models.NewCompany, userId uint) (models.Company, error) {
	c, err := r.rs.CreateC(ctx, ni, userId)
	return c, err
}

func (r NewService) ViewCompany() ([]models.Company, error) {
	c, err := r.rs.ViewCompanies()
	return c, err
}

func (r NewService) GetCompanyInfoByID(uid int) (models.Company, error) {
	c, err := r.rs.GetCompanyByID(uid)
	return c, err
}
