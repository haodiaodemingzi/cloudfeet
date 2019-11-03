package settings

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

// ConfModel ...
type ConfModel struct {
	MySQL  MySQL
	Jwt    Jwt
	Log    Log
	Gin    Gin
	Debug  bool
	URL    URL
	Consul Consul
}

// Log config
type Log struct {
	Level  string
	Format string
	Path   string
}

// Outline ...
type Outline struct {
	Server string
	Port   string
	APIKEY string `mapstructure:"api_key"`
}

// Root ...
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

// URL ...
type URL struct {
	AuthToken     string `mapstructure:"auth_token"`
	PullDomains   string `mapstructure:"pull_domains"`
	UpdateDomains string `mapstructure:"update_domains"`
	UploadDomains string `mapstructure:"upload_domains"`
	UploadDNSFile string `mapstructure:"upload_dns_file"`
	PacConfig     string `mapstructure:"pac_config"`
	InitScript    string `mapstructure:"init_script"`
	ProxyInfo     string `mapstructure:"proxy_info"`
}

// Jwt config
type Jwt struct {
	Secret     string
	ExpireHour int64 `mapstructure:"expire_hour"`
}

// Consul model
type Consul struct {
	Scheme string `mapstructure:"scheme"`
	DC     string `mapstructure:"dc"`
	Addr   string `mapstructure:"addr"`
}

// Viper config
var Viper = viper.New()

// Config model
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
