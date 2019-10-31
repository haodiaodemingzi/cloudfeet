package proxy_service

import (
	"fmt"

	"github.com/haodiaodemingzi/cloudfeet/models"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
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

func AddProxy(server string, port int, method string, password string) error{
	var model = &models.ProxyModel{}

	model.Name = fmt.Sprintf("%s-%d", server, port)
	model.Domain = server
	model.Server = server
	model.Port = port
	model.EncryptMethod = method
	model.Password = password
	model.Status = 1

	return model.FindOrCreate(model.Domain)
}

func RemoveProxy(server string) error{
	var model = &models.ProxyModel{}

	where := map[string]interface{}{"server": server,}
	row, err := model.Select(where)
	if err != nil{
		return err
	}

	err = row.Delete()
	return err
}
