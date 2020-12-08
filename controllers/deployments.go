package controllers

import (
	"context"
	"fmt"

	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	"gopkg.in/yaml.v2"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeploymentList(c *gin.Context) {

	namespace := c.Param("namespaces")

	clientset := utils.KubernetesAuth()

	data, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func DeploymentInfo(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")

	clientset := utils.KubernetesAuth()

	data, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func DeploymentModify(c *gin.Context) {

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

func DeploymentDelete(c *gin.Context) {

	namespace := c.Param("namespaces")
	name := c.Param("name")

	clientset := utils.KubernetesAuth()
	deletePolicy := metav1.DeletePropagationForeground
	deleteOptions := metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}
	err := clientset.AppsV1().Deployments(namespace).Delete(context.TODO(), name, deleteOptions)

	c.JSON(200, gin.H{
		"status": 200,
		"error":  err,
	})
}

//func DeploymentCreate(c *gin.Context)  {
//	var r corev1.ResourceRequirements
//	//资源分配会遇到无法设置值的问题，故采用json反解析
//	j := `{"limits": {"cpu":"2000m", "memory": "1Gi"}, "requests": {"cpu":"2000m", "memory": "1Gi"}}`
//	json.Unmarshal([]byte(j), &r)
//
//	namespace := c.Param("namespaces")
//	//name := c.Param("name")
//	//object := c.PostForm("object")
//	clientset := utils.KubernetesAuth()
//	data := clientset.AppsV1().Deployments(namespace)
//
//	deployment := &appsv1.Deployment {
//		ObjectMeta: metav1.ObjectMeta {
//			Name: "demo-deployment",       // 指定 deployment 名字
//		},
//		Spec: appsv1.DeploymentSpec{
//			Replicas: int32Ptr(2),         // 指定副本数
//			Selector: &metav1.LabelSelector {  // 指定标签
//				MatchLabels: map[string]string {
//					"app": "demo",
//				},
//			},
//			Template: corev1.PodTemplateSpec { // 容器模板
//				ObjectMeta: metav1.ObjectMeta {
//					Labels: map[string]string {
//						"app": "demo",
//					},
//				},
//				Spec: corev1.PodSpec{
//					Containers: []corev1.Container {
//						{
//							Name:            "nginx",
//							Image:           "nginx",
//							Resources: r,
//						},
//					},
//				},
//			},
//		},
//	}
//
//	_, err := data.Create(deployment)
//	if err != nil {
//		panic(err.Error())
//	}
//	c.JSON(200, gin.H{
//		"status":  "200",
//	})
//}
//
//func int32Ptr(i int32) *int32 {
//	return &i
//}

func DeploymentCreate(c *gin.Context) {

	//type ObjectDeployment struct {
	//	Objects *appsv1.Deployment
	//	Types string `json:"types"`
	//}
	//
	//type StringDeployment struct {
	//	Objects string
	//	Types string `json:"types"`
	//}

	namespace := c.Param("namespaces")
	dataType, _ := c.GetRawData()

	clientset := utils.KubernetesAuth()

	var NewObjectDeployment appsv1.Deployment

	json.Unmarshal(dataType, &NewObjectDeployment)

	result, err := clientset.AppsV1().Deployments(namespace).Create(context.TODO(), &NewObjectDeployment, metav1.CreateOptions{})
	fmt.Println(result, err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":      "error",
			"status_code": "500",
			"error":       err,
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      "seccuss",
		"status_code": "200",
	})
}

func DeploymentGetTemplate(c *gin.Context) {

	namespace := c.Param("namespaces")
	var _name string = "\n"

	var _limit corev1.ResourceRequirements
	_limit_template := `{"limits": {"cpu":"500m", "memory": "0.5Gi"}, "requests": {"cpu":"500m", "memory": "0.5Gi"}}`
	json.Unmarshal([]byte(_limit_template), &_limit)

	NewDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      _name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: new(int32),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": _name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": _name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            _name,
							Image:           _name,
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
							Command: []string{""},
							Env: []corev1.EnvVar{
								{
									Name:  "TZ",
									Value: "Asia/Shanghai",
								},
							},
							Resources: _limit,
						},
					},
				},
			},
		},
	}

	json.Marshal(*NewDeployment)

	c.JSON(http.StatusOK, gin.H{
		"status":      "seccuss",
		"status_code": "200",
		"template":    NewDeployment,
	})
}
