package main

import (
	"bjdaos_tool/pkg/daosftp"
	"bjdaos_tool/pkg/env"
	cfg "bjdaos_tool/pkg/daosconfig"
)

func main(){
	configPath := env.GetConfigPath()
	ftpConfig := new(cfg.Config)
	ftpConfig.InitConfig(configPath + "\\config.ini")
	xml_path := ftpConfig.Read("path", "xml_path")
	img_path := ftpConfig.Read("path", "img_path")
	file_img := ftpConfig.Read("file", "file_img")
	file_xml := ftpConfig.Read("file", "file_xml")
	for{
		daosftp.SendXmlFileToFtpServer(xml_path, file_xml)
		daosftp.SendJpgFileToFtpServer(img_path, file_img)
	}
}
