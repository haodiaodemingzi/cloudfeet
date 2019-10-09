package proxy_service

import (
	"github.com/haodiaodemingzi/cloudfeet/models"
	log "github.com/haodiaodemingzi/cloudfeet/pkgs/logging"
)


// 随机获取代理链接信息
// TODO: 后期智能根据负载和网络状况返回
func ProxyConnInfo() (models.ProxyModel, error) {
	var model = &models.ProxyModel{}
	proxyModel, err := model.RandomProxy()
	if err != nil {
		log.Error("获取ss配置失败: ", err.Error())
	}

	return proxyModel, err
}
