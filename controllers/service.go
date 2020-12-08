package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ServiceList(c *gin.Context) {
	namespace := c.Param("namespaces")
	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Services(namespace).List(context.TODO(), meta_v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func ServiceInfo(c *gin.Context) {
	namespace := c.Param("namespaces")
	name := c.Param("name")
	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Services(namespace).Get(context.TODO(), name, meta_v1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func ServicePods(c *gin.Context) {

	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Pods("").List(context.TODO(), meta_v1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	c.JSON(200, data)
}
