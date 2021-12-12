package main

import (
	"github.com/go_example/mylogger"
)

// 定义全局的logger接口变量
var log mylogger.Logger

func main() {
	//log = mylogger.NewConsoleLog("ERROR")
	log = mylogger.NewFileLogger("info", "./", "test.log", 100*1024*1024)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warn("这是一条Warn日志")
		log.Error("这是一条Error日志, id:%v", 1010011314124)
		log.Fatal("这是一条Fatal日志")
	}
}
