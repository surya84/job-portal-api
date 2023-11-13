package handlers

import (
	"bytes"
	"context"
	"errors"
	"job-portal/internal/middleware"
	"job-portal/internal/models"
	"job-portal/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func Test_handler_ViewJobByCompany(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "Invalid companyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "surya"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while fectching job details by companyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ViewJobByCompanyId(c.Request.Context(), gomock.Any()).Return([]models.Job{}, errors.New("mock service error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"mock service error"}`,
		},
		{
			name: "sucess while fectching job details by companyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ViewJobByCompanyId(c.Request.Context(), gomock.Any()).Return([]models.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.ViewJobByCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_ViewJobById(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "Invalid jobId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "surya"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while fectching job details by jobId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().GetJobInfoByID(c.Request.Context(), gomock.Any()).Return(models.Job{}, errors.New("mock service error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "sucess while fectching job details by jobId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().GetJobInfoByID(c.Request.Context(), gomock.Any()).Return(models.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"ID":0,"Title":"","Description":"","CompanyID":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.ViewJobById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_AddJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{

		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "Invalid CompanyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "surya"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "invalid request body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := "invalid string request body"
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", strings.NewReader(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "checking validator function",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				//requestBody := []byte(`{"key": "value"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBufferString(`
				{
					"title": "c developer",
					"budget": 85000,
					"description": "software development",
					"min_exp": 2.5,
					"max_exp": 5.5,
					"qualifications": [1,2],
					"shifts": [1, 2],
					"JobTypes": [1, 2,3],
					"locations": [1, 2, 3],
					"technologies": [1, 2, 3],
					"WorkModes": [1, 2, 3]
				}`))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while adding job",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				//requestBody := []byte(`{"title": "GoLang Developer", "description":"train hire and deploy"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBufferString(`
				{
					"title": "c developer",
					"min_np": 20,
					"max_np": 30,
					"budget": 85000,
					"description": "software development",
					"min_exp": 2.5,
					"max_exp": 5.5,
					"qualifications": [1,2],
					"shifts": [1, 2],
					"JobTypes": [1, 2,3],
					"locations": [1, 2, 3],
					"technologies": [1, 2, 3],
					"WorkModes": [1, 2, 3]
				}`))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().CreateJob(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.NewJobResponse{}, errors.New("error in adding job")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "sucessfully adding job",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				//requestBody := []byte(`{"title": "GoLang Developer", "description":"train hire and deploy"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBufferString(`
				{
					"title": "c developer",
					"min_np": 20,
					"max_np": 30,
					"budget": 85000,
					"description": "software development",
					"min_exp": 2.5,
					"max_exp": 5.5,
					"qualifications": [1,2],
					"shifts": [1, 2],
					"JobTypes": [1, 2,3],
					"locations": [1, 2, 3],
					"technologies": [1, 2, 3],
					"WorkModes": [1, 2, 3]
				}`))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().CreateJob(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.NewJobResponse{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.AddJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_ViewJobs(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error while fectching jobs",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ViewJob(c.Request.Context()).Return([]models.Job{}, errors.New("jobs not found")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"jobs not found"}`,
		},
		{
			name: "sucess while fectching jobs",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ViewJob(c.Request.Context()).Return([]models.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.ViewJobs(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_ProcessJobApplication(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{

		{
			name: "trace id missing",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				//var d []models.ApplicationRequest
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)

				c.Request = httpReq

				return c, rr, nil

			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error while converting into json",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				// requestBody := []byte(``)
				httpReq, _ := http.NewRequest(http.MethodPost, "http://google.com:8080", bytes.NewBufferString(`[indcbcusv],,,`))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
				httpReq = httpReq.WithContext(ctx)
				//c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "validation failed",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {

				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodPost, "http://google.com:8080", bytes.NewBufferString(`[{
					{
						"name":"dhoni",
						"id":2,
					   "title": "java developer",
					  
					   "budget": 1200,
					   "min_exp": 4.5,
					   "qualifications": [1, 2],
					   "shifts": [1, 2, 3],
					   "jobTypes": [1, 2, 3],
					   "locations": [1, 2, 3],
					   "technologies": [1, 2, 3],
					   "workModes": [1, 2, 3]
				   },
				   {
						"name":"ruthu",
						   "id":3,
					   "title": "java developer",
					   
					   "budget": 85000,
					   "min_exp": 3.5,
					   "qualifications": [1, 2],
					   "shifts": [1, 2, 3],
					   "jobTypes": [1, 2, 3],
					   "locations": [1, 2, 3],
					   "technologies": [1, 2, 3],
					   "workModes": [1, 2, 3]
				   },
   
					{
						"name":"jadeja",
						   "id":3,
					   "title": "java developer",
					  
					   "budget": 1200,
					   "min_exp": 4.5,
					   "qualifications": [1, 2],
					   "shifts": [1, 2, 3],
					   "jobTypes": [1, 2, 3],
					   "locations": [1, 2, 3],
					   "technologies": [1, 2, 3],
					   "workModes": [1, 2, 3]
				   }
			   ]`))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				return c, rr, nil

			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error fetching from db",

			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)

				// Create a new request with a specific URL
				httpReq, _ := http.NewRequest(http.MethodGet, "http://example.com:8080", bytes.NewBufferString(`[
					{
					"name":"dhoni",
					"id":2,
					"title": "java developer",
					"noticePeriod": 20,
					"budget": 1200,
					"min_exp": 4.5,
					"qualifications": [1, 2],
					"shifts": [1, 2, 3],
					"jobTypes": [1, 2, 3],
					"locations": [1, 2, 3],
					"technologies": [1, 2, 3],
					"workModes": [1, 2, 3]
				},
				{
					"name":"ruthu",
					"id":3,
					"title": "java developer",
					"noticePeriod": 20,
					"budget": 85000,
					"min_exp": 3.5,
					"qualifications": [1, 2],
					"shifts": [1, 2, 3],
					"jobTypes": [1, 2, 3],
					"locations": [1, 2, 3],
					"technologies": [1, 2, 3],
					"workModes": [1, 2, 3]
				},
				 {
					"name":"jadeja",
					"id":3,
					"title": "java developer",
					"noticePeriod": 20,
					"budget": 1200,
					"min_exp": 4.5,
					"qualifications": [1, 2],
					"shifts": [1, 2, 3],
					"jobTypes": [1, 2, 3],
					"locations": [1, 2, 3],
					"technologies": [1, 2, 3],
					"workModes": [1, 2, 3]
				}

				]`))
				ctx := context.WithValue(httpReq.Context(), middleware.TraceIdKey, "1")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ProcessJob(gomock.Any(), gomock.Any()).Return([]models.ApplicationRequest{}, errors.New("jobs not found")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{

			name: "sucess while applying jobs",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBufferString(
					`[
						{
						 "name":"dhoni",
							"id":2,
						"title": "java developer",
						"noticePeriod": 20,
						"budget": 1200,
						"min_exp": 4.5,
						"qualifications": [1, 2],
						"shifts": [1, 2, 3],
						"jobTypes": [1, 2, 3],
						"locations": [1, 2, 3],
						"technologies": [1, 2, 3],
						"workModes": [1, 2, 3]
					},
					{
						 "name":"ruthu",
							"id":3,
						"title": "java developer",
						"noticePeriod": 20,
						"budget": 85000,
						"min_exp": 3.5,
						"qualifications": [1, 2],
						"shifts": [1, 2, 3],
						"jobTypes": [1, 2, 3],
						"locations": [1, 2, 3],
						"technologies": [1, 2, 3],
						"workModes": [1, 2, 3]
					},

					 {
						 "name":"jadeja",
							"id":3,
						"title": "java developer",
						"noticePeriod": 20,
						"budget": 1200,
						"min_exp": 4.5,
						"qualifications": [1, 2],
						"shifts": [1, 2, 3],
						"jobTypes": [1, 2, 3],
						"locations": [1, 2, 3],
						"technologies": [1, 2, 3],
						"workModes": [1, 2, 3]
					}

					]`,
				))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "1")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ProcessJob(c.Request.Context(), gomock.Any()).Return([]models.ApplicationRequest{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.ProcessJobApplication(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

		})
	}
}
