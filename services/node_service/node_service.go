package node_service

import (
	"github.com/haodiaodemingzi/cloudfeet/pkg/consul"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/hashicorp/consul/api"
)

func AddNode(server string, port int, provider string, region string) error {
	log.Info("add new node with sever: %s , port: %d , provider: %s", server, port, provider)
	return consul.RegisterProxyNode(server, port, provider, region)
}

func GetNodeList(serviceName string) ([]api.AgentServiceChecksInfo, error) {
	return consul.GetHealthServices(serviceName)
}
