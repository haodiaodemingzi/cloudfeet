package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func CorsMiddleware() gin.HandlerFunc {
	// TODO: 从配置文件读取cors配置
	config := cors.Config{
		 Origins:        "*",
		 Methods:        "GET, PUT, POST, DELETE, OPTIONS",
		 RequestHeaders: "Origin, Authorization, Content-Type, Access-Control-Allow-Origin, Token",
		 ExposedHeaders: "",
		 MaxAge: 50 * time.Second,
		 Credentials: true,
		 ValidateHeaders: false,
	}
	middleware := cors.Middleware(config)
	return middleware
}
