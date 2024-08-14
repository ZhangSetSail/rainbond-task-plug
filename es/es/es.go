package es

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	db_model "github.com/goodrain/rainbond-task-plug/db/model"
	"github.com/goodrain/rainbond-task-plug/pkg/config"
	"net/http"
)

var esc *ComponentReportRepo

type ComponentReportRepo struct {
	EsURL      string
	EsIndex    string
	EsUsername string
	EsPassword string
	Client     *http.Client
}

func InitES(esConfig *config.ESConfig) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	esc = &ComponentReportRepo{
		EsURL:      esConfig.EsURL,
		EsIndex:    esConfig.EsIndex,
		EsUsername: esConfig.EsUsername,
		EsPassword: esConfig.EsPassword,
		Client:     &http.Client{Transport: tr},
	}
}

func (repo *ComponentReportRepo) DeleteComponentReports(componentID, reportType string) error {
	if repo.EsURL == "" {
		return fmt.Errorf("ES_URL is not set")
	}
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"component_id": componentID,
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"type": reportType,
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(query)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s/_delete_by_query", repo.EsURL, repo.EsIndex), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.SetBasicAuth(repo.EsUsername, repo.EsPassword)
	req.Header.Set("Content-Type", "application/json")

	resp, err := repo.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete component reports, status code: %d", resp.StatusCode)
	}

	return nil
}

func (repo *ComponentReportRepo) CreateComponentReports(reports []*db_model.ComponentReport) error {
	if repo.EsURL == "" {
		return fmt.Errorf("ES_URL is not set")
	}

	var bulkBody bytes.Buffer
	for _, report := range reports {
		meta := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": repo.EsIndex,
				"_type":  "_doc",
			},
		}
		metaBytes, _ := json.Marshal(meta)
		reportBytes, _ := json.Marshal(report)
		bulkBody.Write(metaBytes)
		bulkBody.Write([]byte("\n"))
		bulkBody.Write(reportBytes)
		bulkBody.Write([]byte("\n"))
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/_bulk", repo.EsURL), &bulkBody)
	if err != nil {
		return err
	}
	req.SetBasicAuth(repo.EsUsername, repo.EsPassword)
	req.Header.Set("Content-Type", "application/json")

	resp, err := repo.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create component reports, status code: %d", resp.StatusCode)
	}

	return nil
}

func GetES() *ComponentReportRepo {
	return esc
}
