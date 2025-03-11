package model

import (
	"github.com/Mohammed-Aadil/risk-management/internal/enum"
	"github.com/google/uuid"
)

type Risk struct {
	ID          uuid.UUID      `json:"uuid"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	State       enum.RiskState `json:"state"`
}

const RiskId = "id"