package node

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/pkg/e"
	res "github.com/haodiaodemingzi/cloudfeet/pkg/http/response"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/haodiaodemingzi/cloudfeet/services/node_service"
)

//NodeInfo
type NodeInfo struct {
	Server string `form:"server" json:"server"`
	Port int `form:"port" json:"port"`
	Provider string `form:"provider" json:"provider"`
	Region string `form:"region" json:"region"`
}


// @Summary register a node into consul service with custom port
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/node [post]
func RegisterNode(c *gin.Context) {
	var nodeInfo NodeInfo
	err := c.ShouldBindJSON(&nodeInfo)
	if err != nil{
		log.Error("node 绑定form表单参数错误")
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	log.Info("node配置: %+v", nodeInfo)
	// outline api server from c.request port is 8081
	if nodeInfo.Server == "" {
		nodeInfo.Server = c.ClientIP()
	}
	if nodeInfo.Port == 0 {
		nodeInfo.Port = 8081
	}

	err = node_service.AddNode(
		nodeInfo.Server, nodeInfo.Port, nodeInfo.Provider, nodeInfo.Region)
	if err != nil {
		log.Error("添加 node err %s", err.Error())
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}




