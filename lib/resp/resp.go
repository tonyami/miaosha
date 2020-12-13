package resp

import (
	"github.com/gin-gonic/gin"
	"miaosha/lib/code"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewSuccess(c *gin.Context) {
	NewSuccessWith(c, nil)
}

func NewSuccessWith(c *gin.Context, data interface{}) {
	newJson(c, code.Success, data)
}

func NewError(c *gin.Context, err error) {
	newJson(c, code.ToCode(err.Error()), nil)
}

func newJson(c *gin.Context, code code.Code, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: code.Code(),
		Msg:  code.Message(),
		Data: data,
	})
}
