package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/Mohammed-Aadil/risk-management/pkg/persistence"
	"go.uber.org/zap"
)

var (
	mockServer *APIHandler
	logger     *zap.Logger
)

func startMockServer(conf *model.Config, store persistence.Persistence, logger *zap.Logger) {
	api := NewAPI(conf, store, logger)
	mockServer = api.StartTestServer()
}

func TestMain(m *testing.M) {
	conf := &model.Config{
		ServiceName:            "test-risk-management",
		HostName:               "testhost",
		Level:                  "debug",
		HttpPort:               8080,
		HttpReadTimeout:        2 * time.Second,
		HttpWriteTimetout:      5 * time.Second,
		ServerShutdownTimeout:  2 * time.Second,
		StorageBackend:         "inmemory",
		DefaultPaginationLimit: 100,
	}

	logger, _ = zap.NewDevelopment()

	store := persistence.Init(conf, logger)
	startMockServer(conf, store, logger)
	os.Exit(m.Run())
}

func callMockAPI(method, url string, payload any, expectedHttpCode int) ([]byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		logger.Fatal("unable to parse payload")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		logger.Fatal("unable to create http request")
	}

	req.Header.Set("Content-Type", "application/json")
	responseRec := httptest.NewRecorder()

	req.RemoteAddr = fmt.Sprintf("10.0.0.1:%d", mockServer.conf.HttpPort)
	mockServer.router.ServeHTTP(responseRec, req)

	if responseRec != nil && responseRec.Code != expectedHttpCode {
		logger.Fatal("http code mismatch", zap.Int("expectedHttpCode", expectedHttpCode), zap.Int("actualHttpCode", responseRec.Code))
	}

	data, err := io.ReadAll(responseRec.Result().Body)
	if err != nil {
		logger.Error("unable to read response", zap.Error(err))
		return nil, err
	}
	return data, nil
}
