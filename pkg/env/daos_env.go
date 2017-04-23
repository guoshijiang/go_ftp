//==================================================================
//创建时间：2017-4-20 首次创建
//功能描述：windows和Linux下路径兼容处理
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package env
import (
	"os"
	"runtime"
)

var ostype = runtime.GOOS

func GetProjectPath() string{
	var projectPath string
	projectPath, _ = os.Getwd()
	return projectPath
}

func GetConfigPath() string{
	path := GetProjectPath()
	if ostype == "windows"{
		path = path + "\\" + "config\\"
	}else if ostype == "linux"{
		path = path +"/" + "config/"
	}
	return  path
}

func GetConLogPath() string{
	path := GetProjectPath()
	if ostype == "windows"{
		path = path + "\\log\\"
	}else if ostype == "linux"{
		path = path + "/log/"
	}
	return  path
}