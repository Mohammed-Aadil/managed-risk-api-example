package service

import (
	"context"

	"github.com/Mohammed-Aadil/common-core/pkg/pagination"
	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/Mohammed-Aadil/risk-management/pkg/persistence"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	ListRisk(ctx context.Context, pagination *pagination.Pagination) ([]*model.Risk, int, int, error)
	GetRisk(ctx context.Context, riskId uuid.UUID) (*model.Risk, int, error)
	CreateRisk(ctx context.Context, risk *model.Risk) (int, error)
}

type ServiceHandler struct {
	conf   *model.Config
	store  persistence.Persistence
	logger *zap.Logger
}

func NewService(conf *model.Config, store persistence.Persistence, logger *zap.Logger) Service {
	return &ServiceHandler{conf: conf, store: store, logger: logger}
}
