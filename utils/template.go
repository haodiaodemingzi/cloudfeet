package utils

import (
	"bytes"
	"html/template"
	"path/filepath"

	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
)

var boxScriptTpl string = `
#/bin/bash

nvram set ss_enable=1
nvram set ss_mtu=1492
nvram set ss_multiport=22,80,443,1935
nvram set ss_lower_port_only=2
#nvram set ss_usage=" -O auth_sha1_v4 -o http_simple"
nvram set ss_server={{.Server}}
nvram set ss_server_port={{.Port}}
nvram set ss_key={{.Password}}
nvram set ss_method={{.Method}}
nvram set ss_obfs=http_simple
nvram set ss_obfs_param=""
nvram set ss_protocol=auth_sha1_v4
nvram set ss_proto_param=""

logger "更新出国服务器配置..更新出国服务器配置..更新出国服务器配置.."
wget --no-check-certificate  -O  /tmp/gfwlist.conf  {{.GfwURL}}

if [ -f /tmp/gfwlist.conf ];then
	mv /tmp/gfwlist.conf /etc/storage/dnsmasq/dnsmasq.d/gfwlist.dnsmasq.conf
fi
ipset_gfw=ss_spec_dst_fw

restart_dnsproxy
sleep 1
restart_dns
sleep 1

/usr/bin/shadowsocks.sh restart
`

type SserverConfig struct {
	Server     string
	UserName   string
	Password   string
	Port       int
	Method     string
	GfwListURL string
	AuthURL    string
	DomainsUploadURL string
}

func RenderBoxScript(server string, username string,
	password string, port int, method string, authURL string, gfwlistURL string,
	domainsFileURL string) (string, error) {
	// t := template.New("box-script")
	buf := new(bytes.Buffer)
	rootDir := settings.FindRootDir()
	tplFile := filepath.Join(rootDir, "templates", "wifi_start.sh.tpl")
	t, _ := template.ParseFiles(tplFile)
	ssConfig := SserverConfig{
		Server: server, UserName: username, Port: port,
		Password: password, Method: method, GfwListURL: gfwlistURL,
		AuthURL: authURL, DomainsUploadURL: domainsFileURL,
	}
	log.Info("ssconfig : %+v", ssConfig)
	err := t.Execute(buf, ssConfig)
	return buf.String(), err
}
