package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type accessTokenInfo struct {
	Species     string
	Description string
}

func main() {

	router := gin.Default()

	router.GET("/check", func(c *gin.Context) {
		server := c.Query("server")
		port := c.DefaultQuery("port", "9090")

		apiPrefix := fmt.Sprintf("https://%s:%s/abcd123", server, port)
		getConfigURL := apiPrefix + "/access-keys/"
		go func() {
			fmt.Print(getConfigURL)
		}()

		c.String(http.StatusOK, "Hello World")
	})
	router.Run(":8083")
}
