package response

import (
	"github.com/gin-gonic/gin"
	"github.com/haodiaodemingzi/cloudfeet/common/e"
)

type API struct {
	C *gin.Context
}

type Template struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Template{
		Code: httpCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
