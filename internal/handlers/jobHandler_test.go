package handlers

import (
	"context"
	"job-portal/internal/auth"
	"job-portal/internal/middleware"
	"job-portal/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Test_handler_AddJob(t *testing.T) {

	tests := []struct {
		name string
		//h    *handler
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.NewService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, service.NewService{}
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error log in",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "123")
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest

				return c, rr, service.NewService{}
			},
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   `{"error":"Unauthorized"}`,
		},

		{
			name: "failure to add",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"title":"dev","salary":"1222"}`))
				ctx := httpRequest.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "123")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Request = httpRequest
				// c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				// mc := gomock.NewController(t)
				// ms := service.NewMockService(mc)

				// ms.EXPECT().CreateJob(gomock.Any(), gomock.Any(), gomock.Any()).Return(models.Job{}, nil).AnyTimes()

				return c, rr, service.NewService{}
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":"Bad Request"}`,
		},

		// {
		// 	name: "conversion error",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		// 		c.Request = httpRequest

		// 		return c, rr, service.NewService{}
		// 	},
		// 	expectedStatusCode: http.StatusBadRequest,
		// 	expectedResponse:   `{"msg":"Bad Request"}`,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &handler{
				S: ms,
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
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.NewService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, service.NewService{}
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},

		// {
		// 	name: "success",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
		// 		ctx := httpRequest.Context()
		// 		ctx = context.WithValue(ctx, middleware.TraceIdKey, "123")
		// 		ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
		// 		httpRequest = httpRequest.WithContext(ctx)
		// 		c.Request = httpRequest
		// 		c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
		// 		mc := gomock.NewController(t)
		// 		ms := service.NewMockService(mc)

		// 		ms.EXPECT().ViewJob(ctx).Return([]models.Job{}, nil).AnyTimes()

		// 		return c, rr, service.NewService{}
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	expectedResponse:   `[]`,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &handler{
				S: ms,
			}

			h.AddJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

		})
	}
}

func Test_handler_ViewJobById(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.NewService)
		expectedStatusCode int
		expectedResponse   string
	}{
		{

			name: "conversion error",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, service.NewService{}
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":"Bad Request"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &handler{
				S: ms,
			}

			h.AddJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

			//tt.h.ViewJobById(tt.args.c)
		})
	}
}

func Test_handler_ViewJobByCompany(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.NewService)
		expectedStatusCode int
		expectedResponse   string
	}{

		{

			name: "conversion error",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, rr, service.NewService{}
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":"Bad Request"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()

			h := &handler{
				S: ms,
			}

			h.AddJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())

			//tt.h.ViewJobById(tt.args.c)
		})
	}
}
