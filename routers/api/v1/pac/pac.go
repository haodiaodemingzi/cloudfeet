package pac

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haodiaodemingzi/cloudfeet/common/e"
	"github.com/haodiaodemingzi/cloudfeet/common/logging"
	"github.com/haodiaodemingzi/cloudfeet/services/pac_service"

	res "github.com/haodiaodemingzi/cloudfeet/common/http/response"
)

var logger = logging.GetLogger()

type DomainInfo struct {
	Source  string `json:"source"`
	Domains string `json:"domains"`
}

// CheckedDomain 检测完成之后提交的域名信息
type CheckedDomain struct{
	Source string
	Domains map[string]interface{}
}

// @Summary 上传app搜集的域名
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/pac/domains [post]
func UploadDomain(c *gin.Context) {
	var domainInfo DomainInfo
	err := c.BindJSON(&domainInfo)
	if err != nil {
		logger.Error("post upload domain json data error", domainInfo)
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	err = pac_service.SavePacDomain(domainInfo.Source, domainInfo.Domains)

	if err !=nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}

// @Summary 拉取域名
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/pac/domains [get]
func PullDomain(c *gin.Context) {
	status, _ := strconv.Atoi(c.DefaultQuery("status", "0"))
	limit := c.DefaultQuery("limit", "1000")

	// add check input
	data := map[string]interface{}{
		"limit":  limit,
		"status": status,
	}
	logger.Info("map data item = ", data)
	domains, err := pac_service.GetDomains(data)
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, domains)
}

// @Summary 更新域名检测信息
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/pac/domains [put]
func UpdateDomains(c *gin.Context) {
	var checkedDomain CheckedDomain
	err := c.BindJSON(&checkedDomain)
	if err != nil {
		logger.Error("put json data error", checkedDomain)
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	err = pac_service.RefreshCheckedDomain(&checkedDomain.Domains)

	if err != nil{
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}
