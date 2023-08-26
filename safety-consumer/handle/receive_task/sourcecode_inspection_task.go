package receive_task

import (
	"encoding/json"
	"fmt"
	"github.com/goodrain/rainbond-safety/model"
	"github.com/goodrain/rainbond-safety/safety-consumer/handle/clone"
	"github.com/goodrain/rainbond-safety/util"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
)

func (t *ManagerReceiveTask) DigestionSourceCodeInspectionTask() error {
	logrus.Infof("begion receive source code inspection task")
	_, err := t.nc.QueueSubscribe(t.config.Subscribe, t.config.SubscribeQueue, func(m *nats.Msg) {
		defer clearDir()
		var cdm model.CodeDetectionModel
		err := json.Unmarshal(m.Data, &cdm)
		if err != nil {
			logrus.Errorf("json unmarshal failure: %v", err)
			return
		}
		_, _, err = clone.GitClone(cdm, t.config.CodeStoragePath, 10, t.ctx)
		if err != nil {
			logrus.Errorf("git clone failure: %v", err)
			return
		}
		srcPath := fmt.Sprintf("SRC_PATH=%v", "/usr/src")
		sonarToken := fmt.Sprintf("SONAR_TOKEN=%v", t.config.SonarToken)
		sonarScannerOpts := fmt.Sprintf("SONAR_SCANNER_OPTS=-Dsonar.projectKey=%v -Dsonar.exclusions=**/*.java", cdm.ProjectName)
		SonarHostURL := fmt.Sprintf("SONAR_HOST_URL=%v", t.config.SonarHostUrl)
		command := "/usr/bin/entrypoint.sh"
		args := []string{"sonar-scanner"}
		envs := []string{sonarToken, sonarScannerOpts, SonarHostURL, srcPath}
		err = util.ExecCommand(command, args, envs)
		if err != nil {
			logrus.Errorf("code inspection execution failure: %v", err)
			return
		}
	})
	return err
}

func clearDir() {
	logrus.Infof("begin remove /usr/src/*")
	dir, _ := ioutil.ReadDir("/usr/src")
	for _, d := range dir {
		p := path.Join([]string{"/usr/src", d.Name()}...)
		err := os.RemoveAll(p)
		if err != nil {
			logrus.Errorf("remove %v failure: %v", p, err)
		}
	}
}
