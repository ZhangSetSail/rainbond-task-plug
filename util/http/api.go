package http

import (
	"context"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"reflect"
)

// ReturnError 返回错误信息
func ReturnError(r *http.Request, w http.ResponseWriter, code int, msg string) {
	logrus.Debugf("error code: %d; error uri: %s; error msg: %s", code, r.RequestURI, msg)
	r = r.WithContext(context.WithValue(r.Context(), render.StatusCtxKey, code))
	render.DefaultResponder(w, r, ResponseBody{Msg: msg})
}

// ReturnSuccess 成功返回
func ReturnSuccess(r *http.Request, w http.ResponseWriter, datas interface{}) {
	r = r.WithContext(context.WithValue(r.Context(), render.StatusCtxKey, http.StatusOK))
	if datas == nil {
		render.DefaultResponder(w, r, ResponseBody{Bean: nil})
		return
	}
	v := reflect.ValueOf(datas)
	if v.Kind() == reflect.Slice {
		render.DefaultResponder(w, r, ResponseBody{List: datas})
		return
	}
	render.DefaultResponder(w, r, ResponseBody{Bean: datas})
	return
}

// ResponseBody api返回数据格式
type ResponseBody struct {
	ValidationError url.Values  `json:"validation_error,omitempty"`
	Msg             string      `json:"msg,omitempty"`
	Bean            interface{} `json:"bean,omitempty"`
	List            interface{} `json:"list,omitempty"`
	//数据集总数
	ListAllNumber int `json:"number,omitempty"`
	//当前页码数
	Page int `json:"page,omitempty"`
}

type Response struct {
	Message string      `json:"msg"`
	Code    int         `json:"code"`
	Bean    interface{} `json:"bean"`
	List    interface{} `json:"list"`
}
