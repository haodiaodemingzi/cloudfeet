package logging

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/haodiaodemingzi/cloudfeet/pkgs/settings"
)

var logger logrus.Logger

func Setup() {
	levelMaps := map[string]logrus.Level{
		"debug": logrus.DebugLevel, "info": logrus.InfoLevel, "warn": logrus.WarnLevel,
		"error": logrus.ErrorLevel, "panic": logrus.PanicLevel,
	}
	var level logrus.Level
	level, ok := levelMaps[settings.Config.Log.Level]
	if !ok {
		level = logrus.DebugLevel
	}

	logfile := settings.Config.Log.Path
	file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		panic(err)
	}

	logger = logrus.Logger{
		Out:          file,
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: false,
		Level:        level,
	}
	logger.SetOutput(os.Stdout)
}

func Test(msg string){
	logger.Debug(msg)
}

func Info(s string, v ...interface{}) {
	logger.Infof(s, v...)
}

func Debug(s string, v ...interface{}) {
	logger.Debugf(s, v...)
}

func Warn(s string, v ...interface{}) {
	logger.Warnf(s, v...)
}

func Fatal(s string, v ...interface{}) {
	logger.Fatalf(s, v...)
}

func Error(s string, v ...interface{}) {
	logger.Errorf(s, v...)
}

func Panic(s string, v ...interface{}) {
	logger.Panicf(s, v...)
}

