package server

import (
	"github.com/goodrain/rainbond-task-plug/cmd/task-plug-producer/config"
	"github.com/goodrain/rainbond-task-plug/pkg"
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/controller"
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/handle"
	init_watch "github.com/goodrain/rainbond-task-plug/task-plug-producer/handle/k8s-watch/init-watch"
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/router"
	"github.com/sirupsen/logrus"
)

func Run() error {
	//初始化所有cli
	err := initCli()
	if err != nil {
		logrus.Errorf("init cli failure: %v", err)
		return err
	}
	//初始化后台服务
	initBackgroundProcess()
	logrus.Infof("task plug producer server exit")
	return nil
}

func initCli() error {
	p := config.GetProducerServer()
	err := pkg.InitK8SClient()
	if err != nil {
		return err
	}
	pkg.InitCTX(3)
	// 初始化消息队列
	natsAddr := p.NatsHost + ":" + p.NatsPort
	logrus.Infof("--------------------%v", natsAddr)
	err = pkg.InitNatsCli(natsAddr)
	if err != nil {
		logrus.Errorf("-----------------error: %v", err)
		return err
	}
	//初始化router路由函数
	err = controller.CreateRouterManager()
	if err != nil {
		return err
	}
	//初始化数据库
	//err = mysql.InitDB(p.DB)
	//if err != nil {
	//	return err
	//}
	//初始化router
	router.InitRouterCli()
	//初始化处理程序,必须放在最后
	handle.InitHandle()
	return nil
}

func initBackgroundProcess() error {
	p := config.GetProducerServer()
	//k8s watch 监听
	mw := init_watch.CreateResourceWatch()
	mw.Start()
	//启动 server
	r := router.GetRouter()
	exitCh := make(chan int)
	s := pkg.InitHttpServer(r, p.Port)
	go pkg.GracefulShutdown(s, exitCh)
	<-exitCh
	return nil
}
