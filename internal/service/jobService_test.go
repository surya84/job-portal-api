package service

import (
	"context"
	"errors"
	"job-portal/internal/models"
	rediscache "job-portal/internal/redisCache"
	"job-portal/internal/repository"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestNewService_ViewJob(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name             string
		r                NewService
		args             args
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "error in db",
			want: []models.Job{},
			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{}, errors.New("error in accessing the db")
			},
			wantErr: true,
		},
		{
			name: "success",
			want: []models.Job{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Title:              "software",
					Min_NoticePeriod:   12,
					Max_NoticePeriod:   3,
					Budget:             23542345,
					Description:        "jeriojgiorejhgioejgoi",
					Minimum_Experience: 3,
					Maximum_Experience: 2,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "dsjhfhsifh",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "hnurhiushf",
						},
					},
					Shifts: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "nbfduihfiuehf",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "fregfrgreg",
						},
					},
					JobTypes: []models.JobType{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "fdsdrrrgv",
						},
					},
					Locations: []models.Location{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "terffefred",
						},
					},
					Technologies: []models.Technology{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "vrfefr",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "fdgaegaergr",
						},
					},
					WorkModes: []models.WorkMode{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "rgdargg",
						},
					},
					CompanyId: 1,
				},
				{
					Model: gorm.Model{
						ID: 2,
					},
					Title:              "hardware",
					Min_NoticePeriod:   12,
					Max_NoticePeriod:   3,
					Budget:             458347589,
					Description:        "hhsdhd",
					Minimum_Experience: 3,
					Maximum_Experience: 2,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "dsjhfhsifh",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "hnurhiushf",
						},
					},
					Shifts: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "nbfduihfiuehf",
						},
					},
					JobTypes: []models.JobType{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "fdsdrrrgv",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "vfgrdfrs",
						},
					},
					Locations: []models.Location{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "terffefred",
						},
					},
					Technologies: []models.Technology{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "vrfefr",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "fdgaegaergr",
						},
					},
					WorkModes: []models.WorkMode{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "rgdargg",
						},
					},
					CompanyId: 2,
				},
			},

			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Model: gorm.Model{
							ID: 1,
						},
						Title:              "software",
						Min_NoticePeriod:   12,
						Max_NoticePeriod:   3,
						Budget:             23542345,
						Description:        "jeriojgiorejhgioejgoi",
						Minimum_Experience: 3,
						Maximum_Experience: 2,
						Qualifications: []models.Qualification{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "dsjhfhsifh",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "hnurhiushf",
							},
						},
						Shifts: []models.Shift{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "nbfduihfiuehf",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "fregfrgreg",
							},
						},
						JobTypes: []models.JobType{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "fdsdrrrgv",
							},
						},
						Locations: []models.Location{
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "terffefred",
							},
						},
						Technologies: []models.Technology{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "vrfefr",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "fdgaegaergr",
							},
						},
						WorkModes: []models.WorkMode{
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "rgdargg",
							},
						},
						CompanyId: 1,
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						Title:              "hardware",
						Min_NoticePeriod:   12,
						Max_NoticePeriod:   3,
						Budget:             458347589,
						Description:        "hhsdhd",
						Minimum_Experience: 3,
						Maximum_Experience: 2,
						Qualifications: []models.Qualification{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "dsjhfhsifh",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "hnurhiushf",
							},
						},
						Shifts: []models.Shift{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "nbfduihfiuehf",
							},
						},
						JobTypes: []models.JobType{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "fdsdrrrgv",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "vfgrdfrs",
							},
						},
						Locations: []models.Location{
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "terffefred",
							},
						},
						Technologies: []models.Technology{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "vrfefr",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "fdgaegaergr",
							},
						},
						WorkModes: []models.WorkMode{
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "rgdargg",
							},
						},
						CompanyId: 2,
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)
			mockCache := rediscache.NewMockCache(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobs().Return(tt.mockRepoResponse()).AnyTimes()

			}

			s := NewServiceStore(mockRepo, mockCache)

			got, err := s.ViewJob(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.ViewJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.ViewJob() = %v, want %v", got, tt.want)
			}

			// assert.Equal(t, tt.want, got)
			// assert.Equal(t, tt.wantErr, err)
		})
	}

}

