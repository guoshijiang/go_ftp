package main

import (
	"bjdaos_tool/pkg/daosftp"
	//cfg "bjdaos_tool/pkg/daosconfig"
	//"bjdaos_tool/pkg/env"
)

func main(){
	//configPath := env.GetConfigPath()
	//ftpConfig := new(cfg.Config)
	//ftpConfig.InitConfig(configPath + "\\config.ini")
	//path := ftpConfig.Read("path", "xml_path")
	//xml := ftpConfig.Read("file", "file_xml")
	//for{
	//	_, slice := daosftp.GetAllFileName("D:\\dian\\out", ".xml")
	//	fmt.Println(slice)
	//	daosftp.SendFileToFtpServer()
	//	daosftp.RemoveFile()
	//}
	daosftp.YT()
}


//func main() {
//	a := env.GetConfigPath()
//	myConfig := new(cfg.Config)
//	myConfig.InitConfig(a + "\\config.yml")
//	b := myConfig.Read("file", "file_img")
//	fmt.Println(b)
//	str := fmt.Sprintf("%s:%s","ss","aaa" )
//	fmt.Println(str)
//}








