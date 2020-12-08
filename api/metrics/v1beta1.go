package metrics

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/controllers"
)

func MetricsPods(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/pods", controllers.MetricsPods)
}

func MetricsNodes(router *gin.RouterGroup) {
	router.GET("/nodes", controllers.MetricsNodes)
}
