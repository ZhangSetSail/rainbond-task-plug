package server

import (
	"github.com/goodrain/rainbond-task-plug/cmd/normative-consumer/config"
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/normative-consumer/handle"
	"github.com/goodrain/rainbond-task-plug/normative-consumer/router"
	"github.com/goodrain/rainbond-task-plug/pkg"
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
	err = startBackgroundProcess()
	if err != nil {
		logrus.Errorf("start background process failure: %v", err)
		return err
	}
	logrus.Infof("task plug producer server exit")
	return nil
}

func initCli() error {
	//获取配置
	s := config.GetNormativeConsumerServer()
	//初始化 k8s cli
	err := pkg.InitK8SClient()
	if err != nil {
		return err
	}
	// 初始化消息队列
	natsAddr := s.NatsHost + ":" + s.NatsPort
	err = pkg.InitNatsCli(natsAddr)
	if err != nil {
		return err
	}
	router.InitRouterCli()
	//初始化数据库
	err = mysql.InitDB(s.DB)
	if err != nil {
		return err
	}
	//初始化 ctx
	pkg.InitCTX(3)
	//初始化处理程序，必须放在最后
	handle.InitHandle()
	return nil
}

func startBackgroundProcess() error {
	s := config.GetNormativeConsumerServer()
	//启动 server
	r := router.GetRouter()
	exitCh := make(chan int)
	err := handle.GetManagerReceiveTasks().DigestionNormativeInspectionTask()
	if err != nil {
		return err
	}
	server := pkg.InitHttpServer(r, s.Port)
	go pkg.GracefulShutdown(server, exitCh)
	<-exitCh
	return nil
}
