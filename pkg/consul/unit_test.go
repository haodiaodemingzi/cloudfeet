package consul

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	a := assert.New(t)
	client := New()
	_, err := client.GetService("proxy-random")
	a.Error(err)
	meta := make(map[string]string)
	meta["name"] = "james"
	meta["outline_id"] = "1"
	meta["method"] = "chacha20"
	meta["access_url"] = "ss://james@chacah20passwd?localhost:3306&chacha20"
	meta["http_check"] = "http://localhost/proxy/check?acceess_url=" + meta["access_url"]
	err = client.RegisterService("proxy-test", "shadowsocks", []string{"ss"}, "ss.csdc.io", 8006, meta)
	a.NoError(err)
	serviceInfo, err := client.GetService("proxy-test")
	a.NoError(err)
	a.Equal(serviceInfo.Port, 8006, "test service port is 8006")
	a.Equal(serviceInfo.Address, "ss.csdc.io", "test service addr is ss.csdc.io")
	a.Equal(serviceInfo.Meta["method"], "chacha20", "test serviceInfo meta method is chacha20")
	a.Equal(serviceInfo.Meta["http_check"], meta["http_check"], "test serviceInfo meta method is chacha20")

	// check service healthy
	_, err = client.GetPassingService("proxy-test")
	a.Error(err)
}
