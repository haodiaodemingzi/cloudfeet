package settings

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

// Config struct
type ConfModel struct {
	MySQL MySQL
	Jwt   Jwt
	Log   Log
	Gin   Gin
	Debug bool
}

// Log config
type Log struct {
	Level  string
	Format string
	Path   string
}

type Outline struct {
	Server string
	Port   string
	ApiKey string `mapstructure:"api_key"`
}

// Root config
type Root struct {
	UserName string `mapstructure:"user_name"`
	Password string
	RealName string `mapstructure:"real_name"`
}

// CORS config
type CORS struct {
	Enable           bool
	AllowOrigins     []string `mapstructure:"allow_origins"`
	AllowMethods     []string `mapstructure:"allow_methods"`
	AllowHeaders     []string `mapstructure:"allow_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
}

// MySQL config
type MySQL struct {
	Host     string
	Port     int
	User     string
	Password string
	DataBase string
	Debug    bool
}

// Gin config
type Gin struct {
	Host    string
	Port    int
	RunMode string
	BaseURL string `mapstructure:"base_url"`
}

// Redis config
type Redis struct {
	Host        string
	Port        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// Jwt config
type Jwt struct {
	Secret     string
	ExpireHour int64 `mapstructure:"expire_hour"`
}

var Viper = viper.New()
var Config ConfModel

// FindRootDir find root dir for project
func FindRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	abspath := path.Join(path.Dir(filename), "../..")
	return abspath
}

// Setup init all config
func Setup() {
	root := FindRootDir()
	configPath := filepath.Join(root, "conf")

	fmt.Println("log path: ", configPath)

	Viper.SetConfigName("app")
	Viper.AddConfigPath(configPath)

	err := Viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	conf := ConfModel{}
	err = Viper.Unmarshal(&conf)
	if err != nil {
		panic(err.Error())
	}
	Config = conf
}
