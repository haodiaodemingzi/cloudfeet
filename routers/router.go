package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/haodiaodemingzi/cloudfeet/middlewares"
	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
	"github.com/haodiaodemingzi/cloudfeet/routers/admin"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/auth"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/node"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/pac"
	"github.com/haodiaodemingzi/cloudfeet/routers/api/v1/proxy"

	//"github.com/swaggo/gin-swagger/swaggerFiles"
	swaggerFiles "github.com/swaggo/files"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.JwtMiddleware())
	// Apply the middleware to the router (works with groups too)
	/*
	r.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin, Token",
		ExposedHeaders: "",
		MaxAge: 50 * time.Second,
		Credentials: true,
		ValidateHeaders: false,
	}))
	*/
	r.Use(middlewares.CorsMiddleware())
	r.Use(gin.Recovery())


	// swagger 文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// TODO: read api version group from config
	api := r.Group("/api/v1")
	//api.GET("/config", config.GetConfig)
	//api.GET("/query", config.QueryUser)

	// auth api maps
	//api.POST("/auth/token", auth.GenToken)
	api.POST(settings.Config.URL.AuthToken, auth.GenToken)

	// proxy api maps
	//api.GET("/proxy", proxy.GetProxy)
	api.GET(settings.Config.URL.ProxyInfo, proxy.GetProxy)
	api.POST(settings.Config.URL.ProxyInfo, proxy.RegisterProxy)
	api.DELETE(settings.Config.URL.ProxyInfo + "/:server", proxy.DeleteProxy)

	// pac api maps
	//api.POST("/pac/domains", pac.UploadDomains)
	api.POST(settings.Config.URL.UploadDomains, pac.UploadDomains)

	//api.GET("/pac/domains", pac.PullDomains)
	api.GET(settings.Config.URL.PullDomains, pac.PullDomains)
	//api.PUT("/pac/domains", pac.UpdateDomains)
	api.PUT(settings.Config.URL.UpdateDomains, pac.UpdateDomains)
	//api.POST("/pac/domains/cache", pac.UploadDomainFile)
	api.POST(settings.Config.URL.UploadDNSFile, pac.UploadDomainFile)

	// api for box
	//api.GET("/pac/script", pac.DownloadBoxScript)
	//api.GET("/pac/config", pac.DownloadBoxConfig)
	api.GET(settings.Config.URL.PacConfig, pac.DownloadBoxConfig)
	api.GET(settings.Config.URL.InitScript, pac.DownloadBoxScript)

	// api for outline node
	api.POST(settings.Config.URL.Node, node.RegisterNode)

	// admin api
	web := r.Group("/admin")
	web.POST("/user/login", admin.Login)
	web.GET("/user/info", admin.UserInfo)

	return r
}
