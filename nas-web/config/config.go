// ---------------------------------
// File Name    : config.go
// Author       : aico
// Mail         : 2237616014@qq.com
// Github       : https://github.com/TBBtianbaoboy
// Site         : https://www.lengyangyu520.cn
// Create Time  : 2021-12-14 13:54:30
// Description  :
// ----------------------------------
package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var IrisConfig Config

type Config struct {
	Web     Web           `yaml:"Web"`
	Mongodb MongodbConfig `yaml:"Mongodb"`
	Redis   RedisConfig   `yaml:"Redis"`
	Kafka   KafkaConfig   `yaml:"Kafka"`
	Log     LogConfig     `yaml:"Log"`
	Other   OtherConfig   `yaml:"Other"`
	RpcUrl  RpcConfig     `yaml:"RpcUrl"`
	Etcd    EtcdConfig    `yaml:"Etcd"`
	Openai  OpenaiConfig  `yaml:"Openai"`
}

type EtcdConfig struct {
	EtcdHost []string `yaml:"EtcdHost"`
}

type RpcConfig struct {
	Forward string `yaml:"Forward"`
}

type MongodbConfig struct {
	Host   string `yaml:"Host"`   //Mongodb 地址
	Port   int    `yaml:"Port"`   //Mongodb 端口
	User   string `yaml:"User"`   //Mongodb 用户名
	Passwd string `yaml:"Passwd"` //Mongodb 密码
	DbName string `yaml:"DbName"` //Mongodb 数据库名
}

type OpenaiConfig struct {
	OpenaiApiKey string `yaml:"OpenaiApiKey"` //OpenaiApiKey
	ProxyUrl     string `yaml:"ProxyUrl"`     //ProxyUrl
}

type RedisConfig struct {
	Host   string `yaml:"Host"`   //Redis 地址
	Port   int    `yaml:"Port"`   //Redis 端口
	Passwd string `yaml:"Passwd"` //Redis 密码
}

type Web struct {
	Host string `yaml:"Host"` //Web 服务地址
}

type LogConfig struct {
	LogPath    string `yaml:"LogPath"`    //日志存放路径
	LogLevel   string `yaml:"LogLevel"`   //日志记录级别
	MaxSize    int    `yaml:"MaxSize"`    //日志分割的尺寸 MB
	MaxAge     int    `yaml:"MaxAge"`     //分割日志保存的时间 day
	Stacktrace string `yaml:"Stacktrace"` //记录堆栈的级别
	IsStdOut   string `yaml:"IsStdOut"`   //是否标准输出console输出
}

type OtherConfig struct {
	IgnoreUrls  []string `yaml:"IgnoreURLs"`  //忽略校验的URL
	JwtTimeOut  int64    `yaml:"JwtTimeOut"`  //jwt超时时间
	JwtLogLevel string   `yaml:"JwtLogLevel"` //jwt日志水平
}

type KafkaConfig struct {
	EndPoints []string `yaml:"EndPoints"` //kafka 地址
}

func ConfigInit(path string) (err error) {
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(conf, &IrisConfig)
	return
}
