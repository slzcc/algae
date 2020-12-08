package k8s

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func ConfigMapList(c *gin.Context) {

	namespace := c.Param("namespaces")

	clientset := utils.KubernetesAuth()

	data, err := clientset.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func ConfigMapInfo(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")

	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func ConfigMapModify(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")
	object := c.PostForm("object")

	clientset := utils.KubernetesAuth()
	result, getErr := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Containers: %s", getErr))
	}

	configmap := &corev1.ConfigMap{}

	yaml.Unmarshal([]byte(object), configmap)

	result.Data = configmap.Data

	_, err := clientset.CoreV1().ConfigMaps(namespace).Update(context.TODO(), result, metav1.UpdateOptions{})

	c.JSON(200, gin.H{
		"status": result,
		"error":  err,
	})
}

func ConfigMapCreate(c *gin.Context) {

	namespace := c.Param("namespaces")
	object := c.PostForm("object")

	clientset := utils.KubernetesAuth()
	decode := scheme.Codecs.UniversalDeserializer().Decode

	obj, _, _ := decode([]byte(object), nil, nil)

	configmap := obj.(*corev1.ConfigMap)

	configmap.Namespace = namespace

	result, err := clientset.CoreV1().ConfigMaps(namespace).Create(context.TODO(), configmap, metav1.CreateOptions{})

	c.JSON(200, gin.H{
		"status": result,
		"error":  err,
	})
}

func ConfigMapDelete(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")

	clientset := utils.KubernetesAuth()
	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}
	err := clientset.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, deleteOptions)

	c.JSON(200, gin.H{
		"status": 200,
		"error":  err,
	})
}
