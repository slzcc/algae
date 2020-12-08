package v1

import (
	"github.com/slzcc/algae/controllers"
	"github.com/gin-gonic/gin"
)

// API Nodes Routers
func NodesList(router *gin.RouterGroup){
	router.GET("", controllers.NodeList)
}

func NodesInfo(router *gin.RouterGroup){
	router.GET("/:name", controllers.NodeInfo)
}