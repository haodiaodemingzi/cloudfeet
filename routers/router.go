package routers

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/haodiaodemingzi/cloudfeet/middlewares"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/auth"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/config"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/pac"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/proxy"

	//"github.com/swaggo/gin-swagger/swaggerFiles"
	swaggerFiles "github.com/swaggo/files"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(middlewares.JwtMiddleware())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	// swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	api.GET("/config", config.GetConfig)
	api.GET("/query", config.QueryUser)

	// auth api maps
	api.POST("/auth/token", auth.GenToken)

	// proxy api maps
	api.GET("/proxy", proxy.GetProxy)

	// pac api maps
	api.POST("/pac/domains", pac.UploadDomains)

	api.GET("/pac/domains", pac.PullDomains)
	api.PUT("/pac/domains", pac.UpdateDomains)
	api.POST("/pac/domains/file", pac.UploadDomainFile)

	// api for box
	api.GET("/pac/script", pac.DownloadBoxScript)
	api.GET("/pac/config", pac.DownloadBoxConfig)

	return r
}
