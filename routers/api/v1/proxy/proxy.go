package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haodiaodemingzi/cloudfeet/common/e"
	"github.com/haodiaodemingzi/cloudfeet/common/logging"
	proxyService "github.com/haodiaodemingzi/cloudfeet/services/proxy_service"

	res "github.com/haodiaodemingzi/cloudfeet/common/http/response"
)

var logger = logging.GetLogger()

// @Summary Test a mysql conn api
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/config/mysql [get]
func GetProxy(c *gin.Context) {
	id := "1"
	port := "7007"
	server := "ss.csdc.io"
	method := "chacha20"
	connInfo := proxyService.ProxyConnInfo{
		ID: id, Server: server, Port: port, Method: method,
	}
	logger.Info("get conn resp ", connInfo)

	res.Response(c, http.StatusOK, e.SUCCESS, connInfo)
}
