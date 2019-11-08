package node

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/haodiaodemingzi/cloudfeet/pkg/e"
	res "github.com/haodiaodemingzi/cloudfeet/pkg/http/response"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/haodiaodemingzi/cloudfeet/services/node_service"
)

type NodeInfo struct {
	Server string `form:"server" json:"server" binding:"required"`
	Port int `form:"port" json:"port" binding:"required"`
	Provider string `form:"provider" json:"provider" binding:"required"`
	Region string `form:"region" json:"region" binding:"required"`
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
	err = node_service.AddNode(
		nodeInfo.Server, nodeInfo.Port, nodeInfo.Provider, nodeInfo.Region)
	if err != nil {
		log.Error("添加 node err %s", err.Error())
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}




