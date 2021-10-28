package devops

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/api/devops/v1"
)

func LRouter(router *gin.Engine) {

	devops_api := router.Group("/api/v1/devops")
	{
		v1.VersionControl(devops_api)
	}
}
