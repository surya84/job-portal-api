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

func TestNewService_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		nu  models.NewUser
	}
	tests := []struct {
		name string
		//r                NewService
		args             args
		want             models.User
		wantErr          bool
		mockRepoResponse func() (models.User, error)
	}{
		{
			name: "error in creating",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:     "surya",
					Email:    "surya@gmail.com",
					Password: "1234",
				},
			},
			want: models.User{},
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("error in creating user in db")
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:     "surya",
					Email:    "surya@gmail.com",
					Password: "1234",
				},
			},
			want: models.User{
				Name:  "surya",
				Email: "surya@gmail.com",
			},
			mockRepoResponse: func() (models.User, error) {
				return models.User{
					Name:  "surya",
					Email: "surya@gmail.com",
				}, nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateU(tt.args.ctx, tt.args.nu).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s := NewServiceStore(mockRepo)
			got, err := s.CreateUser(tt.args.ctx, tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}