package api

import (
	"net/http"

	"github.com/Mohammed-Aadil/common-core/pkg/pagination"
	"github.com/Mohammed-Aadil/common-core/pkg/response"
	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ListRisksAPI doc
// @Summary 		List all risks available with pagination
// @Description 	List all risks available with pagination
// @Tags 			risks
// @Accept			application/json
// @Produce			application/json
// @Success			200 {string} string "OK"
// @Failure			503
// @Router			/api/v1/risks [get]
func (s *APIHandler) listRisksAPI(c *gin.Context) {
	var paginationPayload pagination.Pagination
	if err := c.ShouldBindQuery(&paginationPayload); err != nil {
		response.JsonResponse(c, response.HttpErrorResponse{Message: "unable to parse pagination from query"}, http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	if paginationPayload.SortField == "" {
		paginationPayload.SortField = "title"
	}
	if paginationPayload.SortOrder == "" {
		paginationPayload.SortOrder = pagination.ASC
	}
	if paginationPayload.Limit == 0 {
		paginationPayload.Limit = s.conf.DefaultPaginationLimit
	}

	risks, total, statusCode, err := s.service.ListRisk(ctx, &paginationPayload)
	if err != nil {
		response.JsonResponse(c, response.HttpErrorResponse{Message: err.Error()}, statusCode)
		return
	}

	resp := response.HttpPaginatedResponse{
		Data:       risks,
		Pagination: paginationPayload,
	}
	resp.Pagination.Total = total

	response.JsonResponse(c, resp, statusCode)
}

// GetRisksAPI doc
// @Summary 		get risk details
// @Description 	get risk details
// @Tags 			risks
// @Accept			application/json
// @Produce			application/json
// @Success			200 {string} string "OK"
// @Failure			503
// @Router			/api/v1/risks/:id [get]
func (s *APIHandler) getRisksAPI(c *gin.Context) {
	riskId := c.Param(model.RiskId)
	ctx := c.Request.Context()

	if riskUUID, err := uuid.Parse(riskId); err != nil {
		response.JsonResponse(c, response.HttpErrorResponse{Message: "incorrect risk id provided"}, http.StatusBadRequest)
		return
	} else {
		risk, httpCode, err := s.service.GetRisk(ctx, riskUUID)
		if err != nil {
			response.JsonResponse(c, response.HttpErrorResponse{Message: err.Error()}, httpCode)
			return
		}
		response.JsonResponse(c, risk, http.StatusOK)
	}
}

// CreateRisksAPI doc
// @Summary 		store risk details in system
// @Description 	store risk details in system
// @Tags 			risks
// @Accept			application/json
// @Produce			application/json
// @Success			200 {string} string "OK"
// @Failure			503
// @Router			/api/v1/risks [post]
func (s *APIHandler) createRisksAPI(c *gin.Context) {
	var risk model.Risk

	if err := c.ShouldBind(&risk); err != nil {
		response.JsonResponse(c, response.HttpErrorResponse{Message: "unable to parse payload"}, http.StatusBadRequest)
		return
	}

	ctx := c.Request.Context()

	httpCode, err := s.service.CreateRisk(ctx, &risk)
	if err != nil {
		response.JsonResponse(c, response.HttpErrorResponse{Message: err.Error()}, httpCode)
		return
	}
	response.JsonResponse(c, nil, httpCode)
}
