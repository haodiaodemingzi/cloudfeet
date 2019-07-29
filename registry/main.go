package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/common/gmysql"
	"github.com/haodiaodemingzi/cloudfeet/common/logging"
	"github.com/haodiaodemingzi/cloudfeet/common/settings"
	"github.com/haodiaodemingzi/cloudfeet/config/routers"
)

func init() {
	settings.Setup()
	gmysql.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An config api
// @termsOfService https://github.com/haodiaodemingzi/cloudfeet
func main() {
	gin.SetMode(gin.DebugMode)
	r := routers.InitRouter()
	endPoint := fmt.Sprintf("%s:%d", settings.Config.Gin.Host, settings.Config.Gin.Port)
	logging.Info("Start cloudfeet-config web service with endpoint: %s", endPoint)

	r.Run(endPoint)
}
