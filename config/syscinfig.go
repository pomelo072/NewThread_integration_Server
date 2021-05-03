package config

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
)

//配置信息
type sysconfig struct {
	Port        string `json:"Port"`
	DBUserName  string `json:"DBUserName"`
	DBPassword  string `json:"DBPassword"`
	DBIp        string `json:"DBIp"`
	DBPort      string `json:"DBPort"`
	DBName      string `json:"DBName"`
	Admin       string `json:"Admin"`
	AdminVerify string `json:"AdminVerify"`
}

// Sysconfig 生成配置结构
var Sysconfig = &sysconfig{}

// Init 在main前执行
func init() {
	// 读取配置文件
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Printf("error:%s", err)
		panic(err)
	} else {
		// jsoniter解包至Sysconfig
		err = jsoniter.Unmarshal(b, Sysconfig)
		if err != nil {
			fmt.Printf("error:%s", err)
			panic(err)
		}
	}
}
