package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/pkgs/e"
	"github.com/haodiaodemingzi/cloudfeet/pkgs/settings"

	res "github.com/haodiaodemingzi/cloudfeet/pkgs/http/response"
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

// @Summary Test a mysql conn api
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/config/mysql [get]
func QueryUser(c *gin.Context) {
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}
