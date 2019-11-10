package proxy_service

import (
	"errors"
	"fmt"
	"gopkg.in/resty.v1"

	httpclient "github.com/go-resty/resty/v2"

	"github.com/haodiaodemingzi/cloudfeet/models"
	"github.com/haodiaodemingzi/cloudfeet/pkg/consul"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
)


// 随机获取代理链接信息
// TODO: 后期智能根据负载和网络状况返回
// get proxy info from consul service
func ProxyConnInfo(id string) (models.ProxyModel, error) {
	/*
	var model = &models.ProxyModel{}
	proxyModel, err := model.RandomProxy()
	if err != nil {
		log.Error("获取ss配置失败: ", err.Error())
	}
	*/
	var proxyModel models.ProxyModel
	service, err := consul.GetRandomProxyService("outline-proxy")
	if err != nil{
		return models.ProxyModel{}, err
	}
	if service == nil{
		return models.ProxyModel{}, errors.New("can't find service for outline-proxy")
	}

	client := httpclient.New()
	outlineAPI := "https://" + service.Address + ":8081/access-keys"
	payload := map[string]interface{}{
		"id": "outline-" + id,
		"port": 10247,
		"password": "Divein" + id,
	}
	resp, err := client.R().SetBody(payload).Post(outlineAPI)
	if resp !=nil && resp.StatusCode() == 201{
		return proxyModel, errors.New("get proxy failed")
	}
	proxyModel.Server = service.Address
	proxyModel.Port = 10248
	proxyModel.Name = "outline-proxy-" + service.Address
	proxyModel.EncryptMethod = "chacha20-poly1305"
	proxyModel.Password = "Diveinedu"

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

func AddOutlineProxy() (map[string]interface{},error){
	nodeService, err:= consul.GetRandomProxyService("outline-proxy")
	if err != nil{
		log.Error(err.Error())
		return nil, err
	}
	// with auth api key
	apiKEY := "api"
	outlineAPI := fmt.Sprintf("https://%s:%d/%s/access-key",
		nodeService.Address, nodeService.Port, apiKEY)
	client := resty.New()
	// read from setting
	body := map[string]interface{}{"username": "testuser", "password": "Diveinedu",}
	resp, err := client.R().SetBody(body).Post(outlineAPI)
	if resp.StatusCode() != 201{
		return nil, errors.New("post new access-key failed")
	}
	return body, nil
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
