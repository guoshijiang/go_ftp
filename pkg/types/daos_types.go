package types

type ConfigItem struct {
	xmlPath         string `json:"xml_path"`                     //xml文件所在的路径
	imgPath         string  `json:"img_path"`                    //图片文件所在路径
	ftpfilePath     string `json:"ftpfile_path"`             //在服务器上的文件存储路径
	ftpServerIp     string `json:"ftp_server_ip"`          //ftp服务器的IP
	ftpServerPort   string `json:"ftp_server_port"`      //ftp服务器的端口
	ftpServerName   string `json:"ftp_server_name"`       //ftp服务器的用户名
	ftpServerPwd    string   `json:"ftp_server_pwd"`       //ftp服务器的密码
	localIp         string `json:"local_ip"`                    //本地主机的IP
	fileImg         string `json:"file_img"`                   //传输文件的格式
	fileBmp         string `json:"file_bmp"`                   //传输文件的格式
	fileXml         string `json:"file_xml"`                   //传输文件的格式
	logPrint        string `json:"log_print"`                 //是否打日志，生产环境上不打日志，在调式程序的时候使用
}