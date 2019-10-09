package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/models"
	"github.com/haodiaodemingzi/cloudfeet/pkgs/logging"
	"github.com/haodiaodemingzi/cloudfeet/pkgs/settings"
	"github.com/haodiaodemingzi/cloudfeet/routers"
)

var log = logging.GetLogger()

func init() {
	settings.Setup()
	models.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An config api
// @termsOfService https://github.com/haodiaodemingzi/cloudfeet
func main() {
	gin.SetMode(gin.DebugMode)

	r := routers.InitRouter()

	endPoint := fmt.Sprintf("%s:%d", settings.Config.Gin.Host, settings.Config.Gin.Port)
	log.Info("Start cloudfeet-config web service with endpoint: %s", endPoint)

	_ = r.Run(endPoint)
}
