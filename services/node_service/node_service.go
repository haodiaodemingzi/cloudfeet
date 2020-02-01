package node_service

import (
	"errors"
	"github.com/haodiaodemingzi/cloudfeet/pkg/consul"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/hashicorp/consul/api"
)

func AddNode(server string, port int, provider string, region string) error {
	log.Info("add new node with sever: %s , port: %d , provider: %s", server, port, provider)
	// make metric service in 9092 port
	if consul.RegisterProxyNode(server, port, provider, region) != nil || consul.RegisterMetricService(server, 9092) != nil {
		log.Error("register service with server = %s port = %d failed", server, port)
		return errors.New("register Node or metric service failed!!!")
	}
	return nil
}

func GetNodeList(serviceName string) ([]api.AgentServiceChecksInfo, error) {
	return consul.GetHealthServices(serviceName)
}
