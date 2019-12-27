package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	middleware "github.com/haodiaodemingzi/cloudfeet/middlewares"
	"github.com/haodiaodemingzi/cloudfeet/models"
	"github.com/haodiaodemingzi/cloudfeet/pkg/e"
	res "github.com/haodiaodemingzi/cloudfeet/pkg/http/response"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	auth "github.com/haodiaodemingzi/cloudfeet/services/auth_service"
)


// @Summary user login
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /admin/user/login [get]
func Login(c *gin.Context) {
	var loginInfo models.LoginInfo
	err := c.BindJSON(&loginInfo)
	log.Info("admin login info : %+v", loginInfo)

	if !auth.ValidateUser(loginInfo.Username, loginInfo.Password) {
		log.Error("用户不存在或者用户密码错误")
		res.Response(c, http.StatusUnauthorized, e.ERROR_LOGIN_FAILED, nil)
		return
	}
	token, err := middleware.GenerateToken(loginInfo.Username, loginInfo.Password)
	if err != nil {
		res.Response(c, http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	log.Info("Token: %s", token)

	res.Response(c, http.StatusOK, e.SUCCESS, data)
}

func UserInfo(c *gin.Context) {
	claims, err := middleware.ParseToken(c.Request.Header.Get("Token"))
	if err != nil {
		log.Error(err.Error())
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	userModel := &models.UserModel{}
	where := map[string]interface{}{
		"username": claims.Username,
	}
	userInfo, _:= userModel.Select(where)
	res.Response(c, http.StatusOK, e.SUCCESS, userInfo)
}