func TestNewService_GetJobInfoByID(t *testing.T) {
	type args struct {
		ctx context.Context
		jId int
	}
	tests := []struct {
		name             string
		r                NewService
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "error from db",
			want: models.Job{},
			args: args{
				jId: 12,
			},
			wantErr: true,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("test error")
			},
		},
		{
			name: "success",
			args: args{
				jId: 12,
			},
			want: models.Job{
				Model: gorm.Model{
					ID: 12,
				},
				Title:              "software",
				Min_NoticePeriod:   12,
				Max_NoticePeriod:   3,
				Budget:             23542345,
				Description:        "jeriojgiorejhgioejgoi",
				Minimum_Experience: 3,
				Maximum_Experience: 2,
				Qualifications: []models.Qualification{
					{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "dsjhfhsifh",
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						Name: "hnurhiushf",
					},
				},
				Shifts: []models.Shift{
					{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "nbfduihfiuehf",
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						Name: "fregfrgreg",
					},
				},
				JobTypes: []models.JobType{
					{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "fdsdrrrgv",
					},
				},
				Locations: []models.Location{
					{
						Model: gorm.Model{
							ID: 2,
						},
						Name: "terffefred",
					},
				},
				Technologies: []models.Technology{
					{
						Model: gorm.Model{
							ID: 1,
						},
						Name: "vrfefr",
					},
					{
						Model: gorm.Model{
							ID: 2,
						},
						Name: "fdgaegaergr",
					},
				},
				WorkModes: []models.WorkMode{
					{
						Model: gorm.Model{
							ID: 2,
						},
						Name: "rgdargg",
					},
				},
				CompanyId: 1,
			},
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					Model: gorm.Model{
						ID: 12,
					},
					Title:              "software",
					Min_NoticePeriod:   12,
					Max_NoticePeriod:   3,
					Budget:             23542345,
					Description:        "jeriojgiorejhgioejgoi",
					Minimum_Experience: 3,
					Maximum_Experience: 2,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "dsjhfhsifh",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "hnurhiushf",
						},
					},
					Shifts: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "nbfduihfiuehf",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "fregfrgreg",
						},
					},
					JobTypes: []models.JobType{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "fdsdrrrgv",
						},
					},
					Locations: []models.Location{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "terffefred",
						},
					},
					Technologies: []models.Technology{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "vrfefr",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "fdgaegaergr",
						},
					},
					WorkModes: []models.WorkMode{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "rgdargg",
						},
					},
					CompanyId: 1,
				}, nil
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)
			mockCache := rediscache.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().GetJobById(tt.args.jId).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s := NewServiceStore(mockRepo, mockCache)
			got, err := s.GetJobInfoByID(tt.args.ctx, tt.args.jId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.GetJobInfoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.GetJobInfoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService_ViewJobByCompanyId(t *testing.T) {
	type args struct {
		ctx context.Context
		cId int
	}
	tests := []struct {
		name             string
		r                NewService
		args             args
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "error in db",
			args: args{
				cId: 12,
			},
			want: []models.Job{},

			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{}, errors.New("error in accesing data from db")
			},

			wantErr: true,
		},

		{
			name: "success",
			args: args{
				cId: 12,
			},
			want: []models.Job{
				{
					Model: gorm.Model{
						ID: 1,
					},
					Title:              "software",
					Min_NoticePeriod:   12,
					Max_NoticePeriod:   3,
					Budget:             23542345,
					Description:        "jeriojgiorejhgioejgoi",
					Minimum_Experience: 3,
					Maximum_Experience: 2,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "dsjhfhsifh",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "hnurhiushf",
						},
					},
					Shifts: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "nbfduihfiuehf",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "fregfrgreg",
						},
					},
					JobTypes: []models.JobType{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "fdsdrrrgv",
						},
					},
					Locations: []models.Location{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "terffefred",
						},
					},
					Technologies: []models.Technology{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "vrfefr",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "fdgaegaergr",
						},
					},
					WorkModes: []models.WorkMode{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "rgdargg",
						},
					},
					CompanyId: 12,
				},
				{
					Model: gorm.Model{
						ID: 12,
					},
					Title:              "hardware",
					Min_NoticePeriod:   12,
					Max_NoticePeriod:   3,
					Budget:             458347589,
					Description:        "hhsdhd",
					Minimum_Experience: 3,
					Maximum_Experience: 2,
					Qualifications: []models.Qualification{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "dsjhfhsifh",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "hnurhiushf",
						},
					},
					Shifts: []models.Shift{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "nbfduihfiuehf",
						},
					},
					JobTypes: []models.JobType{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "fdsdrrrgv",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "vfgrdfrs",
						},
					},
					Locations: []models.Location{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "terffefred",
						},
					},
					Technologies: []models.Technology{
						{
							Model: gorm.Model{
								ID: 1,
							},
							Name: "vrfefr",
						},
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "fdgaegaergr",
						},
					},
					WorkModes: []models.WorkMode{
						{
							Model: gorm.Model{
								ID: 2,
							},
							Name: "rgdargg",
						},
					},
					CompanyId: 12,
				},
			},

			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Model: gorm.Model{
							ID: 1,
						},
						Title:              "software",
						Min_NoticePeriod:   12,
						Max_NoticePeriod:   3,
						Budget:             23542345,
						Description:        "jeriojgiorejhgioejgoi",
						Minimum_Experience: 3,
						Maximum_Experience: 2,
						Qualifications: []models.Qualification{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "dsjhfhsifh",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "hnurhiushf",
							},
						},
						Shifts: []models.Shift{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "nbfduihfiuehf",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "fregfrgreg",
							},
						},
						JobTypes: []models.JobType{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "fdsdrrrgv",
							},
						},
						Locations: []models.Location{
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "terffefred",
							},
						},
						Technologies: []models.Technology{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "vrfefr",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "fdgaegaergr",
							},
						},
						WorkModes: []models.WorkMode{
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "rgdargg",
							},
						},
						CompanyId: 12,
					},
					{
						Model: gorm.Model{
							ID: 12,
						},
						Title:              "hardware",
						Min_NoticePeriod:   12,
						Max_NoticePeriod:   3,
						Budget:             458347589,
						Description:        "hhsdhd",
						Minimum_Experience: 3,
						Maximum_Experience: 2,
						Qualifications: []models.Qualification{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "dsjhfhsifh",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "hnurhiushf",
							},
						},
						Shifts: []models.Shift{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "nbfduihfiuehf",
							},
						},
						JobTypes: []models.JobType{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "fdsdrrrgv",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "vfgrdfrs",
							},
						},
						Locations: []models.Location{
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "terffefred",
							},
						},
						Technologies: []models.Technology{
							{
								Model: gorm.Model{
									ID: 1,
								},
								Name: "vrfefr",
							},
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "fdgaegaergr",
							},
						},
						WorkModes: []models.WorkMode{
							{
								Model: gorm.Model{
									ID: 2,
								},
								Name: "rgdargg",
							},
						},
						CompanyId: 12,
					},
				}, nil
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)

			mockRepo := repository.NewMockRepository(mc)
			mockCache := rediscache.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobById(tt.args.cId).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s := NewServiceStore(mockRepo, mockCache)
			got, err := s.ViewJobByCompanyId(tt.args.ctx, tt.args.cId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.ViewJobByCompanyId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.ViewJobByCompanyId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService_ProcessJob(t *testing.T) {
	type args struct {
		ctx    context.Context
		newJob []models.ApplicationRequest
	}
	tests := []struct {
		name    string
		r       NewService
		args    args
		want    []models.ApplicationRequest
		wantErr bool
		setup   func(mockRepo *repository.MockRepository, mockCache *rediscache.MockCache)
	}{
		// {
		// 	name: "error case",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		newJob: []models.ApplicationRequest{

		// 			//budget fail case
		// 			{
		// 				Name:           "surya",
		// 				Id:             3,
		// 				Title:          "java developer",
		// 				NoticePeriod:   intPtr(30),
		// 				Budget:         2500000,
		// 				Experience:     3.5,
		// 				Qualifications: []uint{1, 2},
		// 				Shifts:         []uint{1, 2, 3},
		// 				JobTypes:       []uint{1, 2},
		// 				Locations:      []uint{1, 2},
		// 				Technologies:   []uint{1, 2},
		// 				WorkModes:      []uint{1, 2},
		// 			},
		// 			{
		// 				Name:           "teja",
		// 				Id:             4,
		// 				Title:          "data science",
		// 				NoticePeriod:   intPtr(30),
		// 				Budget:         25000,
		// 				Experience:     3.5,
		// 				Qualifications: []uint{1, 2},
		// 				Shifts:         []uint{1, 2, 3},
		// 				JobTypes:       []uint{1, 2},
		// 				Locations:      []uint{1, 2},
		// 				Technologies:   []uint{1, 2},
		// 				WorkModes:      []uint{1, 2},
		// 			},

		// 			//notice period fail case
		// 			{
		// 				Name:           "Ram",
		// 				Id:             2,
		// 				Title:          "data science",
		// 				NoticePeriod:   intPtr(1000),
		// 				Budget:         25000,
		// 				Experience:     3.5,
		// 				Qualifications: []uint{1, 2},
		// 				Shifts:         []uint{1, 2, 3},
		// 				JobTypes:       []uint{1, 2},
		// 				Locations:      []uint{1, 2},
		// 				Technologies:   []uint{1, 2},
		// 				WorkModes:      []uint{1, 2},
		// 			},

		// 			//experience fail case
		// 			{
		// 				Name:           "Lucky",
		// 				Id:             5,
		// 				Title:          "C developer",
		// 				NoticePeriod:   intPtr(30),
		// 				Budget:         25000,
		// 				Experience:     1.2,
		// 				Qualifications: []uint{1, 2},
		// 				Shifts:         []uint{1, 2, 3},
		// 				JobTypes:       []uint{1, 2},
		// 				Locations:      []uint{1, 2},
		// 				Technologies:   []uint{1, 2},
		// 				WorkModes:      []uint{1, 2},
		// 			},
		// 		},
		// 	},
		// 	want:    nil,
		// 	wantErr: false,
		// 	setup: func(mockRepo *repository.MockRepository, mockCache *rediscache.MockCache) {
		// 		mockRepo.EXPECT().GetJobProcessData(3).Return(models.Job{
		// 			Model:              gorm.Model{ID: 2},
		// 			Title:              "java developer",
		// 			Min_NoticePeriod:   10,
		// 			Max_NoticePeriod:   40,
		// 			Budget:             25090,
		// 			Description:        "java development",
		// 			Minimum_Experience: 2.5,
		// 			Maximum_Experience: 5.5,
		// 			Qualifications: []models.Qualification{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 			Shifts: []models.Shift{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 			JobTypes: []models.JobType{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},

		// 			Locations: []models.Location{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 			Technologies: []models.Technology{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 			WorkModes: []models.WorkMode{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 		}, nil).Times(1)
		// 		mockRepo.EXPECT().GetJobProcessData(4).Return(models.Job{

		// 			Model:              gorm.Model{ID: 4},
		// 			Title:              " data science",
		// 			Min_NoticePeriod:   10,
		// 			Max_NoticePeriod:   40,
		// 			Budget:             25090,
		// 			Description:        "work on data sets",
		// 			Minimum_Experience: 2.5,
		// 			Maximum_Experience: 5.5,
		// 			Qualifications: []models.Qualification{
		// 				{Model: gorm.Model{ID: 4}},
		// 			},
		// 			Shifts: []models.Shift{
		// 				{Model: gorm.Model{ID: 4}},
		// 			},
		// 			JobTypes: []models.JobType{
		// 				{Model: gorm.Model{ID: 4}},
		// 			},

		// 			Locations: []models.Location{
		// 				{Model: gorm.Model{ID: 4}},
		// 			},
		// 			Technologies: []models.Technology{
		// 				{Model: gorm.Model{ID: 4}},
		// 			},
		// 			WorkModes: []models.WorkMode{
		// 				{Model: gorm.Model{ID: 4}},
		// 			},
		// 		}, nil).Times(1)

		// 		mockRepo.EXPECT().GetJobProcessData(2).Return(models.Job{

		// 			Model:              gorm.Model{ID: 2},
		// 			Title:              " data science",
		// 			Min_NoticePeriod:   30,
		// 			Max_NoticePeriod:   60,
		// 			Budget:             25090,
		// 			Description:        "work on data sets",
		// 			Minimum_Experience: 2.5,
		// 			Maximum_Experience: 5.5,
		// 			Qualifications: []models.Qualification{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 			Shifts: []models.Shift{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 			JobTypes: []models.JobType{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},

		// 			Locations: []models.Location{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 			Technologies: []models.Technology{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 			WorkModes: []models.WorkMode{
		// 				{Model: gorm.Model{ID: 2}},
		// 			},
		// 		}, nil).Times(1)

		// 		mockRepo.EXPECT().GetJobProcessData(5).Return(models.Job{

		// 			Model:              gorm.Model{ID: 5},
		// 			Title:              "C developer",
		// 			Min_NoticePeriod:   30,
		// 			Max_NoticePeriod:   40,
		// 			Budget:             25090,
		// 			Description:        "work on data sets",
		// 			Minimum_Experience: 2.5,
		// 			Maximum_Experience: 5.5,
		// 			Qualifications: []models.Qualification{
		// 				{Model: gorm.Model{ID: 5}},
		// 			},
		// 			Shifts: []models.Shift{
		// 				{Model: gorm.Model{ID: 5}},
		// 			},
		// 			JobTypes: []models.JobType{
		// 				{Model: gorm.Model{ID: 5}},
		// 			},

		// 			Locations: []models.Location{
		// 				{Model: gorm.Model{ID: 5}},
		// 			},
		// 			Technologies: []models.Technology{
		// 				{Model: gorm.Model{ID: 5}},
		// 			},
		// 			WorkModes: []models.WorkMode{
		// 				{Model: gorm.Model{ID: 5}},
		// 			},
		// 		}, nil).Times(1)
		// 	},
		// },

		{
			name: "error from redis",
			args: args{
				ctx: context.Background(),
				newJob: []models.ApplicationRequest{

					{
						Name:           "surya",
						Id:             3,
						Title:          "java developer",
						NoticePeriod:   intPtr(20),
						Budget:         25000,
						Experience:     3.5,
						Qualifications: []uint{1, 2},
						Shifts:         []uint{1, 2, 3},
						JobTypes:       []uint{1, 2, 3},
						Locations:      []uint{1, 2, 3},
						Technologies:   []uint{1, 2, 3},
						WorkModes:      []uint{1, 2, 3},
					},
					// {
					// 	Name:           "teja",
					// 	Id:             10,
					// 	Title:          "java developer",
					// 	NoticePeriod:   intPtr(30),
					// 	Budget:         2500000,
					// 	Experience:     3.5,
					// 	Qualifications: []uint{1, 2},
					// 	Shifts:         []uint{1, 2, 3},
					// 	JobTypes:       []uint{1, 2},
					// 	Locations:      []uint{1, 2},
					// 	Technologies:   []uint{1, 2},
					// 	WorkModes:      []uint{1, 2},
					// },
				},
			},
			want: []models.ApplicationRequest{
				{
					Name:           "surya",
					Id:             3,
					Title:          "java developer",
					NoticePeriod:   intPtr(20),
					Budget:         25000,
					Experience:     3.5,
					Qualifications: []uint{1, 2},
					Shifts:         []uint{1, 2, 3},
					JobTypes:       []uint{1, 2, 3},
					Locations:      []uint{1, 2, 3},
					Technologies:   []uint{1, 2, 3},
					WorkModes:      []uint{1, 2, 3},
				},
			},
			wantErr: false,
			setup: func(mockRepo *repository.MockRepository, mockCache *rediscache.MockCache) {
				//mockRepo.EXPECT().GetJobProcessData(10).Return(models.Job{}, errors.New("no data found for job id")).Times(1)
				mockCache.EXPECT().CheckRedisKey(gomock.Any()).Return(models.Job{
					Model:              gorm.Model{ID: 3},
					Title:              "java developer",
					Min_NoticePeriod:   20,
					Max_NoticePeriod:   30,
					Budget:             85000,
					Description:        "java development",
					Minimum_Experience: 2.5,
					Maximum_Experience: 5.5,
					Qualifications: []models.Qualification{
						{Model: gorm.Model{ID: 3}},
					},
					Shifts: []models.Shift{
						{Model: gorm.Model{ID: 3}},
					},
					JobTypes: []models.JobType{
						{Model: gorm.Model{ID: 3}},
					},
					Locations: []models.Location{
						{Model: gorm.Model{ID: 3}},
					},
					Technologies: []models.Technology{
						{Model: gorm.Model{ID: 3}},
					},
					WorkModes: []models.WorkMode{
						{Model: gorm.Model{ID: 3}},
					},
				}, nil).Times(1)
			},
		},

		{
			name: "success",
			args: args{
				ctx: context.Background(),
				newJob: []models.ApplicationRequest{

					{
						Name:           "surya",
						Id:             1,
						Title:          "java developer",
						NoticePeriod:   intPtr(20),
						Budget:         25000,
						Experience:     3.5,
						Qualifications: []uint{1, 2},
						Shifts:         []uint{1, 2, 3},
						JobTypes:       []uint{1, 2, 3},
						Locations:      []uint{1, 2, 3},
						Technologies:   []uint{1, 2, 3},
						WorkModes:      []uint{1, 2, 3},
					},
					// {
					// 	Name:           "teja",
					// 	Id:             1,
					// 	Title:          "java developer",
					// 	NoticePeriod:   intPtr(1000),
					// 	Budget:         2500000,
					// 	Experience:     3.5,
					// 	Qualifications: []uint{1, 2},
					// 	Shifts:         []uint{1, 2, 3},
					// 	JobTypes:       []uint{1, 2, 3},
					// 	Locations:      []uint{1, 2, 3},
					// 	Technologies:   []uint{1, 2, 3},
					// 	WorkModes:      []uint{1, 2, 3},
					// },
				},
			},
			want: []models.ApplicationRequest{
				{
					Name:           "surya",
					Id:             1,
					Title:          "java developer",
					NoticePeriod:   intPtr(20),
					Budget:         25000,
					Experience:     3.5,
					Qualifications: []uint{1, 2},
					Shifts:         []uint{1, 2, 3},
					JobTypes:       []uint{1, 2, 3},
					Locations:      []uint{1, 2, 3},
					Technologies:   []uint{1, 2, 3},
					WorkModes:      []uint{1, 2, 3},
				},
			},
			wantErr: false,
			setup: func(mockRepo *repository.MockRepository, mockCache *rediscache.MockCache) {
				mockCache.EXPECT().CheckRedisKey("1").Return(models.Job{}, errors.New("error from cache")).Times(1)
				mockRepo.EXPECT().GetJobProcessData(1).Return(models.Job{
					Model:              gorm.Model{ID: 1},
					Title:              "java developer",
					Min_NoticePeriod:   20,
					Max_NoticePeriod:   30,
					Budget:             25000,
					Description:        "java development",
					Minimum_Experience: 2.5,
					Maximum_Experience: 5.5,
					Qualifications: []models.Qualification{
						{Model: gorm.Model{ID: 1}},
					},
					Shifts: []models.Shift{
						{Model: gorm.Model{ID: 1}},
					},
					JobTypes: []models.JobType{
						{Model: gorm.Model{ID: 1}},
					},
					Locations: []models.Location{
						{Model: gorm.Model{ID: 1}},
					},
					Technologies: []models.Technology{
						{Model: gorm.Model{ID: 1}},
					},
					WorkModes: []models.WorkMode{
						{Model: gorm.Model{ID: 1}},
					},
				}, nil).Times(1)

				mockCache.EXPECT().SetRedisKey("1", models.Job{
					Model:              gorm.Model{ID: 1},
					Title:              "java developer",
					Min_NoticePeriod:   20,
					Max_NoticePeriod:   30,
					Budget:             25000,
					Description:        "java development",
					Minimum_Experience: 2.5,
					Maximum_Experience: 5.5,
					Qualifications: []models.Qualification{
						{Model: gorm.Model{ID: 1}},
					},
					Shifts: []models.Shift{
						{Model: gorm.Model{ID: 1}},
					},
					JobTypes: []models.JobType{
						{Model: gorm.Model{ID: 1}},
					},
					Locations: []models.Location{
						{Model: gorm.Model{ID: 1}},
					},
					Technologies: []models.Technology{
						{Model: gorm.Model{ID: 1}},
					},
					WorkModes: []models.WorkMode{
						{Model: gorm.Model{ID: 1}},
					},
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)
			mockCache := rediscache.NewMockCache(mc)

			tt.setup(mockRepo, mockCache)
			//tt.setupCache(mockCache)
			s := NewServiceStore(mockRepo, mockCache)
			got, err := s.ProcessJob(tt.args.ctx, tt.args.newJob)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.ProcessJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.ProcessJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService_CreateJob(t *testing.T) {
	type args struct {
		ctx    context.Context
		newJob models.NewJobRequest
		cId    int
	}
	tests := []struct {
		name             string
		r                NewService
		args             args
		want             models.NewJobResponse
		wantErr          bool
		mockRepoResponse func() (models.NewJobResponse, error)
	}{
		{
			name: "error in creating job",
			args: args{
				ctx: context.Background(),
				newJob: models.NewJobRequest{
					Title:              "java developer",
					Min_NoticePeriod:   intPtr(20),
					Max_NoticePeriod:   intPtr(40),
					Budget:             25000,
					Description:        "java development",
					Minimum_Experience: 2.2,
					Maximum_Experience: 5.5,
					Qualifications:     []uint{1, 2},
					Shifts:             []uint{1, 2, 3},
					JobTypes:           []uint{1, 2},
					Locations:          []uint{1, 2},
					Technologies:       []uint{1, 2, 3},
					WorkModes:          []uint{1, 2},
				},
			},
			want: models.NewJobResponse{},
			mockRepoResponse: func() (models.NewJobResponse, error) {
				return models.NewJobResponse{}, errors.New("error in creating job")
			},
			wantErr: true,
		},

		{
			name: "success",
			args: args{
				ctx: context.Background(),

				newJob: models.NewJobRequest{
					Title:              "java developer",
					Min_NoticePeriod:   intPtr(20),
					Max_NoticePeriod:   intPtr(40),
					Budget:             25000,
					Description:        "java development",
					Minimum_Experience: 2.2,
					Maximum_Experience: 5.5,
					Qualifications:     []uint{1, 2},
					Shifts:             []uint{1, 2, 3},
					JobTypes:           []uint{1, 2},
					Locations:          []uint{1, 2},
					Technologies:       []uint{1, 2, 3},
					WorkModes:          []uint{1, 2},
				},
			},

			want: models.NewJobResponse{
				ID: uint(2),
			},

			mockRepoResponse: func() (models.NewJobResponse, error) {
				return models.NewJobResponse{
					ID: uint(2),
				}, nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)
			mockCache := rediscache.NewMockCache(mc)

			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateJ(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s := NewServiceStore(mockRepo, mockCache)
			got, err := s.CreateJob(tt.args.ctx, tt.args.newJob, tt.args.cId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.CreateJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.CreateJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func intPtr(i int) *int {
	return &i
}
