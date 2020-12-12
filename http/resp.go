package http

import (
	"encoding/json"
	"miaosha/lib/code"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func newSuccess(w http.ResponseWriter) {
	newSuccessWith(w, nil)
}

func newSuccessWith(w http.ResponseWriter, data interface{}) {
	resp(w, code.Success, data)
}

func newError(w http.ResponseWriter, err error) {
	resp(w, code.String(err.Error()), nil)
}

func resp(w http.ResponseWriter, code code.Code, data interface{}) {
	res := Response{
		Code: code.Code(),
		Msg:  code.Message(),
		Data: data,
	}
	bytes, _ := json.Marshal(res)
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Headers", "none")
	//w.Header().Set("Access-Control-Allow-Methods", "none")
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
