package handlers

import (
	"encoding/json"
	"fmt"
	"job-portal/internal/middleware"
	"job-portal/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

func (h *handler) AddJob(c *gin.Context) {

	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)

	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": http.StatusText((http.StatusInternalServerError))})

		return
	}
	cIdstr := c.Param("id")
	cId, err := strconv.Atoi(cIdstr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": http.StatusText((http.StatusBadRequest))})

		return
	}
	var newJob models.NewJobRequest
	err = json.NewDecoder(c.Request.Body).Decode(&newJob)
	if err != nil {
		log.Info().Msg("error while converting request body to json")
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": http.StatusText(http.StatusBadRequest)})

		return
	}
	validate := validator.New()
	err = validate.Struct(newJob)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("validation failed")
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": http.StatusText(http.StatusBadRequest)})

		return
	}

	job, err := h.s.CreateJob(ctx, newJob, cId)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("error while adding job")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": http.StatusText(http.StatusInternalServerError)})

		return
	}

	c.JSON(http.StatusOK, job)

}

func (h *handler) ViewJobs(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": http.StatusText(http.StatusInternalServerError)})

		return
	}

	data, err := h.s.ViewJob(ctx)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": http.StatusText(http.StatusInternalServerError)})

		return
	}

	c.JSON(http.StatusOK, data)

}
func (h *handler) ViewJobById(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": http.StatusText(http.StatusInternalServerError)})

		return
	}

	id := c.Param("id")
	jId, err := strconv.Atoi(id)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})

		return
	}
	job, err := h.s.GetJobInfoByID(ctx, jId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, job)
}
func (h *handler) ViewJobByCompany(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": http.StatusText(http.StatusInternalServerError)})

		return
	}
	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})

		return
	}

	jobs, err := h.s.ViewJobByCompanyId(ctx, cId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func (h *handler) ProcessJobApplication(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	fmt.Println("trace id", traceId)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": http.StatusText(http.StatusInternalServerError)})

		return
	}

	//fmt.Println("///////////////////////////")
	var newApplication []models.ApplicationRequest
	err := json.NewDecoder(c.Request.Body).Decode(&newApplication)
	//fmt.Println("[[[[[[[[[[[[[[]]]]]]]]]]]]]]", err, newApplication)
	if err != nil {
		log.Error().Msg("error while converting requested body to json")
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": http.StatusText(http.StatusInternalServerError)})

		return
	}
	//fmt.Println("dsgdgedfg---------------------")
	for _, val := range newApplication {
		//fmt.Println("-----------------------------")
		validate := validator.New()
		if err := validate.Struct(val); err != nil {
			log.Error().Err(err).Str("Trace Id", traceId).Msg("error while converting to struct")
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError))

			return
		}

	}

	jobData, err := h.s.ProcessJob(ctx, newApplication)
	//fmt.Println("=========================================================", err)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("error while applying job")
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": http.StatusText(http.StatusBadRequest)})

		return
	}

	c.IndentedJSON(http.StatusOK, jobData)

}
