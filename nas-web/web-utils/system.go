//---------------------------------
//File Name    : web-utils/system.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-27 13:45:08
//Description  :
//----------------------------------
package webutils

import (
	"net"
	"strings"
	"time"
)

type system struct{}

var System system

//@Func 获取系统唯一表示 - MAC
func (system) GetWebMacAddress() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return time.Now().String()
	}
	return netInterfaces[1].HardwareAddr.String()
}

//@Func 获取系统IP
func (system) GetWebIP() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "Unknow"
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
