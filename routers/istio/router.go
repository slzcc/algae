package istio

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/api/istio"
)

func LRouter(router *gin.Engine) {

	istio_api := router.Group("/api/networking.istio.io/v1beta1")
	{
		istio.VirtualServicesList(istio_api)
		istio.VirtualServicesInfo(istio_api)
		istio.GatewaysList(istio_api)
		istio.DestinationRulesList(istio_api)
		istio.DestinationRulesInfo(istio_api)
		istio.ServiceEntrysList(istio_api)
	}

}
