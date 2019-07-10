package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haodiaodemingzi/cloudfeet/common/e"
	"github.com/haodiaodemingzi/cloudfeet/common/settings"

	res "github.com/haodiaodemingzi/cloudfeet/common/http/response"
)

// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/config/{id} [get]
func GetConfig(c *gin.Context) {
	res.Response(c, http.StatusOK, e.SUCCESS, settings.Config)
}
