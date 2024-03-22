package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
	//七牛
	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	LoadServer(cfg)
	LoadData(cfg)
	LoadQiniu(cfg)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("eternal")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("ljy123...")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("fXBIBlHwHbcItG4B5XGTkeub4uIfUdxJeFEUl-6t")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("v0aXVRCNId1lBoCmkEfjgzd6qzBpdbLBztuVrL49")
	Bucket = file.Section("qiniu").Key("Bucket").MustString("binlog")
	QiniuSever = file.Section("qiniu").Key("QiniuSever").MustString("http://rkllm6psp.hn-bkt.clouddn.com/")

}
