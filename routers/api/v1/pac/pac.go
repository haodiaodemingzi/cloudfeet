package pac

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"

	"github.com/haodiaodemingzi/cloudfeet/pkgs/e"
	"github.com/haodiaodemingzi/cloudfeet/pkgs/logging"
	"github.com/haodiaodemingzi/cloudfeet/services/pac_service"

	res "github.com/haodiaodemingzi/cloudfeet/pkgs/http/response"
)

var logger = logging.GetLogger()

type DomainInfo struct {
	Source  string `json:"source"`
	Domains string `json:"domains"`
}

// CheckedDomain 检测完成之后提交的域名信息
type CheckedDomain struct {
	Source  string
	Domains map[string]string
}

// @Summary 上传app搜集的域名
// @Produce  json
// @Success 200 {object} response.Template
// @Failure 500 {object} response.Template
// @Router /api/v1/pac/domains [post]
func UploadDomains(c *gin.Context) {
	var domainInfo DomainInfo
	err := c.BindJSON(&domainInfo)
	if err != nil {
		logger.Error("post upload domain json data error", domainInfo)
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	err = pac_service.SavePacDomain(domainInfo.Source, &domainInfo.Domains)

	if err != nil {
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
func PullDomains(c *gin.Context) {
	status, _ := strconv.Atoi(c.DefaultQuery("status", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5000"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	if limit > 5000 {
		log.Error("domain pull max = 5000")
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}

	// add check input
	data := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"status": status,
	}

	logger.Info("map data item = ", data)
	domainList := []string{}
	pacList, err := pac_service.GetDomains(data)
	for _, pac := range *pacList {
		domainList = append(domainList, pac.Domain)
	}
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, domainList)
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
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}

// UploadDomains ...
func UploadDomainFile(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		logger.Fatal(err.Error())
		res.Response(c, http.StatusOK, e.ERROR, nil)
		return
	}
	var buf bytes.Buffer

	_, err = io.Copy(&buf, file)
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}
	content := buf.String()
	scanner := bufio.NewScanner(strings.NewReader(content))
	var domainList []string
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			domainList = append(domainList, line)
		}
	}
	domains := strings.Join(domainList, ",")
	err = pac_service.SavePacDomain("box", &domains)
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}

	res.Response(c, http.StatusOK, e.SUCCESS, nil)
}

// 下载盒子pac文件
func DownloadBoxConfig(c *gin.Context) {
	// status = 1 被屏蔽的域名状态
	domainLines, err := pac_service.GenBoxConfig()
	if err != nil {
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}

	c.String(http.StatusOK, domainLines)
}

// 下载盒子启动脚本
func DownloadBoxScript(c *gin.Context) {
	// status = 1 被屏蔽的域名状态
	content, err := pac_service.GetBoxStartScript()
	if err != nil {
		log.Error(err)
		res.Response(c, http.StatusBadRequest, e.ERROR, nil)
		return
	}

	c.String(http.StatusOK, content)
}
