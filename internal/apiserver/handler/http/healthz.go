package http

import (
	apiv1 "miniblog/pkg/api/apiserver/v1"
	"time"

	"github.com/gin-gonic/gin"
)

// Healthz 服务健康检查.
func (h *Handler) Healthz(c *gin.Context) {
	c.JSON(200,&apiv1.HealthzResponse{
		Status: apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	})
}
