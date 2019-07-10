package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var Logger *logrus.Logger

func init() {
	logger := &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}
	Logger = logger
}

func GetLogger() *logrus.Logger {
	return Logger
}

func Info(s string, v ...interface{}) {
	Logger.Infof(s, v...)
}

func Debug(s string, v ...interface{}) {
	Logger.Debugf(s, v...)
}

func Warn(s string, v ...interface{}) {
	Logger.Warnf(s, v...)
}

func Fatal(s string, v ...interface{}) {
	Logger.Fatalf(s, v...)
}

func Error(s string, v ...interface{}) {
	Logger.Errorf(s, v...)
}

func Panic(s string, v ...interface{}) {
	Logger.Panicf(s, v...)
}
