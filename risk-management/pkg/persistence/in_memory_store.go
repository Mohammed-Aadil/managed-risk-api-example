package persistence

import (
	"container/list"

	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type InMemoryStore struct {
	conf         *model.Config
	logger       *zap.Logger
	risks        *list.List
	riskIdMap    map[uuid.UUID]*list.Element
	riskTitleMap map[string]*list.Element
}

func NewInMemoryStore(conf *model.Config, logger *zap.Logger) *InMemoryStore {
	return &InMemoryStore{conf: conf, risks: list.New(), riskIdMap: make(map[uuid.UUID]*list.Element), riskTitleMap: make(map[string]*list.Element), logger: logger}
}
