package persistence

import (
	"context"

	"github.com/Mohammed-Aadil/common-core/pkg/pagination"
	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Persistence interface {
	ListRisks(ctx context.Context, sortField string, sortOrder pagination.SortOrder, limit, offset int) ([]*model.Risk, int, error)
	GetRisks(ctx context.Context, riskId uuid.UUID) (*model.Risk, error)
	CreateRisks(ctx context.Context, risk *model.Risk) error
}

func Init(conf *model.Config, logger *zap.Logger) Persistence {
	if conf.StorageBackend == "inmemory" {
		return NewInMemoryStore(conf, logger)
	} else {
		logger.Fatal("unknown backend storage config")
	}
	return nil
}
