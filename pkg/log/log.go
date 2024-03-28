package log

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
)

type Fields logrus.Fields

type LogInfo struct {
	isMultiFile bool
	logHandler  map[string]*logrus.Logger
	logLevel    logrus.Level
	logPath     string
}

var (
	logInfo  LogInfo
	logLevel []string
)

func Init(logPath string) {
	logInfo = LogInfo{
		isMultiFile: true,
		logHandler:  make(map[string]*logrus.Logger),
		logLevel:    logrus.DebugLevel,
		logPath:     logPath,
	}
	loglevel := []string{"debug", "info", "warn", "error", "fatal", "access"}

	for _, level := range loglevel {
		logInfo.logHandler[level] = initHandler(level, logPath)
	}
}

func initHandler(level string, logPath string) *logrus.Logger {
	_, err := os.Stat(logPath)

	if os.IsNotExist(err) {
		err = os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("create log path failed:%s", err.Error()))
		}
	}

	file, err := os.OpenFile(logPath+level+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(fmt.Sprintf("open log path failed:%s", err.Error()))
	}

	logHandler := logrus.New()
	logHandler.SetLevel(logInfo.logLevel)
	logHandler.SetOutput(file)
	logHandler.SetFormatter(&Formatter{})

	return logHandler
}

type Formatter struct {
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	time := entry.Time.Format("2006-01-02 15:04:05")

	ip := ""

	if entry.Data["ip"] != nil {
		ip = entry.Data["ip"].(string)
		delete(entry.Data, "ip")
	}

	uniqId := ""

	if entry.Data["uniqId"] != nil {
		uniqId = entry.Data["uniqId"].(string)
		delete(entry.Data, "uniqId")
	}

	file := getCaller()

	msg := entry.Message

	dataByte, _ := json.Marshal(entry.Data)

	logMsg := fmt.Sprintf("[%s] [%s] [%s] [%s] [%s] [%s]\n", time, ip, uniqId, file, msg, string(dataByte))

	return []byte(logMsg), nil
}

func getCaller() string {
	_, file, line, _ := runtime.Caller(7)
	return file + ":" + strconv.Itoa(line)
}

func getHandler(level string) *logrus.Logger {
	var logHandler *logrus.Logger

	if logInfo.isMultiFile {
		logHandler, _ = logInfo.logHandler[level]
	} else {
		logHandler, _ = logInfo.logHandler["info"]
	}

	return logHandler
}

//调试
func Debug(ctx context.Context, msg string, data Fields) {
	data["uniqId"] = ctx.Value("uuid")
	data["ip"] = ctx.Value("ip")
	getHandler("debug").WithFields(logrus.Fields(data)).Debug(msg)
}

//正常
func Info(ctx context.Context, msg string, data Fields) {
	data["uniqId"] = ctx.Value("uuid")
	data["ip"] = ctx.Value("ip")
	getHandler("info").WithFields(logrus.Fields(data)).Info(msg)
}

//警告
func Warn(ctx context.Context, msg string, data Fields) {
	data["uniqId"] = ctx.Value("uuid")
	data["ip"] = ctx.Value("ip")
	getHandler("warn").WithFields(logrus.Fields(data)).Warn(msg)
}

//错误
func Error(ctx context.Context, msg string, data Fields) {
	data["uniqId"] = ctx.Value("uuid")
	data["ip"] = ctx.Value("ip")
	getHandler("error").WithFields(logrus.Fields(data)).Error(msg)
}

//严重
func Fatal(ctx context.Context, msg string, data Fields) {
	data["uniqId"] = ctx.Value("uuid")
	data["ip"] = ctx.Value("ip")
	getHandler("fatal").WithFields(logrus.Fields(data)).Fatal(msg)
}
