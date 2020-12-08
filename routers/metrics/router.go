package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/api/metrics"
)

func LRouter(router *gin.Engine) {

	metrics_api := router.Group("/api/metrics.k8s.io/v1beta1")
	{
		metrics.MetricsNodes(metrics_api)
		metrics.MetricsPods(metrics_api)
	}
}
