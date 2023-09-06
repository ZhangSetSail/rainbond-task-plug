package util

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	restclient "k8s.io/client-go/rest"
)

func InitK8SClient() (*kubernetes.Clientset, *rest.Config, error) {
	logrus.Infof("begin init k8s client")
	config, err := restclient.InClusterConfig()
	if err != nil {
		return nil, nil, err
	}
	// 创建 client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Errorf("create clientset failure: %v", err)
	}

	return clientSet, config, nil
}
