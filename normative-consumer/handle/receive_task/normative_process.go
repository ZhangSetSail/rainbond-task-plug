package receive_task

import (
	"bytes"
	"context"
	"fmt"
	db_model "github.com/goodrain/rainbond-task-plug/db/model"
	"github.com/goodrain/rainbond-task-plug/db/mysql"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/goodrain/rainbond-task-plug/pkg"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/remotecommand"

	"gorm.io/gorm"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"strconv"
	"strings"
	"time"
)

type ProcessNormative struct {
	ctx        context.Context
	DB         *gorm.DB
	kubeClient *kubernetes.Clientset
	config     *rest.Config
}

func (s ProcessNormative) Check(ni model.NormativeInspectionModel) {
	selector := fmt.Sprintf("service_id=%v", ni.ComponentID)
	podL, err := s.kubeClient.CoreV1().Pods("").List(s.ctx, v1.ListOptions{LabelSelector: selector})
	if err != nil {
		logrus.Errorf("by service id select pod failure: %v", err)
		return
	}
	var pod corev1.Pod
	if podL == nil || len(podL.Items) == 0 {
		pod = podL.Items[0]
	} else {
		logrus.Errorf("pod is not exist")
		return
	}
	containerName := pod.Spec.Containers[0].Name
	req := s.kubeClient.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(pod.Name).
		Namespace(pod.Namespace).
		SubResource("exec")
	req.VersionedParams(&corev1.PodExecOptions{
		Container: containerName,
		Command:   []string{"sh", "-c", "ps -ef |grep -v ps |grep -v bash|grep -v 'nginx: worker process'|wc -l"},
		Stdin:     false,
		Stdout:    true,
		Stderr:    true,
		TTY:       false,
	}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(s.config, "POST", req.URL())
	if err != nil {
		logrus.Errorf("remotecommand newSPDYExecutor failure: %v", err)
		return
	}
	var stdout, stderr bytes.Buffer
	err = executor.Stream(remotecommand.StreamOptions{
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		logrus.Errorf("executor stream failure: %v", err)
		return
	}
	num := 0
	numString := strings.Replace(stdout.String(), "\n", "", -1)
	num, err = strconv.Atoi(numString)
	if err != nil {
		logrus.Errorf("strconv atoi failure: %v", err)
		return
	}
	if num-2 != 1 {
		probeReport := db_model.ComponentReport{
			CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
			Level:       0,
			Message:     "组件主容器进程数不为 1",
			ComponentID: ni.ComponentID,
			PrimaryLink: "",
			Type:        "normative",
		}
		err := s.DB.Debug().Create(probeReport).Error
		if err != nil {
			logrus.Errorf("create service normative probe record failure: %v", err)
		}
	}

}

func NewProcessNormative() *ProcessNormative {
	db := mysql.GetDB()
	client := pkg.GetClientSet()
	config := pkg.GetConfig()
	ctx := pkg.GetCTX()
	return &ProcessNormative{
		ctx:        ctx,
		DB:         db,
		config:     config,
		kubeClient: client,
	}
}
