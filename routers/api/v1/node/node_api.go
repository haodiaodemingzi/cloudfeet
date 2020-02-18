package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"

	"github.com/haodiaodemingzi/cloudfeet/pkg/e"
	res "github.com/haodiaodemingzi/cloudfeet/pkg/http/response"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/haodiaodemingzi/cloudfeet/services/node_service"
)

//NodeInfo
type NodeInfo struct {
	Server   string `form:"server" json:"server"`
	Port     int    `form:"port" json:"port"`
	Provider string `form:"provider" json:"provider"`
	Region   string `form:"region" json:"region"`
}

// @Summary register a node into consul service with custom port
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/node [post]
func RegisterNode(c *gin.Context) {
	var nodeInfo NodeInfo
	c.BindJSON(&nodeInfo)
	/*
	err := c.ShouldBindJSON(&nodeInfo)
	if err != nil {
		log.Error("node 绑定form表单参数错误")
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	log.Info("node配置: %+v", nodeInfo)
	*/
	log.Info("node配置: %+v", nodeInfo)

	// outline api server from c.request port is 8081
	if nodeInfo.Server == "" {
		nodeInfo.Server = c.ClientIP()
	}
	if nodeInfo.Port == 0 {
		nodeInfo.Port = 8081
	}

	err := node_service.AddNode(
		nodeInfo.Server, nodeInfo.Port, nodeInfo.Provider, nodeInfo.Region)
	if err != nil {
		log.Error("添加 node err %s", err.Error())
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}

// @Summary get healty node info
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/node [get]
func GetNodeList(c *gin.Context) {
	serviceName := `outline-proxy`
	if serviceName == "" {
		res.Response(c, http.StatusNotFound, e.ERROR, nil)
		return
	}

	checkedServices, err := node_service.GetNodeList(serviceName)
	if err != nil {
		res.Response(c, http.StatusNotFound, e.ERROR, nil)
		return
	}
	var nodeList []*api.AgentService
	for _, item := range checkedServices {
		nodeList = append(nodeList, item.Service)
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nodeList)
}
