package v1

import (
	"github.com/slzcc/algae/controllers/k8s"
	"github.com/gin-gonic/gin"
)

// API Nodes Routers
func NodesList(router *gin.RouterGroup){
	router.GET("", k8s.NodeList)
}

func NodesInfo(router *gin.RouterGroup){
	router.GET("/:name", k8s.NodeInfo)
}