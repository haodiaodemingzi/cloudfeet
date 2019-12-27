package agent

import (
	"fmt"
	"strconv"

	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
	"github.com/smallnest/goreq"
)

func init() {
	settings.Setup()
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
func NewOutline() *Client {
	server := settings.Config.Outline.Server
	port := settings.Config.Outline.Port
	apiKey := settings.Config.Outline.ApiKey
	apiURL := fmt.Sprintf("https://%s:%s/%s/access-keys/", server, port, apiKey)
	return &Client{Server: server, Port: port, APIKey: apiKey, APIURL: apiURL}
}

// Add a proxy
func (c *Client) CreateProxy(username string) (*ProxyConfig, error) {
	var proxyConfig ProxyConfig
	_, _, err := goreq.New().Post(c.APIURL).ContentType("json").BindBody(&proxyConfig).End()
	if err != nil {
		return nil, nil
	}

	// 修改代理名称
	renameURL := fmt.Sprintf("%s%s/name", c.APIURL, proxyConfig.ID)
	q := fmt.Sprintf("name=%s", username)
	_, _, err = goreq.New().Put(renameURL).SendMapString(q).End()
	if err != nil {
		return nil, nil
	}

	return &proxyConfig, nil
}

// Delete proxy
func (c *Client) DeleteProxy(id int) bool {
	url := c.APIURL + strconv.Itoa(id)
	_, _, err := goreq.New().Post(url).End()
	if err != nil {
		return false
	}
	return true
}
