package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/kubernetes/scheme"
)

func EventList(c *gin.Context) {

	namespace := c.Param("namespaces")

	clientset := utils.KubernetesAuth()

	data, err := clientset.CoreV1().Events(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func EventInfo(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")

	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Events(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}
