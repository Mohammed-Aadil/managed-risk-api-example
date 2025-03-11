package api

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func (a *APIHandler) registerCommonMiddleware() {
	a.router.Use(gin.Recovery())
	a.router.Use(ginzap.RecoveryWithZap(a.logger, true))
	a.router.Use(ginzap.Ginzap(a.logger, time.RFC3339, true))
	a.router.Use(ginzap.GinzapWithConfig(a.logger, &ginzap.Config{SkipPaths: []string{"/readyz", "healthz"}}))
}

func (a *APIHandler) registerCommonRoutes() {
	a.router.GET("/", a.InfoAPI)
	a.router.GET("/readyz", a.ReadyzAPI)
	a.router.GET("/healtz", a.HealthzAPI)
}

func (a *APIHandler) registerAPIRoutes() {
	api := a.router.Group("/api/v1")

	api.GET("/risks", a.listRisksAPI)
	api.GET("/risks/:id", a.getRisksAPI)
	api.POST("/risks", a.createRisksAPI)
}
