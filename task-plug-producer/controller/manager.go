package controller

import (
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/api"
	"github.com/sirupsen/logrus"
)

var defaultManager Manager

type Manager interface {
	api.ProducerInterface
	api.NormativeInterface
}

func CreateRouterManager() (err error) {
	logrus.Infof("create router manager")
	defaultManager = NewManager()
	return err
}

func NewManager() *RouterManager {
	return &RouterManager{}
}

func GetManager() Manager {
	return defaultManager
}
