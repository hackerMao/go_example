package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

/**
自定义一个日志库
*/

// LogLevel 日志等级
type LogLevel int16

// Logger 接口
type Logger interface {
	Debug(format string, arg ...interface{})

	Trace(format string, arg ...interface{})

	Info(format string, arg ...interface{})

	Warn(format string, arg ...interface{})

	Error(format string, arg ...interface{})

	Fatal(format string, arg ...interface{})
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLevelString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(n int) (funcName, filename string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	filename = path.Base(file)
	return funcName, filename, lineNo
}
