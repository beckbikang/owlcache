package log

import (
	"fmt"
	"time"

	config "github.com/beckbikang/owlcache/config"
)

//创建一个全局应用日志
var OwlLogRun *OwlLog

//创建一个全局HTTP日志
var OwlLogHttp *OwlLog

func LogInit() {
	//执行步骤信息
	fmt.Println("owlcache  logger running...")
	//日志目录
	logFilePath := config.OwlConfigModel.Logfile
	//注册全局应用日志
	OwlLogRunRegister(logFilePath)
	//注册全局HTTP日志
	OwlLogHttpRegister(logFilePath)
}

//注册全局应用日志
func OwlLogRunRegister(logFilePath string) {
	//日志文件
	formatLogFileName := "owlcache_run_" + time.Now().Format("20060102_150405") + ".log"
	//创建资源
	OwlLogRun = NewOwlLog(logFilePath, formatLogFileName)
}

//注册全局HTTP日志
func OwlLogHttpRegister(logFilePath string) {
	//日志文件
	formatLogFileName := "owlcache_http_" + time.Now().Format("20060102_150405") + ".log"
	//创建资源
	OwlLogHttp = NewOwlLog(logFilePath, formatLogFileName)
}
