package istio

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ServiceEntrysList(c *gin.Context) {

	namespace := c.Param("namespaces")

	clientset := utils.IstioAuth()
	data, err := clientset.NetworkingV1beta1().ServiceEntries(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}
