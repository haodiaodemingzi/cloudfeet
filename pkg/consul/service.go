package consul

import (
	"fmt"
	"log"
	"strings"

	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
	"github.com/hashicorp/consul/api"
)

// ConsulConn client
var ConsulConn *api.Client

// Setup comment
func Setup() {
	apiConfig := api.DefaultConfig()
	apiConfig.Datacenter = settings.Config.Consul.DC
	apiConfig.Address = settings.Config.Consul.Addr
	apiConfig.Scheme = settings.Config.Consul.Scheme
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	ConsulConn = client
}

// GetService ...
func GetService(serviceID string) (*api.AgentService, error) {
	service, _, err := ConsulConn.agent.Service(serviceID, &api.QueryOptions{})
	return service, err
}

// GetServices ...
func GetServices(map[string]*api.AgentService, error) {
	services, err := ConsulConn.agent.Services()
	if err != nil {
		log.Fatal("Get service list error :" + err.Error())
	}
	return services, err

}

// RegisterService ...
// only support http check
func RegisterService(
	serviceID string, serviceName string, tags []string,
	addr string, port int, meta map[string]string) error {
	var serviceInfo = api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Meta:    meta,
		Address: addr,
		Port:    port,
		Tags:    tags}
	err := ConsulConn.agent.ServiceRegister(&serviceInfo)
	return err
}

// DeRegisterService ...
func DeRegisterService(serviceID string) error {
	return ConsulConn.agent.ServiceDeregister(serviceID)
}

// GetPassingService ...
func GetPassingService(serviceID string) (*api.AgentService, error) {
	serviceInfo, err := GetService(serviceID)
	if err != nil {
		return serviceInfo, err
	}
	passingOnly := true
	addr, _, err := ConsulConn.health.Service(serviceID, strings.Join(serviceInfo.Tags, ""), passingOnly, nil)
	if len(addr) == 0 && err == nil {
		err = fmt.Errorf("service = %s not found", serviceID)
		log.Println(err)
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return serviceInfo, err
}
