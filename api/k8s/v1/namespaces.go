package v1

import (
	"github.com/slzcc/algae/controllers/k8s"
	"github.com/gin-gonic/gin"
)

// https://books.studygolang.com/advanced-go-programming-book/ch6-web/ch6-02-router.html
// API Namespace Routers
func Namespaces(router *gin.RouterGroup){
	router.GET("", k8s.NamespacesList)
	router.POST("", k8s.NamespacesCreate)
}

func NamespacesInfo(router *gin.RouterGroup){
	router.GET("/:namespaces", k8s.NamespacesInfo)
	router.DELETE("/:namespaces", k8s.NamespacesDelete)
}

// API Pods Routers
func Pods(router *gin.RouterGroup){
	router.GET("/:namespaces/pods", k8s.PodsList)
	router.POST("/:namespaces/pods", k8s.PodsGetSelector)
}

func PodsInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name", k8s.PodsInfo)
}

func PodsInfox(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name/info", k8s.PodsInfo)
}

func PodsWatch(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name/watch", k8s.PodsWatch)
}

func PodsLogs(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name/logs/:container", k8s.PodsLogs)
}

func PodsExec(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name/exec/:container", k8s.PodsExec)
}

func ServiceList(router *gin.RouterGroup){
	router.GET("/:namespaces/services", k8s.ServiceList)
}

func ServiceInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/services/:name", k8s.ServiceInfo)
}

func ConfigMapList(router *gin.RouterGroup){
	router.GET("/:namespaces/configmaps", k8s.ConfigMapList)
	router.POST("/:namespaces/configmaps", k8s.ConfigMapCreate)
}

func ConfigMapInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/configmaps/:name", k8s.ConfigMapInfo)
	router.PUT("/:namespaces/configmaps/:name", k8s.ConfigMapModify)
	router.DELETE("/:namespaces/configmaps/:name", k8s.ConfigMapDelete)
}

func LimitRangeList(router *gin.RouterGroup){
	router.GET("/:namespaces/limitranges", k8s.LimitRangeList)
	router.POST("/:namespaces/limitranges", k8s.LimitRangeCreate)
}

func LimitRangeInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/limitranges/:name", k8s.LimitRangeInfo)
	router.PUT("/:namespaces/limitranges/:name", k8s.LimitRangeModify)
	router.DELETE("/:namespaces/limitranges/:name", k8s.LimitRangeDelete)
}

func ResourceQuotaList(router *gin.RouterGroup){
	router.GET("/:namespaces/resourcequotas", k8s.ResourceQuotaList)
	router.POST("/:namespaces/resourcequotas", k8s.ResourceQuotaCreate)
}

func ResourceQuotaInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/resourcequotas/:name", k8s.ResourceQuotaInfo)
	router.PUT("/:namespaces/resourcequotas/:name", k8s.ResourceQuotaModify)
	router.DELETE("/:namespaces/resourcequotas/:name", k8s.ResourceQuotaDelete)
}

func DeploymentList(router *gin.RouterGroup){
	router.GET("/:namespaces/deployments", k8s.DeploymentList)
	router.POST("/:namespaces/deployments", k8s.DeploymentCreate)
}

func DeploymentInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/deployments/:name", k8s.DeploymentInfo)
	router.PUT("/:namespaces/deployments/:name", k8s.DeploymentModify)
	router.DELETE("/:namespaces/deployments/:name", k8s.DeploymentDelete)
}

func DeploymentGetTemplate(router *gin.RouterGroup){
	router.GET("/:namespaces/deployment/templates", k8s.DeploymentGetTemplate)
}

func EventList(router *gin.RouterGroup){
	router.GET("/:namespaces/events", k8s.EventList)
}

func EventInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/events/:name", k8s.EventInfo)
}

func ReplicaSetList(router *gin.RouterGroup){
	router.GET("/:namespaces/replicasets", k8s.ReplicaSetList)
	router.POST("/:namespaces/replicasets", k8s.ReplicaSetGetSelector)
}

func ReplicaSetInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/replicasets/:name", k8s.ReplicaSetInfo)
}