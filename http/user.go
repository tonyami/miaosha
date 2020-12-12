package http

import (
	"miaosha/lib/code"
	"net/http"
	"strconv"
)

func userInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		newError(w, code.ConvertErr)
		return
	}
	user, err := userSvc.GetInfo(id)
	if err != nil {
		newError(w, err)
		return
	}
	newSuccessWith(w, user)
}
