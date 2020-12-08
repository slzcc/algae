package v1

import (
	"github.com/slzcc/algae/controllers"
	"github.com/gin-gonic/gin"
)

// https://books.studygolang.com/advanced-go-programming-book/ch6-web/ch6-02-router.html
// API Namespace Routers
func Namespaces(router *gin.RouterGroup){
	router.GET("", controllers.NamespacesList)
	router.POST("", controllers.NamespacesCreate)
}

func NamespacesInfo(router *gin.RouterGroup){
	router.GET("/:namespaces", controllers.NamespacesInfo)
	router.DELETE("/:namespaces", controllers.NamespacesDelete)
}

//func NamespacesPods(router *gin.RouterGroup){
//	router.GET("/:namespaces/all", controllers.NamespacePods)
//}

// API Pods Routers
func Pods(router *gin.RouterGroup){
	router.GET("/:namespaces/pods", controllers.PodsList)
	router.POST("/:namespaces/pods", controllers.PodsGetSelector)
}

func PodsInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name", controllers.PodsInfo)
}

func PodsInfox(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name/info", controllers.PodsInfo)
}

func PodsWatch(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name/watch", controllers.PodsWatch)
}

func PodsLogs(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name/logs/:container", controllers.PodsLogs)
}

func PodsExec(router *gin.RouterGroup){
	router.GET("/:namespaces/pods/:name/exec/:container", controllers.PodsExec)
}

func ServiceList(router *gin.RouterGroup){
	router.GET("/:namespaces/services", controllers.ServiceList)
}

func ServiceInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/services/:name", controllers.ServiceInfo)
}

func ConfigMapList(router *gin.RouterGroup){
	router.GET("/:namespaces/configmaps", controllers.ConfigMapList)
	router.POST("/:namespaces/configmaps", controllers.ConfigMapCreate)
}

func ConfigMapInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/configmaps/:name", controllers.ConfigMapInfo)
	router.PUT("/:namespaces/configmaps/:name", controllers.ConfigMapModify)
	router.DELETE("/:namespaces/configmaps/:name", controllers.ConfigMapDelete)
}

func LimitRangeList(router *gin.RouterGroup){
	router.GET("/:namespaces/limitranges", controllers.LimitRangeList)
	router.POST("/:namespaces/limitranges", controllers.LimitRangeCreate)
}

func LimitRangeInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/limitranges/:name", controllers.LimitRangeInfo)
	router.PUT("/:namespaces/limitranges/:name", controllers.LimitRangeModify)
	router.DELETE("/:namespaces/limitranges/:name", controllers.LimitRangeDelete)
}

func ResourceQuotaList(router *gin.RouterGroup){
	router.GET("/:namespaces/resourcequotas", controllers.ResourceQuotaList)
	router.POST("/:namespaces/resourcequotas", controllers.ResourceQuotaCreate)
}

func ResourceQuotaInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/resourcequotas/:name", controllers.ResourceQuotaInfo)
	router.PUT("/:namespaces/resourcequotas/:name", controllers.ResourceQuotaModify)
	router.DELETE("/:namespaces/resourcequotas/:name", controllers.ResourceQuotaDelete)
}

func DeploymentList(router *gin.RouterGroup){
	router.GET("/:namespaces/deployments", controllers.DeploymentList)
	router.POST("/:namespaces/deployments", controllers.DeploymentCreate)
}

func DeploymentInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/deployments/:name", controllers.DeploymentInfo)
	router.PUT("/:namespaces/deployments/:name", controllers.DeploymentModify)
	router.DELETE("/:namespaces/deployments/:name", controllers.DeploymentDelete)
}

func DeploymentGetTemplate(router *gin.RouterGroup){
	router.GET("/:namespaces/deployment/templates", controllers.DeploymentGetTemplate)
}

func EventList(router *gin.RouterGroup){
	router.GET("/:namespaces/events", controllers.EventList)
}

func EventInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/events/:name", controllers.EventInfo)
}

func ReplicaSetList(router *gin.RouterGroup){
	router.GET("/:namespaces/replicasets", controllers.ReplicaSetList)
	router.POST("/:namespaces/replicasets", controllers.ReplicaSetGetSelector)
}

func ReplicaSetInfo(router *gin.RouterGroup){
	router.GET("/:namespaces/replicasets/:name", controllers.ReplicaSetInfo)
}