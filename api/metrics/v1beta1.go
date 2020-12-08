package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/controllers/metrics"
)

func MetricsPods(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/pods", metrics.MetricsPods)
}

func MetricsNodes(router *gin.RouterGroup) {
	router.GET("/nodes", metrics.MetricsNodes)
}
