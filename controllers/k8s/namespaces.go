package k8s

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

func NamespacesList(c *gin.Context) {

	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func NamespacesInfo(c *gin.Context) {
	name := c.Param("namespaces")
	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func NamespacesCreate(c *gin.Context) {

	type Namespace struct {
		Name string `json:"name"`
	}

	type ObjectNamespace struct {
		Objects Namespace
		Types   string `json:"types"`
	}

	type StringNamespace struct {
		Objects string
		Types   string `json:"types"`
	}

	clientset := utils.KubernetesAuth()

	dataType, _ := c.GetRawData()

	var NewObjectNamespace ObjectNamespace

	json.Unmarshal(dataType, &NewObjectNamespace)

	if NewObjectNamespace.Types == "edit" {

		var NewStringNamespace StringNamespace

		json.Unmarshal(dataType, &NewStringNamespace)

		decodes := scheme.Codecs.UniversalDeserializer().Decode

		obj, _, _ := decodes([]byte(NewStringNamespace.Objects), nil, nil)

		namespace := obj.(*corev1.Namespace)

		_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":      "error",
				"status_code": "500",
				"error":       err,
			})
			c.Abort()
		}

	} else {

		namespace := corev1.Namespace{}

		namespace.Name = NewObjectNamespace.Objects.Name

		_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":      "error",
				"status_code": "500",
				"error":       err,
			})
			c.Abort()
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "seccuss",
		"status_code": "200",
	})
}

func NamespacesDelete(c *gin.Context) {

	name := c.Param("namespaces")

	clientset := utils.KubernetesAuth()
	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}
	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), name, deleteOptions)

	c.JSON(200, gin.H{
		"status": 200,
		"error":  err,
	})
}
