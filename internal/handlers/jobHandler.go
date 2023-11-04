package handlers

import (
	"encoding/json"
	"job-portal/internal/auth"
	"job-portal/internal/middleware"
	"job-portal/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func (h *handler) AddJob(c *gin.Context) {

	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	//id := c.Param("cid")

	cIdstr := c.Param("id")
	cId, err := strconv.Atoi(cIdstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": http.StatusText((http.StatusBadRequest))})
		return
	}
	var newJob models.NewJob
	err = json.NewDecoder(c.Request.Body).Decode(&newJob)
	if err != nil {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	validate := validator.New()
	err = validate.Struct(newJob)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceid).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"msg": "please provide valid details"})
		return
	}

	job, err := h.S.CreateJob(ctx, newJob, cId)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceid)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Job creation failed"})
		return
	}

	c.JSON(http.StatusOK, job)

}

func (h *handler) ViewJobs(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	data, err := h.S.ViewJob(ctx)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing Jobs"})
		return
	}

	c.JSON(http.StatusOK, data)

}
func (h *handler) ViewJobById(c *gin.Context) {

	id := c.Param("id")
	jId, err := strconv.Atoi(id)
	if err != nil {
		// Handle invalid ID
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Call the service layer to get company information
	job, err := h.S.GetJobInfoByID(jId)
	if err != nil {
		// Handle errors, e.g., company not found
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the company data as JSON response
	c.JSON(http.StatusOK, job)
}
func (h *handler) ViewJobByCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceid, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceid missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	_, ok = ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceid).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Call the service layer to get company information
	jobs, err := h.S.ViewJobByCompanyId(cId)
	if err != nil {
		// Handle errors, e.g., company not found
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Return the company data as JSON response
	c.JSON(http.StatusOK, jobs)
}
