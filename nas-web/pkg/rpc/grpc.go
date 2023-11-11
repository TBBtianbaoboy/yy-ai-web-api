//---------------------------------
//File Name    : grpc.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-30 14:23:57
//Description  :
//----------------------------------
package rpc

import (
	"nas-common/common"
	"nas-common/mlog"
	"nas-common/rpcapi/forward"
	"nas-web/config"

	"go.uber.org/zap"
)

var RpcForwardClient forward.ForwardClient

func InitRpcForwardClient(addr string) error {
	zClient, err := NewClient(RpcClientConf{
		Target: addr,
		App:    DefaultApp,
		Token:  DefaultAppToken,
		EtcdKey: common.EtcdKeyForward,
		EtcdHosts: config.IrisConfig.Etcd.EtcdHost,
	})
	if err != nil {
		mlog.Error("Connect Forward Rpc Server failed", zap.Error(err))
		return err
	}
	RpcForwardClient = forward.NewForwardClient(zClient.Conn())
	return nil
}
