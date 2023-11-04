package handlers

import (
	"context"
	"job-portal/internal/auth"
	"job-portal/internal/middleware"
	"job-portal/internal/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Test_handler_ViewCompany(t *testing.T) {

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
				responseRecorder := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(responseRecorder)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, responseRecorder, service.NewService{}
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
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

		// 		mockCtrl := gomock.NewController(t)
		// 		mockService := service.NewMockService(mockCtrl)
		// 		mockService.EXPECT().ViewCompany(gomock.Any()).Return([]models.Company{}, nil).AnyTimes()

		// 		return c, rr, service.NewService{}
		// 	},
		// 	expectedStatusCode: http.StatusOK,
		// 	expectedResponse:   ``,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				S: ms,
			}
			h.ViewCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_GetCompanyById(t *testing.T) {
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
				responseRecorder := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(responseRecorder)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, responseRecorder, service.NewService{}
			},

			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		{
			name: "invalid id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
				responseRecorder := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(responseRecorder)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
				ctx := httpRequest.Context()

				ctx = context.WithValue(ctx, middleware.TraceIdKey, "152")
				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
				httpRequest = httpRequest.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})
				c.Request = httpRequest

				return c, responseRecorder, service.NewService{}

			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"msg":"Bad Request"}`,
		},

		// {
		// 	name: "error in fecting company details from service",
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

		// 		mockCtrl := gomock.NewController(t)
		// 		mockService := service.NewMockService(mockCtrl)
		// 		mockService.EXPECT().GetCompanyInfoByID(c.Request.Context()).Return(models.Company{}, errors.New("error in service")).AnyTimes()

		// 		return c, rr, service.NewService{}
		// 	},

		// 	expectedStatusCode: http.StatusInternalServerError,
		// 	expectedResponse:   `{"error":"test service error"}`,
		// },
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				S: service.NewService{Service: ms},
			}
			h.GetCompanyById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_CreateCompany(t *testing.T) {

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
				responseRecorder := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(responseRecorder)
				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
				c.Request = httpRequest

				return c, responseRecorder, service.NewService{}
			},

			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"msg":"Internal Server Error"}`,
		},
		// {
		// 	name: "missing claims",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.NewService) {
		// 		responseRecorder := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(responseRecorder)
		// 		httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
		// 		c.Request = httpRequest

		// 		return c, responseRecorder, service.NewService{}
		// 	},

		// 	expectedStatusCode: http.StatusUnauthorized,
		// 	expectedResponse:   ``,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				S: service.NewService{Service: ms},
			}
			h.GetCompanyById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
