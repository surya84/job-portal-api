package handlers

import (
	"job-portal/internal/auth"
	"job-portal/internal/middleware"
	rediscache "job-portal/internal/redisCache"
	"job-portal/internal/repository"
	"job-portal/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func API(a *auth.Auth, c repository.Repository, redis rediscache.Cache) *gin.Engine {

	r := gin.New()

	m, err := middleware.NewMid(a)
	s := service.NewServiceStore(c, redis)
	h := handler{
		a: a,
		s: s,
	}

	if err != nil {
		log.Panic().Msg("middlewares not set up")
	}

	r.Use(m.Log(), gin.Recovery())

	r.POST("/api/register", h.UserRegister)
	r.POST("/api/login", h.UserLogin)
	r.POST("/api/companies", m.Authenticate(h.CreateCompany))
	r.GET("/api/companies", m.Authenticate(h.ViewCompany))
	r.GET("/api/companies/:id", m.Authenticate(h.GetCompanyById))
	r.POST("/api/companies/:id/jobs", m.Authenticate(h.AddJob))
	r.GET("/api/jobs", m.Authenticate(h.ViewJobs))
	r.GET("/api/jobs/:id", m.Authenticate(h.ViewJobById))
	r.GET("/api/companies/:id/jobs", m.Authenticate(h.ViewJobByCompany))
	r.POST("/api/processjobapplication", m.Authenticate(h.ProcessJobApplication))
	r.POST("/api/forgetpassword", h.ForgetPassword)
	r.POST("/api/changePassword", h.ChangePassword)

	// Return the prepared Gin engine
	return r
}
