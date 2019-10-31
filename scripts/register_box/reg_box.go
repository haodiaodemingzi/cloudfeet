package main

import (
	"fmt"

	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
	"github.com/haodiaodemingzi/cloudfeet/utils"
)

func init(){
	settings.Setup()
}

func main(){
	//deviceID := "D26384628B2C1237" //ptx
	deviceID := "D26384628B3D1723" // daxue
	//
	passwordMD5 := utils.EncodeMD5(deviceID + `|` + settings.Config.Jwt.Secret)
	fmt.Println("password_md5 -> " + passwordMD5)
}
