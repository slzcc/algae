package istio

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func VirtualServicesList(c *gin.Context) {

	namespace := c.Param("namespaces")

	clientset := utils.IstioAuth()
	data, err := clientset.NetworkingV1beta1().VirtualServices(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func VirtualServicesInfo(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")

	clientset := utils.IstioAuth()
	data, err := clientset.NetworkingV1alpha3().VirtualServices(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}