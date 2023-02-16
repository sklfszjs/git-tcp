package utils

import (
	"encoding/json"
	"go-tcp/ginterface"
	"io/ioutil"
)

// 全局参数配置，通过json文件
type GlobalObj struct {
	//server
	TcpServer ginterface.IServer
	Host      string
	TcpPort   int
	Name      string
	//go-tcp框架
	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("config/go-tcp.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:           "go-tcp server",
		Version:        "V0.4",
		TcpPort:        8888,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	GlobalObject.Reload()

}
