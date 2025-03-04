package log

import (
	"io"
	"m-server-api/config"
	"m-server-api/utils/file"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *logrus.Logger

func InitLog() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pwd = pwd + "/logs"
	if !file.FileExists(pwd) {
		err = file.CreateDirectory(pwd)
	}
	if err != nil {
		panic(err)
	}
	if log == nil {
		log = logrus.New()
	}
	logCfg := config.Get().Log
	switch logCfg.Level {
	case "":
		log.SetLevel(logrus.ErrorLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		panic("log level is not support: " + logCfg.Level)
	}
	fileLog := &lumberjack.Logger{
		Filename:   pwd + "/" + config.PROJECT + ".log",
		MaxSize:    cast.ToInt(10), // 日志文件最大 size, 单位是 MB
		MaxBackups: 0,              // 最大过期日志保留的个数 0 表示不删除
		MaxAge:     0,              //保留过期文件的最大时间间隔,单位是天 0 表示不删除
		Compress:   true,           // disabled by default,是否需要压缩滚动日志, 使用的 gzip 压缩
	}
	var writers []io.Writer
	writers = append(writers, os.Stdout, fileLog)
	multiWriter := io.MultiWriter(writers...)
	log.SetOutput(multiWriter)
}

func Error(args ...interface{}) {
	Get().Error(args...)
}

func ErrorTraceId(traceId any, args ...interface{}) {
	Get().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Error(args...)
}
func Fatal(args ...interface{}) {
	Get().Fatal(args...)
}
func FatalTraceId(traceId any, args ...interface{}) {
	Get().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Fatal(args...)
}
func Warn(args ...interface{}) {
	Get().Warn(args...)
}
func WarnTraceId(traceId any, args ...interface{}) {
	Get().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Warn(args...)
}
func Info(args ...interface{}) {
	Get().Info(args...)
}
func InfoTraceId(traceId any, args ...interface{}) {
	Get().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Info(args...)
}
func Debug(args ...interface{}) {
	Get().Debug(args...)
}
func DebugTraceId(traceId any, args ...interface{}) {
	Get().WithFields(logrus.Fields{
		"traceId": traceId,
	}).Debug(args...)
}
func Errorf(format string, args ...interface{}) {
	Get().Errorf(format, args...)
}
func Warnf(format string, args ...interface{}) {
	Get().Warnf(format, args...)
}
func Infof(format string, args ...interface{}) {
	Get().Infof(format, args...)
}

func Debugf(format string, args ...interface{}) {
	Get().Debugf(format, args...)
}

func Get() *logrus.Logger {
	if log == nil {
		InitLog()
	}
	return log
}
