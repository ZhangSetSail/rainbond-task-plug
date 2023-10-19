package pkg

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	restclient "k8s.io/client-go/rest"
)

var (
	config    *rest.Config
	clientSet *kubernetes.Clientset
)

func InitK8SClient() error {
	logrus.Infof("begin init k8s client")
	var err error
	config, err = restclient.InClusterConfig()
	if err != nil {
		return err
	}
	clientSet, err = kubernetes.NewForConfig(config)
	return err
}

func GetConfig() *rest.Config {
	return config
}

func GetClientSet() *kubernetes.Clientset {
	return clientSet
}
