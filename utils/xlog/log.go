/**
 * Created by goland.
 * @file   log.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/1 19:10
 * @desc   log.go
 */

package xlog

import (
	"github.com/druidcaesa/gotool"
	"github.com/druidcaesa/gotool/openfile"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/x-module/utils/utils/fileutil"
	"github.com/x-module/utils/utils/xerror"
	"os"
	"path"
	"time"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

// Logger 系统日志
var Logger *logrus.Logger

func Debug(args ...interface{}) {
	Logger.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}
func Info(args ...interface{}) {
	Logger.Info(args...)
}
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}
func Error(args ...interface{}) {
	Logger.Error(args...)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return Logger.WithField(key, value)
}

// InitLogger 日志模块初始化
func InitLogger(logFilePath, logFileName string, model string) *logrus.Logger {
	if !fileutil.IsExist(logFilePath) {
		err := os.MkdirAll(logFilePath, os.ModePerm)
		xerror.PanicErr(err, "init system error. make log data err.path:"+logFilePath)
	}
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if !gotool.FileUtils.Exists(fileName) {
		openfile.Create(fileName)
		if !gotool.FileUtils.Exists(fileName) {
			panic("init system error. create log file err. log file:" + fileName)
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	xerror.PanicErr(err, "open log file error")
	// 实例化
	Logger = logrus.New()
	// 设置日志级别
	switch model {
	case ReleaseMode:
		Logger.SetLevel(logrus.WarnLevel)
	case DebugMode:
		Logger.SetLevel(logrus.DebugLevel)
	case TestMode:
		Logger.SetLevel(logrus.InfoLevel)
	}

	// 设置输出
	Logger.Out = src
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	Logger.SetOutput(os.Stdout)

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(30*24*time.Hour),
		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))
	Logger.SetReportCaller(true)
	return Logger
}
