package node_service

import (
	"github.com/haodiaodemingzi/cloudfeet/pkg/consul"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
)

func AddNode(server string, port int, provider string, region string) error{
	log.Info("add new node with sever: %s , port: %d , provider: %s", server, port, provider)
	return consul.RegisterProxyNode(server, port, provider, region)
}
