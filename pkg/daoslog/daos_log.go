//==================================================================
//创建时间：2017-4-19 首次创建
//功能描述：打日志，日志业务开关，日志级别
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package daoslog

import (
	"fmt"
	"log"
	"os"
	"github.com/golang/glog"
	"time"
	"bjdaos_tool/pkg/env"
)

func WriteLog(islog,  logtype , errcontent string) {
	if islog == "false" {
		return
	}
	if logtype != "Info" && logtype!= "Debug" && logtype!= "Error" && logtype != "System"  {
		glog.Error("this is not a logtype ")
		return
	}
	data := time.Now().Format("20060102")
	logPath := env.GetConLogPath()
	logFilea := logPath + "\\" + data+"_"+ logtype+".log"
	errcontent = "[" +errcontent + "]"

	logFile, err := os.OpenFile(logFilea, os.O_RDWR | os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("open file error=%s\r\n", err.Error())
		os.Exit(-1)
	}
	logger := log.New(logFile, "{"+logtype+"} ", log.Ldate | log.Ltime | log.Lshortfile)
	logger.Println(errcontent)
}

