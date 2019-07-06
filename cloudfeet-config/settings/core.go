package settings

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)
var Cfg = viper.New()

// FindRootDir find root dir for project
func FindRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	abspath := path.Join(path.Dir(filename), "..")
	return abspath
}

// Setup init all config
func Setup() {
	configPath := path.Join(FindRootDir(), "conf")
	fmt.Println(configPath)

	Cfg.SetConfigName("app")
	Cfg.AddConfigPath(configPath)
	err := Cfg.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
