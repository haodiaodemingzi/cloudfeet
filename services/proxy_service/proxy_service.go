package proxy_service

import (
	"crypto/tls"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/resty.v1"

	httpclient "github.com/go-resty/resty/v2"

	"github.com/haodiaodemingzi/cloudfeet/models"
	"github.com/haodiaodemingzi/cloudfeet/pkg/consul"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
)

//ProxyConnInfo ...
// 随机获取代理链接信息
// TODO: 后期智能根据负载和网络状况返回
// get proxy info from consul service
func ProxyConnInfo(username string) (models.ProxyModel, error) {
	var proxyModel models.ProxyModel
	// var userModel models.UserModel
	service, err := consul.GetRandomProxyService("outline-proxy")
	if err != nil {
		return models.ProxyModel{}, err
	}
	if service == nil {
		return models.ProxyModel{}, errors.New("can't find service for outline-proxy")
	}

	client := httpclient.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// TODO: add api key from settings
	outlineAPIKey := settings.Config.Outline.APIKEY
	outlineAPIPort := settings.Config.Outline.Port
	outlineAPI := "https://" + service.Address + ":" + outlineAPIPort + "/" + outlineAPIKey + "/access-keys"
	// userInfo, _ := userModel.Select(map[string]interface{}{"username": username})
	// will cause outline api create key bug
	// outlineID := strings.Join([]string{"outline", username[0:3], userInfo.Comment}, "-")
	outlineID := strings.Join([]string{"outline", username}, "-")

	log.Info("outline api - %s", outlineAPI)
	// TODO: add salt from config in outline password, add port from outline
	payload := map[string]interface{}{
		"id":       outlineID,
		"port":     10247,
		"password": "Divein" + username,
	}
	log.Info("payload json = %+v", payload)

	//resp, err := client.R().SetBody(payload).Post(outlineAPI)

	/*
	if err != nil || resp.StatusCode() != 201 {
		return proxyModel, errors.New("get proxy failed")
	}
		proxyModel.Server = service.Address
		proxyModel.Port = payload["port"].(int)
		proxyModel.Name = payload["id"].(string)
		proxyModel.EncryptMethod = "chacha20-ietf-poly1305"
		proxyModel.Password = payload["password"].(string)
	*/
	// 写死大茶的地址
	proxyModel.Server = "box.csdc.io"
	proxyModel.Port = 8788
	proxyModel.Name = "bigtea"
	proxyModel.EncryptMethod = "chacha20-ietf-poly1305"
	proxyModel.Password = "rongjin"

	return proxyModel, err
}

// AddProxy ...
func AddProxy(server string, port int, method string, password string) error {
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

// AddOutlineProxy ...
func AddOutlineProxy() (map[string]interface{}, error) {
	nodeService, err := consul.GetRandomProxyService("outline-proxy")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	// with auth api key
	apiKEY := "api"
	outlineAPI := fmt.Sprintf("https://%s:%d/%s/access-key",
		nodeService.Address, nodeService.Port, apiKEY)
	// read from setting
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	body := map[string]interface{}{"username": "testuser", "password": "Diveinedu"}
	resp, err := client.R().SetBody(body).Post(outlineAPI)
	if resp != nil && resp.StatusCode() != 201 {
		return nil, errors.New("post new access-key failed")
	}
	return body, nil
}

// RemoveProxy ...
func RemoveProxy(server string) error {
	var model = &models.ProxyModel{}

	where := map[string]interface{}{"server": server}
	row, err := model.Select(where)
	if err != nil {
		return err
	}

	err = row.Delete()
	return err
}
