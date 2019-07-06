package v1

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/articles/{id} [get]
func GetConfig(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "get config successfully!"})
}

