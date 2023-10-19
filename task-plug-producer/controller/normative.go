package controller

import (
	"encoding/json"
	"fmt"
	"github.com/goodrain/rainbond-task-plug/task-plug-producer/handle"
	httputil "github.com/goodrain/rainbond-task-plug/util/http"
	"github.com/sirupsen/logrus"
	"net/http"
)

type NormativeController struct{}

func (n NormativeController) RetrieveData(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	serviceID := values.Get("service_id")
	serviceIDListJson := values.Get("service_list")
	var serviceIDList []string
	if serviceIDListJson != "" {
		err := json.Unmarshal([]byte(serviceIDListJson), &serviceIDList)
		if err != nil {
			logrus.Errorf("json unmarshal service id list failure: %v", err)
			httputil.ReturnError(r, w, 500, fmt.Sprintf("retrive normative data failure: %v", err))
		}
	}

	rnd, err := handle.GetDBHandle().RetrieveNormativeData(serviceID, serviceIDList)
	if err != nil {
		logrus.Errorf("retrive normative data failure: %v", err)
		httputil.ReturnError(r, w, 500, fmt.Sprintf("retrive normative data failure: %v", err))
	}

	httputil.ReturnSuccess(r, w, &httputil.Response{
		Message: "请求成功",
		Code:    http.StatusOK,
		List:    rnd,
	})
	return
}
