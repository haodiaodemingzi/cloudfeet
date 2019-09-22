package pac_service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/haodiaodemingzi/cloudfeet/common/logging"
	"github.com/haodiaodemingzi/cloudfeet/common/settings"
	"github.com/haodiaodemingzi/cloudfeet/common/utils"
	"github.com/haodiaodemingzi/cloudfeet/models"
)

var logger = logging.GetLogger()

type Pac struct {
	ID         int       `db:"id"`
	Domain     string    `db:"domain"`
	Region     string    `db:"region"`
	Status     string    `db:"status"`
	Source     string    `db:"source"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// SavePacDomain ...
func SavePacDomain(source string, domains *string) error {
	var pac = &models.Pac{}
	var domainList = strings.Split(*domains, ",")
	var data []map[string]interface{}
	for _, item := range domainList {
		data = append(data, map[string]interface{}{
			"domain": item,
			"status": 0,
			"source": source,
		})
	}
	_, err := pac.AddDomains(data)
	return err
}

// SavePacDomain ...
func GetDomains(cond map[string]interface{}) ([]string, error) {
	var pac = &models.Pac{}
	pacList, err := pac.Query(cond)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	var domainList []string
	for i := 0; i < len(pacList); i++ {
		domainList = append(domainList, pacList[i].Domain)
	}

	return domainList, nil
}

// 刷新域名检测
func RefreshCheckedDomain(data *map[string]interface{}) error {
	var pac = &models.Pac{}
	return pac.UpdateCheckedDomain(data)
}

// 生成盒子配置内nil
func GenBoxConfig() (string, error) {
	var domainList []string

	var data = make(map[string]interface{})
	data["limit"] = 900000
	data["status"] = 1
	domainList, _ = GetDomains(data)
	var configStr string
	for _, domain := range domainList {
		configStr += utils.DomainToGFWConf(utils.ParseTopDomain(domain))
		logger.Debug("gen pac line ", configStr)
	}
	if len(configStr) <= 0 {
		return "", errors.New("gen pac line failed")
	}
	return configStr, nil
}

// 生成盒子配置内nil
func GetBoxStartScript() (string, error) {
	// TODO: 从数据库获取可用的ss服务器配置
	var proxy = &models.Proxy{}
	var where = map[string]interface{}{
		"status": 1, "limit": 1,
	}
	var proxyList []models.Proxy
	proxyList, err := proxy.Query(where)
	if err != nil {
		return "", err
	}
	if len(proxyList) <= 0 {
		return "", errors.New("没有找到proxy配置")
	}
	proxyConfig := proxyList[0]

	gfwlistURL := settings.Config.Gin.BaseURL + `pac/config`
	authURL := settings.Config.Gin.BaseURL + `auth/token`
	logger.Debug(settings.Config)
	logger.Info("获取proxyconfig配置 = ", proxyConfig)
	logger.Info("api url = ", gfwlistURL)
	//gfwURL := settings.R
	content, err := utils.RenderBoxScript(
		proxyConfig.Server, proxyConfig.Name, proxyConfig.Password, proxyConfig.Port,
		proxyConfig.EncryptMethod, authURL, gfwlistURL)
	return content, err
}
