package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NodeList(c *gin.Context) {
	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Nodes().List(context.TODO(), meta_v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func NodeInfo(c *gin.Context) {
	nodes := c.Param("name")
	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Nodes().Get(context.TODO(), nodes, meta_v1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}
