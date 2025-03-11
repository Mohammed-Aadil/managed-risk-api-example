package persistence

import (
	"context"
	"sort"

	"github.com/Mohammed-Aadil/common-core/pkg/pagination"
	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/google/uuid"
)

func (store *InMemoryStore) getStoredRisks() []*model.Risk {
	storeRisks := make([]*model.Risk, 0)
	head := store.risks.Front()

	for head != nil {
		storeRisks = append(storeRisks, head.Value.(*model.Risk))
		head = head.Next()
	}
	return storeRisks
}

func (store *InMemoryStore) ListRisks(ctx context.Context, sortField string, sortOrder pagination.SortOrder, limit, offset int) ([]*model.Risk, int, error) {
	storedRisks := store.getStoredRisks()

	sort.Slice(storedRisks, func(i, j int) bool {
		var field1, field2 string
		if sortField == "title" {
			field1, field2 = storedRisks[i].Title, storedRisks[j].Title
		} else if sortField == "uuid" {
			field1, field2 = storedRisks[i].ID.String(), storedRisks[j].ID.String()
		} else if sortField == "state" {
			field1, field2 = storedRisks[i].State.String(), storedRisks[j].State.String()
		}

		if sortOrder == pagination.ASC {
			return field1 < field2
		} else {
			return field1 > field2
		}
	})

	upperLimit := offset + limit
	if upperLimit > len(storedRisks) {
		upperLimit = len(storedRisks)
	}
	risks := storedRisks[offset:upperLimit]
	return risks, len(storedRisks), nil
}

func (store *InMemoryStore) GetRisks(ctx context.Context, riskId uuid.UUID) (*model.Risk, error) {
	if node, ok := store.riskIdMap[riskId]; !ok {
		return nil, model.ErrRiskNotFound
	} else {
		risk := node.Value.(*model.Risk)
		return risk, nil
	}
}

func (store *InMemoryStore) CreateRisks(ctx context.Context, risk *model.Risk) error {
	if _, ok := store.riskTitleMap[risk.Title]; ok {
		return model.ErrRiskAlreadyPresent
	}

	risk.ID = uuid.New()
	store.risks.PushBack(risk)
	store.riskIdMap[risk.ID] = store.risks.Back()
	store.riskTitleMap[risk.Title] = store.risks.Back()
	return nil
}
