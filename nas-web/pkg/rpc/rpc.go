//---------------------------------
//File Name    : rpc.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2022-01-10 21:44:01
//Description  : 
//----------------------------------
package rpc

import (
	"fmt"

	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type RpcClientConf struct {
	Target    string
	EtcdHosts []string
	EtcdKey   string
	App       string
	Token     string
}

func NewClient(cfg RpcClientConf, opts ...grpc.DialOption) (zrpc.Client, error) {
	mergeOpts := make([]grpc.DialOption, 0)
	mergeOpts = append(mergeOpts, newDefaultClientOption()...)
	mergeOpts = append(mergeOpts, opts...)

	formatOpts := make([]zrpc.ClientOption, 0)
	for _, opt := range mergeOpts {
		formatOpts = append(formatOpts, zrpc.WithDialOption(opt))
	}

	endpoints := make([]string, 0)
	if cfg.Target != "" {
		endpoints = append(endpoints, cfg.Target)
	}

	client, err := zrpc.NewClient(
		zrpc.RpcClientConf{
			Endpoints: endpoints,
			Etcd: discov.EtcdConf{
				Hosts: cfg.EtcdHosts,
				Key:   fmt.Sprintf("/lengyangyu520.cn/nas/rpc/%s", cfg.EtcdKey),
			},
			App:   cfg.App,
			Token: cfg.Token,
		},
		formatOpts...,
	)
	if err != nil {
		return nil, fmt.Errorf("rpc NewClient error: %s", err.Error())
	}

	return client, nil
}

func newDefaultClientOption() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(MaxPackageSize),
			grpc.MaxCallSendMsgSize(MaxPackageSize),
		),
	}
}
