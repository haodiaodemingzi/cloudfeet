package settings

import (
	"path"
	"runtime"
	"time"

	"github.com/haodiaodemingzi/cloudfeet/common/logging"
	"github.com/spf13/viper"
)

// Config struct
type ConfModel struct {
	PrefixUrl       string `json:"prefix_url"`
	PageSize        int    `json:"page_size"`
	JwtSecret       string `json:"jwt_secret"`
	RuntimeRootPath string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExts  []string
	LogSavePath     string
	LogSaveName     string
	Root            Root    `json:"root"`
	CORS            CORS    `json:"cors"`
	MySQL           MySQL   `json:"mysql"`
	Redis           Redis   `json:"redis"`
	Outline         Outline `json:"outline"`
	Log             Log     `json:"log"`
	Gin             Gin     `Gin:"gin"`
	Jwt             Jwt     `Jwt:"jwt"`
}

// Log config
type Log struct {
	Level  string `json:"level"`
	Format string `json:"format"`
	Path   string `json:"path"`
}

type Outline struct {
	Server string `json:"server"`
	Port   string `json:"port"`
	ApiKey string `json:"api_key"`
}

// Root config
type Root struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	RealName string `json:"real_name"`
}

// CORS config
type CORS struct {
	Enable           bool     `json:"enable"`
	AllowOrigins     []string `json:"allow_origins"`
	AllowMethods     []string `json:"allow_methods"`
	AllowHeaders     []string `json:"allow_headers"`
	AllowCredentials bool     `json:"allow_credentials"`
	MaxAge           int      `json:"max_age"`
}

// MySQL config
type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DataBase string `json:"database"`
}

// Gin config
type Gin struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	RunMode string `json:"run_mode"`
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
	Secret     string `json:"secret"`
	ExpireHour int64  `json:"expire_hour"`
}

var RedisConfig = &Redis{}
var Viper = viper.New()
var Config ConfModel

// FindRootDir find root dir for project
func FindRootDir() string {
	_, filename, _, _ := runtime.Caller(0)
	abspath := path.Join(path.Dir(filename), "../..")
	logging.Info("日志路径: %s", abspath)
	return abspath
}

// Setup init all config
func Setup() {
	configPath := FindRootDir()
	Viper.SetConfigName("app")
	Viper.AddConfigPath(configPath)
	err := Viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	conf := ConfModel{}
	Viper.Unmarshal(&conf)
	Config = conf
}
