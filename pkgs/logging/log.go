package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type Log struct {
	Logger *logrus.Logger
}

func GetLogger() *Log {
	log := &Log{
		Logger: &logrus.Logger{
			Out:   os.Stdout,
			Level: logrus.DebugLevel,
			Formatter: &easy.Formatter{
				TimestampFormat: "2006-01-02 15:04:05",
				LogFormat:       "[%lvl%]: %time% - %msg%\n",
			},
		},
	}
	return log
}

func (log *Log) Info(s string, v ...interface{}) {
	log.Logger.Infof(s, v...)
}

func (log *Log) Debug(s string, v ...interface{}) {
	log.Logger.Debugf(s, v...)
}

func (log *Log) Warn(s string, v ...interface{}) {
	log.Logger.Warnf(s, v...)
}

func (log *Log) Fatal(s string, v ...interface{}) {
	log.Logger.Fatalf(s, v...)
}

func (log *Log) Error(s string, v ...interface{}) {
	log.Logger.Errorf(s, v...)
}

func (log *Log) Panic(s string, v ...interface{}) {
	log.Logger.Panicf(s, v...)
}
