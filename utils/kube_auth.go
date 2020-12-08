package utils

import (
	"path/filepath"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func KubernetesAuth() (kubernetes.Clientset){

	// creates kubeconfig object
	kubeconfig := filepath.Join("utils", "kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return *clientset

}

func KubeConfig() (*rest.Config){

	// creates kubeconfig object
	kubeconfig := filepath.Join("utils", "kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err.Error())
	}

	return config

}