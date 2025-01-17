package auth_service

import (
	"github.com/haodiaodemingzi/cloudfeet/models"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
)


func ValidateUser(username string, password string) bool {
	var userModel models.UserModel
	where := map[string]interface{}{"username": username}
	userInfo, err := userModel.Select(where)

	if userInfo.ID == 0 || err != nil {
		return false
	}

	log.Info("Userinfo : %+v", userInfo)
	// 验证密码
	if userInfo.Password == password {
		return true
	}

	return false
}
