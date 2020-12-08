package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/api/k8s/v1"
)

func LRouter(router *gin.Engine) {

	namespaces_api := router.Group("/api/v1/namespaces")
	{
		v1.Namespaces(namespaces_api)
		v1.NamespacesInfo(namespaces_api)

		v1.Pods(namespaces_api)
		v1.PodsInfo(namespaces_api)
		v1.PodsInfox(namespaces_api)
		v1.PodsWatch(namespaces_api)
		v1.PodsLogs(namespaces_api)
		v1.PodsExec(namespaces_api)

		v1.DeploymentList(namespaces_api)
		v1.DeploymentInfo(namespaces_api)
		v1.DeploymentGetTemplate(namespaces_api)

		v1.ServiceList(namespaces_api)
		v1.ServiceInfo(namespaces_api)

		v1.ConfigMapList(namespaces_api)
		v1.ConfigMapInfo(namespaces_api)

		v1.LimitRangeList(namespaces_api)
		v1.LimitRangeInfo(namespaces_api)

		v1.ResourceQuotaList(namespaces_api)
		v1.ResourceQuotaInfo(namespaces_api)

		v1.ReplicaSetList(namespaces_api)
		v1.ReplicaSetInfo(namespaces_api)

		v1.EventList(namespaces_api)
		v1.EventInfo(namespaces_api)
	}
	
	nodes_api := router.Group("/api/v1/nodes")
	{
		v1.NodesList(nodes_api)
		v1.NodesInfo(nodes_api)
	}

}
