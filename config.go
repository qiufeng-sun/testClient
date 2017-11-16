package main

import (
	"path/filepath"

	"github.com/astaxie/beego/config"

	"util/logs"
)

var _ = logs.Debug

//
type Config struct {
	Name      string
	ClientNum int

	// server
	GateAddrs []string
}

func (this *Config) init(fileName string) bool {
	confd, e := config.NewConfig("ini", fileName)
	if e != nil {
		logs.Panicln("load config file failed! file:", fileName, "error:", e)
	}

	// [client]
	this.Name = confd.String("client::name")
	this.ClientNum = confd.DefaultInt("client::num", 5)

	// [gate]
	this.GateAddrs = confd.Strings("gate::addrs")

	// echo
	logs.Info("gate config:%+v", *this)

	return true
}

//
var Cfg = &Config{}

//
func LoadConfig(confPath string) bool {
	// config
	confFile := filepath.Clean(confPath + "/self.ini")

	return Cfg.init(confFile)
}

// to do add check func
