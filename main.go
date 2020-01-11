package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/models"
	"github.com/haodiaodemingzi/cloudfeet/pkg/consul"
	"github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
	"github.com/haodiaodemingzi/cloudfeet/routers"
)

func init() {
	settings.Setup()
	models.Setup()
	logging.Setup()
	consul.Setup()

}

// @title Golang Gin API
// @version 1.0
// @description An config api
// @termsOfService https://github.com/haodiaodemingzi/cloudfeet
func main() {
	gin.SetMode(gin.DebugMode)

	r := routers.InitRouter()

	endPoint := fmt.Sprintf("%s:%d", settings.Config.Gin.Host, settings.Config.Gin.Port)
	_ = r.Run(endPoint)
}
