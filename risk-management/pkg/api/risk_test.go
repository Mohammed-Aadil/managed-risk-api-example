package api

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Mohammed-Aadil/common-core/pkg/response"
	"github.com/Mohammed-Aadil/risk-management/internal/model"
)

func TestRiskCreateAPI(t *testing.T) {
	var riskResponse response.HttpPaginatedResponse
	resp, _ := callMockAPI(http.MethodGet, "/api/v1/risks", nil, http.StatusOK)

	if err := json.Unmarshal(resp, &riskResponse); err != nil {
		t.Fatal("unable to parse response")
	}

	if riskResponse.Pagination.Total > 0 {
		t.Fatal("not expecting any risks")
	}

	risk1 := model.Risk{
		Title:       "CVE123",
		Description: "os vulnerabilities",
	}
	risk2 := model.Risk{
		Title:       "CVE566",
		Description: "Appliance vulnerabilities",
	}

	callMockAPI(http.MethodPost, "/api/v1/risks", risk1, http.StatusOK)
	callMockAPI(http.MethodPost, "/api/v1/risks", risk2, http.StatusOK)

	resp, _ = callMockAPI(http.MethodGet, "/api/v1/risks", nil, http.StatusOK)

	if err := json.Unmarshal(resp, &riskResponse); err != nil {
		t.Fatal("unable to parse response")
	}

	risks := riskResponse.Data.([]any)
	if len(risks) != 2 {
		t.Fatal("incorrect number of data received")
	}
	callMockAPI(http.MethodPost, "/api/v1/risks", risk2, http.StatusConflict)

}

func TestGetRiskAPI(t *testing.T) {
	var riskResponse response.HttpPaginatedResponse

	risk1 := model.Risk{
		Title:       "CVE1235",
		Description: "os vulnerabilities",
	}
	risk2 := model.Risk{
		Title:       "CVE5665",
		Description: "Appliance vulnerabilities",
	}

	callMockAPI(http.MethodPost, "/api/v1/risks", risk1, http.StatusOK)
	callMockAPI(http.MethodPost, "/api/v1/risks", risk2, http.StatusOK)

	resp, _ := callMockAPI(http.MethodGet, "/api/v1/risks?sortField=title&sortOrder=desc", nil, http.StatusOK)

	if err := json.Unmarshal(resp, &riskResponse); err != nil {
		t.Fatal("unable to parse response")
	}

	for _, riskAny := range riskResponse.Data.([]any) {
		var riskData model.Risk
		risk := riskAny.(map[string]any)
		id := risk["uuid"].(string)
		resp, _ := callMockAPI(http.MethodGet, "/api/v1/risks/"+id, nil, http.StatusOK)

		if err := json.Unmarshal(resp, &riskData); err != nil {
			t.Fatal("unable to parse response")
		}

		if riskData.Title != risk["title"].(string) {
			t.Fatal("incorrect title for the risk id")
		}
	}
}
