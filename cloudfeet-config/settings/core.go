package settings

import (
	"fmt"
	"path"
	"runtime"

	"github.com/spf13/viper"
)
var Cfg = &Config{}

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

	viper.SetConfigName("app")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	viper.Unmarshal(Cfg)
}

