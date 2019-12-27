package user_service

import (
	"github.com/haodiaodemingzi/cloudfeet/models"
)


// SavePacDomain ...
func GetUserInfo(username string) (*models.UserModel, error){
	var userModel models.UserModel
	where := make(map[string]interface{})
	where["username"] = username

	return userModel.Select(where)
}


