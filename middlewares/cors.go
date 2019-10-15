package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "true")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Token, Content-Type, Access-Control-Allow-Origin, Content-Length, accept, origin, " +
				"Cache-Control, X-Requested-With, Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELET, POST, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		fmt.Println("loging in cors middleware.....")

		c.Next()
	}
}
