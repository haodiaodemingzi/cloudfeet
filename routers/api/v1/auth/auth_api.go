package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haodiaodemingzi/cloudfeet/common/logging"
	middleware "github.com/haodiaodemingzi/cloudfeet/middlewares"
)

var logger = logging.GetLogger()

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary 获取 api token
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/auth/token [get]
func GenToken(c *gin.Context) {
	var loginInfo LoginInfo
	err := c.BindJSON(&loginInfo)
	logger.Info("loginInfo username ", loginInfo.Username)

	// TODO: check user password in db model
	var msg = ""
	token, err := middleware.GenerateToken(loginInfo.Username, loginInfo.Password)
	if err != nil {
		logger.Info("get token failed :", err.Error())
		msg = "生成token失败"
		c.JSON(400, gin.H{"code": 400, "msg": msg, "data": nil})
		return
	}
	data := make(map[string]interface{})
	data["token"] = token

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "", "data": data})
}
