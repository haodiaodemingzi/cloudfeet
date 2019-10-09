package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haodiaodemingzi/cloudfeet/pkgs/e"
	"github.com/haodiaodemingzi/cloudfeet/pkgs/logging"
	proxyService "github.com/haodiaodemingzi/cloudfeet/services/proxy_service"

	res "github.com/haodiaodemingzi/cloudfeet/pkgs/http/response"
)

var logger = logging.GetLogger()

// @Summary Test a mysql conn api
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/config/mysql [get]
func GetProxy(c *gin.Context) {
	connInfo, err := proxyService.ProxyConnInfo()
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
	}
	res.Response(c, http.StatusOK, e.SUCCESS, connInfo)
}
