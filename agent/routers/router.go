package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/haodiaodemingzi/cloudfeet/config/docs"
	v1 "github.com/haodiaodemingzi/cloudfeet/config/routers/api/v1"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/files"
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

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/config", v1.GetConfig)
	apiv1.GET("/query", v1.QueryUser)

	return r
}
