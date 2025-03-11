package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Mohammed-Aadil/common-core/pkg/kubestatus"
	"github.com/Mohammed-Aadil/risk-management/internal/model"
	"github.com/Mohammed-Aadil/risk-management/internal/service"
	"github.com/Mohammed-Aadil/risk-management/pkg/persistence"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type API interface {
	StartServer(stopCh chan struct{})
	StartTestServer() *APIHandler
}

type APIHandler struct {
	conf    *model.Config
	service service.Service
	router  *gin.Engine
	logger  *zap.Logger
}

func NewAPI(conf *model.Config, store persistence.Persistence, logger *zap.Logger) API {
	return &APIHandler{conf: conf, service: service.NewService(conf, store, logger), logger: logger}
}

func (a *APIHandler) StartServer(stopCh chan struct{}) {
	a.router = gin.Default()
	a.registerCommonMiddleware()
	a.registerCommonRoutes()
	a.registerAPIRoutes()

	server := a.ListenAndServe()

	kubestatus.SetHealthy()
	kubestatus.SetReady()
	<-stopCh

	ctx, cancel := context.WithTimeout(context.Background(), a.conf.ServerShutdownTimeout)
	defer cancel()

	kubestatus.SetUnHealthy()
	kubestatus.SetUnReady()

	a.logger.Info("Shutting down the server listening on", zap.Int("port", a.conf.HttpPort), zap.String("service", a.conf.ServiceName))

	if err := server.Shutdown(ctx); err != nil {
		a.logger.Warn("unable to shutdown server gracefully", zap.String("service", a.conf.ServiceName), zap.Error(err))
	}
}

func (a *APIHandler) ListenAndServe() *http.Server {
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", a.conf.HttpPort),
		Handler:      a.router,
		ReadTimeout:  a.conf.HttpReadTimeout,
		WriteTimeout: a.conf.HttpWriteTimetout,
	}

	go func() {
		a.logger.Info("Started listening on port", zap.Int("port", a.conf.HttpPort), zap.String("service", a.conf.ServiceName))
		httpServer.ListenAndServe()
	}()

	return httpServer
}

func (a *APIHandler) StartTestServer() *APIHandler {
	a.router = gin.Default()

	gin.SetMode(gin.TestMode)
	gin.EnableJsonDecoderDisallowUnknownFields()

	a.registerCommonMiddleware()
	a.registerCommonRoutes()

	a.registerAPIRoutes()

	return a
}
