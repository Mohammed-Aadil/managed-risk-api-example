package api

import (
	"net/http"
	"sync/atomic"

	"github.com/Mohammed-Aadil/common-core/pkg/kubestatus"
	"github.com/Mohammed-Aadil/common-core/pkg/response"
	"github.com/gin-gonic/gin"
)

// Healthz doc
// @Summary 		Liveness check
// @Description 	Used by kubernetes liveness prob
// @Tags 			Kubernetes
// @Accept			json
// @Produce			json
// @Success			200 {string} string "OK"
// @Failure			503
// @Router			/healthz [get]
func (h *APIHandler) HealthzAPI(c *gin.Context) {
	if atomic.LoadInt32(&kubestatus.Healthy) == 1 {
		response.JsonResponse(c, map[string]string{"status": "OK"}, http.StatusOK)
		return
	}
	response.JsonResponse(c, nil, http.StatusServiceUnavailable)
}

// Readyz doc
// @Summary 		Readiness check
// @Description 	Used by kubernetes liveness prob
// @Tags 			Kubernetes
// @Accept			json
// @Produce			json
// @Success			200 {string} string "OK"
// @Failure			503
// @Router			/readyz [get]
func (h *APIHandler) ReadyzAPI(c *gin.Context) {
	if atomic.LoadInt32(&kubestatus.Ready) == 1 {
		response.JsonResponse(c, map[string]string{"status": "OK"}, http.StatusOK)
		return
	}
	response.JsonResponse(c, nil, http.StatusServiceUnavailable)
}
