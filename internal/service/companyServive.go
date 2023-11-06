package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateCompany(ctx context.Context, ni models.NewCompany) (models.Company, error) {
	CompanyData, err := r.rp.CreateC(ctx, ni)
	if err != nil {
		return models.Company{}, err
	}
	return CompanyData, err
}

func (r NewService) ViewCompany(ctx context.Context) ([]models.Company, error) {
	CompanyData, err := r.rp.ViewCompanies()
	if err != nil {
		return []models.Company{}, err
	}
	return CompanyData, err
}

func (r NewService) GetCompanyInfoByID(ctx context.Context, uid int) (models.Company, error) {
	CompanyData, err := r.rp.GetCompanyByID(uid)
	if err != nil {
		return models.Company{}, err
	}
	return CompanyData, err
}
