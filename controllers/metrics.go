package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MetricsNodes(c *gin.Context) {

	clientset := utils.MetricsAuth()
	data, err := clientset.MetricsV1beta1().NodeMetricses().List(context.TODO(), meta_v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func MetricsPods(c *gin.Context) {

	name := c.Param("namespaces")
	clientset := utils.MetricsAuth()
	data, err := clientset.MetricsV1beta1().PodMetricses(name).List(context.TODO(), meta_v1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}
