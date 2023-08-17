package server

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func InitK8SClient() (*kubernetes.Clientset, *rest.Config, error) {
	logrus.Infof("begin init k8s client")
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
	if err != nil {
		logrus.Errorf("get client config failure: %v", err)

	}
	// 创建 client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Errorf("create clientset failure: %v", err)
	}

	return clientSet, config, nil
}
