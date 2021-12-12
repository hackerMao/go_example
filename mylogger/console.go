package mylogger

import (
	"fmt"
	"time"
)

/**
往终端输出日志
*/

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog 日志构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(level LogLevel) bool {
	return level >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, arg ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		timeFmt := time.Now().Format("2006-01-02 15:04:05")
		funcName, filename, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", timeFmt, getLevelString(lv), filename, funcName, lineNo, msg)
	}

}

func (c ConsoleLogger) Debug(format string, arg ...interface{}) {
	c.log(DEBUG, format, arg...)
}

func (c ConsoleLogger) Trace(format string, arg ...interface{}) {
	c.log(TRACE, format, arg...)
}

func (c ConsoleLogger) Info(format string, arg ...interface{}) {
	c.log(INFO, format, arg...)
}

func (c ConsoleLogger) Warn(format string, arg ...interface{}) {
	c.log(WARN, format, arg...)
}

func (c ConsoleLogger) Error(format string, arg ...interface{}) {
	c.log(ERROR, format, arg...)
}

func (c ConsoleLogger) Fatal(format string, arg ...interface{}) {
	c.log(FATAL, format, arg...)
}
