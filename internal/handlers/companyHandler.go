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

func (h *handler) CreateCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	claims, ok := ctx.Value(auth.Key).(jwt.RegisteredClaims)
	if !ok {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}

	var newCom models.NewCompany
	err := json.NewDecoder(c.Request.Body).Decode(&newCom)
	if err != nil {
		log.Error().Str("Trace Id", traceId).Msg("login first")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
		return
	}
	validate := validator.New()
	err = validate.Struct(newCom)

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Send()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"msg": "please provide valid details"})
		return
	}
	uid, err := strconv.ParseUint(claims.Subject, 10, 64)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.
			StatusInternalServerError)})
		return
	}
	com, err := h.S.CreateCompany(ctx, newCom, uint(uid))
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Company creation failed"})
		return
	}

	c.JSON(http.StatusOK, com)

}
func (h *handler) ViewCompany(c *gin.Context) {
	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	data, err := h.S.ViewCompany()

	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing C"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *handler) GetCompanyById(c *gin.Context) {

	ctx := c.Request.Context()
	traceId, ok := ctx.Value(middleware.TraceIdKey).(string)
	if !ok {
		log.Error().Msg("traceId missing from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	id := c.Param("id")
	cId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": http.StatusText(http.StatusBadRequest)})
		return
	}

	company, err := h.S.GetCompanyInfoByID(cId)
	if err != nil {

		log.Error().Err(err).Str("Trace Id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "problem in viewing Company"})
		return
	}

	c.JSON(http.StatusOK, company)
}
