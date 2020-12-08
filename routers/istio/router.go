package istio

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/api/istio"
)

func LRouter(router *gin.Engine) {

	istio_api := router.Group("/api/networking.istio.io/v1beta1")
	{
		istio.VirtualServicesListInfo(istio_api)
		istio.GatewaysListInfo(istio_api)
		istio.DestinationRulesListInfo(istio_api)
		istio.ServiceEntrysListInfo(istio_api)
	}

}
