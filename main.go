// test server入口文件
package main

import (
	"net"
	"sync"

	"util/logs"
	"util/run"
)

// 程序入口
func main() {
	logs.Info("\n\n\ntest start...")
	defer logs.Info("test end!")

	defer run.Recover(true)

	//
	ok := LoadConfig("conf/")
	if !ok {
		logs.Error("load config file failed!")
		return
	}

	//
	addr := Cfg.GateAddrs[0]
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr)

	num := Cfg.ClientNum
	logs.Infoln("max conn:", num)

	w := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()

			conn, e := net.DialTCP("tcp", nil, tcpAddr)
			if e != nil {
				logs.Warnln(i, e)
				return
			}
			defer conn.Close()

			client := &Client{conn, i}
			TestClient(client)
		}(i)
	}

	w.Wait()
}
