package proxy

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/pkg/e"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	proxyService "github.com/haodiaodemingzi/cloudfeet/services/proxy_service"

	"github.com/haodiaodemingzi/cloudfeet/middlewares"
	res "github.com/haodiaodemingzi/cloudfeet/pkg/http/response"
)

type ProxyInfo struct {
	Method string `form:"method" json:"method" binding:"required"`
	Port int `form:"port" json:"port" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @Summary Test a mysql conn api
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/config/mysql [get]
func GetProxy(c *gin.Context) {
	claims, err := middlewares.ParseToken(c.Request.Header.Get("Token"))
	if err != nil{
		log.Error("claims err %s", err.Error())
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}

	if claims == nil {
		log.Error("get user claims nil")
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
	}
	connInfo, err := proxyService.ProxyConnInfo(claims.Username)
	fmt.Println(err)
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, connInfo)
}

// @Summary register a proxy from node
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/proxy [post]
func RegisterProxy(c *gin.Context) {
	var proxyInfo ProxyInfo
	err := c.ShouldBindJSON(&proxyInfo)
	if err != nil{
		log.Error("绑定form表单参数错误")
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	log.Info("ss配置: %+v", proxyInfo)
	err = proxyService.AddProxy(
		c.ClientIP(), proxyInfo.Port, proxyInfo.Method, proxyInfo.Password)
	if err != nil {
		log.Error("添加proxy失败 %s", err.Error())
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}

// @Summary delete a proxy from node
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/proxy [post]
func DeleteProxy(c *gin.Context) {
	server := c.Param("server")
	if server == "" {
		log.Error("没有指定删除节点")
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	log.Info("server -> %s", server)
	err := proxyService.RemoveProxy(server)
	if err != nil {
		log.Error("删除proxy失败 %s", err.Error())
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}
