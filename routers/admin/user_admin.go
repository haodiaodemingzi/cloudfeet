package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/pkgs/e"
	res "github.com/haodiaodemingzi/cloudfeet/pkgs/http/response"
	log "github.com/haodiaodemingzi/cloudfeet/pkgs/logging"
	user "github.com/haodiaodemingzi/cloudfeet/services/user_service"
	"github.com/haodiaodemingzi/cloudfeet/utils"
)

// @Summary user login
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /admin/user/login [get]
func UserInfo(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		log.Error("解析 jwt token 失败")
		res.Response(c, http.StatusUnauthorized, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	userInfo, _ := user.GetUserInfo(claims.Username)
	log.Info("userinfo = %+v", userInfo)
	if userInfo.ID > 0 {
		res.Response(c, http.StatusOK, e.SUCCESS, userInfo)
		return
	}
}
