package pac_service

import (
	"strconv"
	"strings"
	"time"

	"github.com/haodiaodemingzi/cloudfeet/models"
	log "github.com/haodiaodemingzi/cloudfeet/pkgs/logging"
	"github.com/haodiaodemingzi/cloudfeet/pkgs/settings"
	"github.com/haodiaodemingzi/cloudfeet/utils"
)

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
	var domainList = strings.Split(*domains, ",")

	for _, item := range domainList {
		pac := &models.PacModel{Domain: item, Status: 0, Source: source}
		if err := pac.FindOrCreate(item); err != nil {
			return err
		}
	}

	return nil
}

// SavePacDomain ...
func GetDomains(cond map[string]interface{}) (*[]models.PacModel, error) {
	// 处理分页查询参数，删除分页参数生成 sql 条件
	var limits = "limit"
	var offsets = "offset"
	var offset = cond[offsets].(int)
	var limit = cond[limits].(int)

	delete(cond, limits)
	delete(cond, offsets)

	pacModel := models.PacModel{}
	pacList, err := pacModel.Query(cond, offset, limit)
	log.Debug("paclist = %+v", pacList)
	if err != nil {
		return nil, err
	}

	return &pacList, nil
}

// 刷新域名检测
func RefreshCheckedDomain(data *map[string]string) error {
	var pac *models.PacModel

	for k, v := range *data {
		where := make(map[string]interface{})
		where["domain"] = k
		item, _ := pac.Select(where)
		if item.Domain == "" {
			continue
		}

		item.Status, _ = strconv.Atoi(v)
		item.Edit()
	}
	return nil
}

// 生成盒子配置内nil
func GenBoxConfig() (string, error) {
	pacModel := &models.PacModel{}
	where := make(map[string]interface{})

	// status = 1 表示被墙的域名
	// TODO: status 定义常量
	where["status"] = 1
	pageNum := 0
	pageSize := 99999
	pacList, _ := pacModel.Query(where, pageNum, pageSize)

	var configStr string
	for i := 0; i < len(pacList); i++ {
		configStr += utils.DomainToGFWConf(utils.ParseTopDomain(pacList[i].Domain))
	}
	log.Debug("config pac domain config: %+v", configStr)

	return configStr, nil
}

// 生成盒子配置内nil
func GetBoxStartScript() (string, error) {
	var proxyModel = &models.ProxyModel{}

	proxyConfig, err := proxyModel.RandomProxy()
	if err != nil {
		return "", err
	}
	gfwlistURL := settings.Config.Gin.BaseURL + `pac/config`
	authURL := settings.Config.Gin.BaseURL + `auth/token`

	log.Info("获取proxyconfig配置 = ", proxyConfig)
	log.Info("api url = ", gfwlistURL)

	//gfwURL := settings.R
	content, err := utils.RenderBoxScript(
		proxyConfig.Server, proxyConfig.Name, proxyConfig.Password, proxyConfig.Port,
		proxyConfig.EncryptMethod, authURL, gfwlistURL)
	return content, err
}
