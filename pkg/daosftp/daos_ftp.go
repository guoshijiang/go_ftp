//==================================================================
//创建时间：2017-4-17 首次创建
//功能描述：实现ftp批量上传文件
//创建人：郭世江
//修改记录：若要修改请记录
//==================================================================
package daosftp

import (
	"fmt"
	"net"
	"os"
	"strings"
	ftp "github.com/ftp"
	"io/ioutil"
	"regexp"
	"path/filepath"
	cfg "bjdaos_tool/pkg/daosconfig"
	"bjdaos_tool/pkg/env"
	//"bjdaos_tool/pkg/daoslog"
)

func getListDir(dirPth string) (files []string,files1 []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil,nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			files1 = append(files1, dirPth+PthSep+fi.Name())
			getListDir(dirPth + PthSep + fi.Name())
		}else{
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files,files1, nil
}

func GetAllFileName(path string, str string) (int, []string ) {
	files, _, _ := getListDir(path)
	fileLen := len(files)
	fileSlice := make([]string,0, fileLen)
	//configPath := env.GetConfigPath()
	//ftpConfig := new(cfg.Config)
	//ftpConfig.InitConfig(configPath + "\\config.ini")
	//suffix1 := ftpConfig.Read("file", "file_img")
	//suffix2 := ftpConfig.Read("file", "file_xml")
	reg_front := regexp.MustCompile("\\d{8}")
	reg_end := regexp.MustCompile("\\d{14}")
	if str == ".xml"{
		for i := 0; i < fileLen; i++{
			data_front := reg_front.FindString(files[i])
			date_end := reg_end.FindString(files[i])
			imgName := data_front + "_" + date_end + str
			fileSlice = append(fileSlice, imgName)
		}
	}else if str == ".jpg" {
		for i := 0; i < fileLen; i++{
			data_front := reg_front.FindString(files[i])
			date_end := reg_end.FindString(files[i])
			imgName := data_front + "_" + date_end + str
			fileSlice = append(fileSlice, imgName)
		}
	}
	return fileLen, fileSlice
}

func getLocalIpAddr() string {
	configPath := env.GetConfigPath()
	ftpConfig := new(cfg.Config)
	ftpConfig.InitConfig(configPath + "\\config.ini")
	network := ftpConfig.Read("ftp", "comm_way")
	ip := ftpConfig.Read("ftp", "local_ip")
	port := ftpConfig.Read("ftp", "local_port")
	address := ip + ":" + port
	conn, err := net.Dial(network, address)
	if err != nil {
		return "127.0.0.1"
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

func ftpUploadFile(ftpserver, ftpuser, pw, localFile, remoteSavePath, saveName string) {
	//configPath := env.GetConfigPath()
	//ftpConfig := new(cfg.Config)
	//ftpConfig.InitConfig(configPath + "\\config.ini")
	//ftpfile_path := ftpConfig.Read("ftp", "comm_way")
	ftp, err := ftp.Connect(ftpserver)
	if err != nil {
		fmt.Println(err)
	}
	err = ftp.Login(ftpuser, pw)
	if err != nil {
		fmt.Println(err)
	}
	ftp.ChangeDir("D:\\iTudou")
	dir, err := ftp.CurrentDir()
	ftp.MakeDir(remoteSavePath)
	ftp.ChangeDir(remoteSavePath)
	dir, _ = ftp.CurrentDir()
	fmt.Println(dir)
	file, err := os.Open(localFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	err = ftp.Stor(saveName, file)
	if err != nil {
		fmt.Println(err)
	}
	ftp.Logout()
	ftp.Quit()
	//islog := ftpConfig.Read("file", "log_print")
	//logcotent := fmt.Sprintf("%s:%s","成功上传文件",localFile)
	//daoslog.WriteLog(islog, "System", logcotent)

}

func SendFileToFtpServer() {
	//configPath := env.GetConfigPath()
	//ftpConfig := new(cfg.Config)
	//ftpConfig.InitConfig(configPath + "\\config.yml")
	//xmlPath := ftpConfig.Read("path", "xml_path")
	flen, imgFileName := GetAllFileName("D:\\dian\\out\\wave", ".jpg")
	serverIp := getLocalIpAddr()
	dir := "D:\\dian\\out\\wave"
	//ftpserver := ftpConfig.Read("ftp", "ftp_server_ip")
	//ftpPort := ftpConfig.Read("ftp", "ftp_server_port")
	//ftpuser := ftpConfig.Read("ftp", "ftp_server_name")
	//pw := ftpConfig.Read("ftp", "ftp_server_pwd")
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		for i := 0; i < flen; i++{
			if f.Name() == imgFileName[i] {
				fmt.Println("path=" + path)
				pathFields := strings.Split(path, "\\")
				var domainName string
				if len(pathFields) > 3 {
					domainName = pathFields[len(pathFields)-3]
				}
				ftpUploadFile("192.168.56.1:21", "20123762", "123456", path, domainName, serverIp+"_"+imgFileName[i])
			}
		}
		return nil
	})
}

func RemoveFile(){
	fileLen, fileSlice := GetAllFileName("D:\\dian\\out", ".xml")
	for i := 0; i < fileLen; i++{
		err := os.Remove("D:\\dian\\out\\" + fileSlice[i])
		if err != nil {
			fmt.Println("file remove Error!")
			fmt.Printf("%s", err)
		} else {
			fmt.Print("file remove OK!")
		}
	}
}

func YT()  {
	fmt.Println(getLocalIpAddr())
}
