package outlines

import (
	"fmt"

	"github.com/haodiaodemingzi/cloudfeet/common/settings"
	"github.com/smallnest/goreq"
)

func init() {
	settings.Setup()
}

// AccessKeyInfo ...
type AccessKeyInfo struct {
	ID        string `json:"id"`
	Server    string `json:"server"`
	Port      int    `json:"port"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Method    string `json:"method"`
	AccessURL string `json:"accessUrl"`
}

type Client struct {
	Server string
	Port   string
	APIKey string
	APIURL string
}

// AccessKey ...
type ProxyConfig struct {
	ID        string
	Port      int
	Method    string `json:"method"`
	Password  string
	Name      string
	AccessURL string `json:"accessUrl"`
}

// New ...
func New() *Client {
	server := settings.Config.Outline.Server
	port := settings.Config.Outline.Port
	apiKey := settings.Config.Outline.ApiKey
	apiURL := fmt.Sprintf("https://%s:%s/%s/access-keys/", server, port, apiKey)
	return &Client{Server: server, Port: port, APIKey: apiKey, APIURL: apiURL}
}

// CreateAccessKey ...
func (c *Client) CreateProxy(username string) *ProxyConfig {
	var proxyConfig ProxyConfig
	_, _, err := goreq.New().Post(c.APIURL).ContentType("json").BindBody(&proxyConfig).End()
	if err != nil {
		return nil
	}

	// 修改代理名称
	renameURL := fmt.Sprintf("%s%s/name", c.APIURL, proxyConfig.ID)
	q := fmt.Sprintf("name=%s", username)
	_, _, err = goreq.New().Put(renameURL).SendMapString(q).End()
	if err != nil {
		return nil
	}

	return &proxyConfig
}
