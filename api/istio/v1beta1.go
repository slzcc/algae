package istio

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/controllers/istio"
)

func IstioInfo(router *gin.RouterGroup) {
	router.GET("/namespaces/:namespaces/virtualservices", istio.VirtualServicesList)
}
