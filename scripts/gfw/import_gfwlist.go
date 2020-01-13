package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"net/http"

	"github.com/haodiaodemingzi/cloudfeet/models"
	log "github.com/haodiaodemingzi/cloudfeet/pkg/logging"
	"github.com/haodiaodemingzi/cloudfeet/pkg/settings"
)

func init() {
	settings.Setup()
	models.Setup()
	log.Setup()
}

func getGfwContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Read body: %v", err)
	}

	return string(data), nil
}

func main() {
	url := `https://cokebar.github.io/gfwlist2dnsmasq/gfwlist_domain.txt`
	content, err := getGfwContent(url)
	if err != nil{
		panic("get gfwlist file failed!!")
	}
	domainList := strings.Split(content, "\n")
	fmt.Println(len(domainList))
	for _, item := range domainList{
		pac := models.PacModel{}
		pac.Domain = item
		pac.Status = 1
		pac.Region = "china"
		pac.FindOrCreate(pac.Domain)
	}
}
