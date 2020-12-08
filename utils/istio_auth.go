package utils

import (
	"path/filepath"

	"istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/tools/clientcmd"
)

func IstioAuth() versioned.Clientset {

	// creates kubeconfig object
	kubeconfig := filepath.Join("utils", "kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return *clientset

}
