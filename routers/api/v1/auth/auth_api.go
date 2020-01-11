package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	middleware "github.com/haodiaodemingzi/cloudfeet/middlewares"
	"github.com/haodiaodemingzi/cloudfeet/pkg/e"
	res "github.com/haodiaodemingzi/cloudfeet/pkg/http/response"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	auth "github.com/haodiaodemingzi/cloudfeet/services/auth_service"
)


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
	log.Info("login info : %+v", loginInfo)

	if !auth.ValidateUser(loginInfo.Username, loginInfo.Password) {
		log.Error("用户不存在或者用户密码错误")
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	token, err := middleware.GenerateToken(loginInfo.Username, loginInfo.Password)
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	log.Info("Token: %s", token)

	res.Response(c, http.StatusOK, e.SUCCESS, data)
}
