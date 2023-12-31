package webutils

import (
	"io"
	"mime/multipart"
	"net"
	"os"
	"strings"
	"time"
)

type system struct{}

var System system

// @Func 获取系统唯一表示 - MAC
func (system) GetWebMacAddress() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return time.Now().String()
	}
	return netInterfaces[1].HardwareAddr.String()
}

// @Func 获取系统IP
func (system) GetWebIP() (ip string) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "Unknow"
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// @Func save to file
func (system) SaveFile(file multipart.File, fullPath string) error {
	physical_f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(physical_f, file)
	if err != nil {
		return err
	}
	return nil
}
