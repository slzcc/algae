package istio

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/controllers/istio"
)

func VirtualServicesList(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/virtualservices", istio.VirtualServicesList)
}

func VirtualServicesInfo(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/virtualservices/:name", istio.VirtualServicesInfo)
}

func GatewaysList(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/gateways", istio.GatewaysList)
}

func DestinationRulesList(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/destinationrules", istio.DestinationRulesList)
}

func DestinationRulesInfo(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/destinationrules/:name", istio.DestinationRulesInfo)
}

func ServiceEntrysList(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/serviceentrys", istio.ServiceEntrysList)
}
