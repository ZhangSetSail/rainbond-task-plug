package receive_task

import (
	"encoding/json"
	"fmt"
	db_model "github.com/goodrain/rainbond-task-plug/db/model"
	"github.com/goodrain/rainbond-task-plug/model"
	"github.com/goodrain/rainbond-task-plug/pkg"
	gojsonq "github.com/thedevsaddam/gojsonq/v2"

	"github.com/goodrain/rainbond-task-plug/safety-consumer/handle/clone"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

func (t *ManagerReceiveTask) DigestionSourceCodeInspectionTask() error {
	logrus.Infof("begion receive source code inspection task")
	_, err := t.nc.QueueSubscribe(model.SOURCE_CODE_INSPECTION, "", func(m *nats.Msg) {
		defer clearDir()
		var cdm model.CodeInspectionModel
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
		if t.config.SonarToken == "" {
			t.config.SonarToken = os.Getenv("SONAR_TOKEN")
		}
		sonarToken := fmt.Sprintf("SONAR_TOKEN=%v", t.config.SonarToken)
		sonarScannerOpts := fmt.Sprintf("SONAR_SCANNER_OPTS=-Dsonar.projectKey=%v -Dsonar.exclusions=**/*.java", cdm.ProjectName)
		sonarAddr := t.config.SonarHost + ":" + t.config.SonarPort
		SonarHostURL := fmt.Sprintf("SONAR_HOST_URL=http://%v", sonarAddr)
		command := "/usr/bin/entrypoint.sh"
		args := []string{"sonar-scanner"}
		envs := []string{sonarToken, sonarScannerOpts, SonarHostURL, srcPath}
		err = pkg.ExecCommand(command, args, envs)
		if err != nil {
			logrus.Errorf("code inspection execution failure: %v", err)
			return
		}

		time.Sleep(10 * time.Second)
		var codeIssuesList []model.CodeIssues
		p := 1
		for {
			url := fmt.Sprintf("http://%v/api/issues/search?", sonarAddr)
			url = url + fmt.Sprintf("componentKeys=%v", cdm.ProjectName)
			url = url + fmt.Sprintf("&p=%v", p)
			url = url + "&ps=500"
			url = url + "&s=FILE_LINE&resolved=false&facets=severities%2Ctypes&additionalFields=_all&timeZone=Asia%2FShanghai"

			var client = &http.Client{
				Timeout: time.Second * 5,
			}

			rqst, err := http.NewRequest("GET", url, nil)
			if err != nil {
				logrus.Errorf("new request failure: %v", err)
				return
			}
			rqst.SetBasicAuth(t.config.SonarToken, "")
			rsps, err := client.Do(rqst)
			if err != nil {
				logrus.Errorf("request sonar failure: %v", err)
				return
			}

			body, _ := ioutil.ReadAll(rsps.Body)
			issuesInterface := gojsonq.New().FromString(string(body)).Find("issues")
			total := gojsonq.New().FromString(string(body)).Find("total")
			var ciList []model.CodeIssues
			issuesJson, err := json.Marshal(issuesInterface)
			if err != nil {
				logrus.Errorf("json marshal issues interface failure: %v", err)
				return
			}
			err = json.Unmarshal(issuesJson, &ciList)
			if err != nil {
				logrus.Errorf("json unmarshal issues json failure: %v", err)
				return
			}
			codeIssuesList = append(codeIssuesList, ciList...)
			rsps.Body.Close()
			if p*500 >= int(total.(float64)) {
				break
			}
			p += 1
		}
		var componentReportList []db_model.ComponentReport
		t.db.Debug().Where("component_id = ?", cdm.ProjectName).Delete(&db_model.ComponentReport{})
		for _, code := range codeIssuesList {
			url := fmt.Sprintf("/project/issues?resolved=false&open=%v&id=%v", sonarAddr, code.Key, code.Project)
			level := 1
			if code.Severity == "CRITICAL" || code.Severity == "MAJOR" {
				level = 0
			}
			componentReportList = append(componentReportList, db_model.ComponentReport{
				PrimaryLink: url,
				Level:       level,
				Message:     code.Message,
				ComponentID: code.Project,
				CreateTime:  time.Now(),
				Type:        "code",
			})
		}
		err = t.db.Debug().Create(&componentReportList).Error
		if err != nil {
			logrus.Errorf("create service normative record failure: %v", err)
		}
		return
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
