package k8s

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/slzcc/algae/utils"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func PodsList(c *gin.Context) {

	namespace := c.Param("namespaces")
	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	c.JSON(200, data)
}

func PodsInfo(c *gin.Context) {
	namespace := c.Param("namespaces")
	name := c.Param("name")
	clientset := utils.KubernetesAuth()
	data, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func PodsGetSelector(c *gin.Context) {
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

	data, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), options)

	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func PodsWatch(c *gin.Context) {
	namespace := c.Param("namespaces")
	clientset := utils.KubernetesAuth()

	data, err := clientset.CoreV1().Pods(namespace).Watch(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, data)
}

func PodsLogs(c *gin.Context) {
	namespace := c.Param("namespaces")
	name := c.Param("name")
	container := c.Param("container")
	clientset := utils.KubernetesAuth()

	podLogOpts := apiv1.PodLogOptions{
		Container: container,
	}

	data := clientset.CoreV1().Pods(namespace).GetLogs(name, &podLogOpts)

	podLogs, err := data.Stream(context.TODO())
	if err != nil {
		panic(err.Error())
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		panic(err.Error())
	}
	str := buf.String()

	c.JSON(200, str)
}

func PodsExec(c *gin.Context) {
	namespace := c.Param("namespaces")
	name := c.Param("name")
	container := c.Param("container")
	command := c.Query("command")

	if command == "" {
		c.JSON(200, gin.H{"Stdout": "", "Stderr": "Invalid instruction"})
		c.Abort()
	}
	clientset := utils.KubernetesAuth()

	config := utils.KubeConfig()

	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(name).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&apiv1.PodExecOptions{
		Container: container,
		Command:   []string{command},
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
	}, scheme.ParameterCodec)

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		panic(err.Error())
	}

	buf := &bytes.Buffer{}
	errBuf := &bytes.Buffer{}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: buf,
		Stderr: errBuf,
		Tty:    true,
	})
	if err != nil {
		panic(err.Error())
	}
	c.JSON(200, gin.H{"Stdout": buf.String(), "Stderr": errBuf.String()})
}
