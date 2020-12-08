package controllers

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	//"k8s.io/client-go/kubernetes/scheme"
)

func ReplicaSetList(c *gin.Context) {

	namespace := c.Param("namespaces")

	clientset := utils.KubernetesAuth()

	data, err := clientset.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func ReplicaSetInfo(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")

	clientset := utils.KubernetesAuth()
	data, err := clientset.AppsV1().ReplicaSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func ReplicaSetGetSelector(c *gin.Context) {
	type matchLabels struct {
		Objects map[string]string
	}
	namespace := c.Param("namespaces")
	dataType, _ := c.GetRawData()

	var matchLabelsObject matchLabels

	json.Unmarshal(dataType, &matchLabelsObject)

	listOptions := &metav1.LabelSelector{
		MatchLabels: matchLabelsObject.Objects,
	}

	labelMap, _ := metav1.LabelSelectorAsMap(listOptions)

	options := metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(labelMap).String(),
	}

	clientset := utils.KubernetesAuth()

	data, err := clientset.AppsV1().ReplicaSets(namespace).List(context.TODO(), options)

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}
