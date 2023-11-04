package service

import (
	"context"
	"errors"
	"job-portal/internal/models"
	"job-portal/internal/repository"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestNewService_ViewCompany(t *testing.T) {
	tests := []struct {
		name string
		//r                NewService
		want             []models.Company
		wantErr          bool
		mockRepoResponse func() ([]models.Company, error)
	}{
		{
			name:    "error from db",
			want:    []models.Company{},
			wantErr: true,
			mockRepoResponse: func() ([]models.Company, error) {
				return []models.Company{}, errors.New("test error")
			},
		},

		{
			name: "sucsess",
			want: []models.Company{
				models.Company{
					Name:     "tcs",
					Location: "banglore",
				},
			},
			wantErr: false,
			mockRepoResponse: func() ([]models.Company, error) {
				return []models.Company{
					models.Company{
						Name:     "tcs",
						Location: "banglore",
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			mc := gomock.NewController(t)

			mockRepo := repository.NewMockRepository(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewCompanies(ctx).Return(tt.mockRepoResponse())
			}

			s := NewServiceStore(mockRepo)

			got, err := s.ViewCompany(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.ViewCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.ViewCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService_GetCompanyInfoByID(t *testing.T) {
	type args struct {
		uid int
	}
	tests := []struct {
		name string
		//r                NewService
		args             args
		want             models.Company
		wantErr          bool
		mockRepoResponse func() (models.Company, error)
	}{
		{
			name: "error in db",
			args: args{
				uid: 12,
			},
			want: models.Company{},
			mockRepoResponse: func() (models.Company, error) {
				return models.Company{}, errors.New("error in accessing data from db")

			},
			wantErr: true,
		},

		{
			name: "success",
			args: args{
				uid: 4,
			},
			want: models.Company{
				Name:     "tcs",
				Location: "banglore",
			},

			mockRepoResponse: func() (models.Company, error) {
				return models.Company{
					Name:     "tcs",
					Location: "banglore",
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetCompanyByID(tt.args.uid).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s := NewServiceStore(mockRepo)
			got, err := s.GetCompanyInfoByID(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.GetCompanyInfoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.GetCompanyInfoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService_CreateCompany(t *testing.T) {
	type args struct {
		ctx    context.Context
		ni     models.NewCompany
		userId uint
	}
	tests := []struct {
		name string
		//r       NewService
		args             args
		want             models.Company
		wantErr          bool
		mockRepoResponse func() (models.Company, error)
	}{
		{
			name: "error in creation",
			args: args{
				ctx: context.Background(),
				ni: models.NewCompany{
					Name:     "Tcs",
					Location: "banglore",
				},
				userId: 3,
			},

			want: models.Company{},
			mockRepoResponse: func() (models.Company, error) {

				return models.Company{}, errors.New("error in creation")

			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateC(tt.args.ctx, tt.args.ni, tt.args.userId).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s := NewServiceStore(mockRepo)

			got, err := s.CreateCompany(tt.args.ctx, tt.args.ni, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.CreateCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.CreateCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}
