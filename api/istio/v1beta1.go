package istio

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/controllers/istio"
)

func VirtualServicesListInfo(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/virtualservices", istio.VirtualServicesList)
}

func GatewaysListInfo(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/gateways", istio.GatewaysList)
}

func DestinationRulesListInfo(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/destinationrules", istio.DestinationRulesList)
}

func ServiceEntrysListInfo(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/serviceentrys", istio.ServiceEntrysList)
}
