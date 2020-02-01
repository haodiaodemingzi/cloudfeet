package consul

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
	"github.com/hashicorp/consul/api"
)

// ConsulConn client
var ConsulConn *api.Client

// Setup comment
func Setup() {
	client, err := api.NewClient(&api.Config{
		Address:    settings.Config.Consul.Addr,
		Scheme:     "http",
		Datacenter: settings.Config.Consul.DC})
	if err != nil {
		panic(err)
	}
	ConsulConn = client
}

// GetService ...
func GetService(serviceID string) (*api.AgentService, error) {
	service, _, err := ConsulConn.Agent().Service(serviceID, nil)
	return service, err
}

// GetServices ...
func GetServices(map[string]*api.AgentService, error) (map[string]*api.AgentService, error) {
	services, err := ConsulConn.Agent().Services()
	if err != nil {
		log.Fatal("Get service list error :" + err.Error())
	}
	return services, err
}

// TODO: set constant in module, add prometheus service register
func RegisterProxyNode(server string, port int, provider string, region string) error {
	serviceId := "outline-" + server
	serviceName := "outline-proxy"
	reg := makeServiceReg(serviceId, serviceName, []string{"proxy", "v1"}, server, port)
	metaInfo := map[string]string{
		"region": region, "provider": provider,
	}
	reg.Meta = metaInfo
	reg.Check = &api.AgentServiceCheck{
		Interval:                       "120s",
		Timeout:                        "10s",
		TCP:                            fmt.Sprintf("%s:%d", server, port),
		DeregisterCriticalServiceAfter: "15m",
	}

	// make service
	err := ConsulConn.Agent().ServiceRegister(&reg)
	if err != nil {
		panic(err)
	}
	return err
}

func RegisterMetricService(server string, port int) error {
	serviceId := "metric-" + server
	serviceName := "ss-metrics"
	reg := makeServiceReg(serviceId, serviceName, []string{"metrics"}, server, port)
	reg.Check = &api.AgentServiceCheck{
		Interval:                       "120s",
		Timeout:                        "10s",
		HTTP:                           fmt.Sprintf("http://%s:%d/metrics", server, port),
		DeregisterCriticalServiceAfter: "15m",
	}

	// make service
	err := ConsulConn.Agent().ServiceRegister(&reg)
	if err != nil {
		panic(err)
	}
	return err
}

func makeServiceCheckReg(serviceName string, serviceID string) {
}

func makeServiceReg(
	serviceID string, serviceName string, tags []string,
	addr string, port int) api.AgentServiceRegistration {
	return api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Address: addr,
		Port:    port,
		Tags:    tags}
}

// DeRegisterService ...
func DeRegisterService(serviceID string) error {
	return ConsulConn.Agent().ServiceDeregister(serviceID)
}

func GetHealthServices(serviceName string) ([]api.AgentServiceChecksInfo, error) {
	_, serviceList, err := ConsulConn.Agent().AgentHealthServiceByName(serviceName)
	return serviceList, err
}

func GetRandomProxyService(serviceName string) (*api.AgentService, error) {
	serviceList, err := GetHealthServices(serviceName)
	if err != nil || len(serviceList) == 0 {
		return &api.AgentService{}, err
	}

	random := serviceList[rand.Intn(len(serviceList))]
	return random.Service, nil
}
