package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/cloudfeet-config/routers"
	"github.com/haodiaodemingzi/cloudfeet/cloudfeet-config/settings"
)

func init() {
	settings.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/EDDYCJY/go-gin-example
// @license.name MIT
// @license.url https://github.com/EDDYCJY/go-gin-example/blob/master/LICENSE
func main() {
	gin.SetMode(gin.DebugMode)
	routersInit := routers.InitRouter()
	// readTimeout := setting.ServerSetting.ReadTimeout
	// writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", settings.Cfg.GetInt("gin.port"))
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    60,
		WriteTimeout:   60,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
