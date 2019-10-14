package main

import (
	"net/http"

	"github.com/haodiaodemingzi/cloudfeet/models"
	log "github.com/haodiaodemingzi/cloudfeet/pkgs/logging"
	"github.com/haodiaodemingzi/cloudfeet/pkgs/settings"
	"github.com/haodiaodemingzi/cloudfeet/utils"
)

func init(){
	settings.Setup()
	models.Setup()
	log.Setup()
}


func main() {
	userRules := []string{"||4ter2n.com", "|https://85.17.73.31/"}
	client := &http.Client{}

	gfw, err := utils.NewGFWList("https://raw.githubusercontent." +
		"com/gfwlist/gfwlist/master/gfwlist.txt",
		client, userRules, "gfwlist.txt", false)
	if err != nil{
		panic(err)
	}
	log.Debug("test map %+v", gfw.RuleMap)
	for k, _ := range gfw.RuleMap{
		pac := models.PacModel{}
		pac.Domain = k
		pac.Status = 1
		pac.Region = "china"
		pac.FindOrCreate(pac.Domain)

	}
}
