package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonResponse(c *gin.Context, resp any, code int) {
	if code >= http.StatusBadRequest {
		c.AbortWithStatusJSON(code, resp)
	} else {
		c.JSON(code, resp)
	}
}
