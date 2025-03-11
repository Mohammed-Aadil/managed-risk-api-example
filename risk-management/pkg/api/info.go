package api

import (
	"net/http"
	"runtime"

	"github.com/Mohammed-Aadil/common-core/pkg/response"
	"github.com/gin-gonic/gin"
)

// Info doc
// @Summary 		Runtime information
// @Description 	returns the runtime information
// @Tags 			HTTP API
// @Accept			json
// @Produce			json
// @Success			200 {object} response.RuntimeResponse
// @Router			/api/info [get]
func (h *APIHandler) InfoAPI(c *gin.Context) {
	data := response.RuntimeResponse{
		Hostname:     h.conf.HostName,
		GOOS:         runtime.GOOS,
		GOARCH:       runtime.GOARCH,
		Runtime:      runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
		NumCPU:       runtime.NumCPU(),
	}
	response.JsonResponse(c, data, http.StatusOK)
}
