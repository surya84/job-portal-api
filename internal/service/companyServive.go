package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateCompany(ctx context.Context, ni models.NewCompany, userId uint) (models.Company, error) {
	company, err := r.rs.CreateC(ctx, ni, userId)
	if err != nil {
		return models.Company{}, err
	}
	return company, err
}

func (r NewService) ViewCompany() ([]models.Company, error) {
	companyData, err := r.rs.ViewCompanies()

	if err != nil {
		return []models.Company{}, err
	}
	return companyData, err
}

func (r NewService) GetCompanyInfoByID(uid int) (models.Company, error) {
	companyData, err := r.rs.GetCompanyByID(uid)
	if err != nil {
		return models.Company{}, err
	}
	return companyData, err
}
