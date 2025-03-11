package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/Mohammed-Aadil/common-core/pkg/pagination"
	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/google/uuid"
)

func (s *ServiceHandler) ListRisk(ctx context.Context, pagination *pagination.Pagination) ([]*model.Risk, int, int, error) {
	if !(pagination.SortField == "uuid" || pagination.SortField == "title" || pagination.SortField == "description") {
		return nil, 0, http.StatusBadRequest, model.ErrRiskSortFieldNotAllowed
	}
	if risks, total, err := s.store.ListRisks(ctx, pagination.SortField, pagination.SortOrder, pagination.Limit, pagination.Offset); err != nil {
		return nil, 0, http.StatusInternalServerError, err
	} else {
		return risks, total, http.StatusOK, nil
	}
}

func (s *ServiceHandler) GetRisk(ctx context.Context, riskId uuid.UUID) (*model.Risk, int, error) {
	if risk, err := s.store.GetRisks(ctx, riskId); err != nil {
		if errors.Is(err, model.ErrRiskNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	} else {
		return risk, http.StatusOK, nil
	}
}

func (s *ServiceHandler) CreateRisk(ctx context.Context, risk *model.Risk) (int, error) {
	if err := s.store.CreateRisks(ctx, risk); err != nil {
		if errors.Is(err, model.ErrRiskAlreadyPresent) {
			return http.StatusConflict, err
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
