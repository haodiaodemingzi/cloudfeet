package main

import (
	"log"

	"github.com/haodiaodemingzi/cloudfeet/pkg/consul"
	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
)

func init(){
	settings.Setup()
	consul.Setup()
}



func main() {
	_ = consul.RegisterProxyNode("prom2.switfin.org", 9000, "gcp", "usa")
	_ = consul.RegisterProxyNode("prom.switfin.org", 9000, "gcp", "usa")
	service, _ := consul.GetService("outline-prom.switfin.org")
	log.Printf("get prom service -> %+v", service)
	service2, _ := consul.GetService("outline-prom2.switfin.org")
	log.Printf("get prom2 service -> %+v", service2)

	healthService, _ := consul.GetRandomProxyService("outline-proxy")
	log.Printf("get random health service -> %+v", healthService)
}


