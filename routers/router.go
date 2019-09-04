package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/config"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/pac"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/proxy"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"

	//"github.com/swaggo/gin-swagger/swaggerFiles"
	swaggerFiles "github.com/swaggo/files"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	// format log
	var log = logrus.New()
	log.Out = os.Stdout
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	api.GET("/config", config.GetConfig)
	api.GET("/query", config.QueryUser)

	// proxy api maps
	api.GET("/proxy", proxy.GetProxy)

	// pac api maps
	api.POST("/pac/domains", pac.UploadDomain)
	api.GET("/pac/domains", pac.PullDomain)
	api.PUT("/pac/domains", pac.UpdateDomains)

	return r
}
