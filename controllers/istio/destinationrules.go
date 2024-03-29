package istio

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DestinationRulesList(c *gin.Context) {

	namespace := c.Param("namespaces")

	clientset := utils.IstioAuth()
	data, err := clientset.NetworkingV1beta1().DestinationRules(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func DestinationRulesInfo(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")

	clientset := utils.IstioAuth()
	data, err := clientset.NetworkingV1alpha3().DestinationRules(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}