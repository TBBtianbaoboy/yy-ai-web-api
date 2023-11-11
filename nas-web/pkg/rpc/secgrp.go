// ---------------------------------
// File Name    : pkg/rpc/secgrp.go
// Author       : aico
// Mail         : 2237616014@qq.com
// Github       : https://github.com/TBBtianbaoboy
// Site         : https://www.lengyangyu520.cn
// Create Time  : 2021-12-30 15:43:31
// Description  :
// ----------------------------------
package rpc

import (
	"context"
	"nas-common/rpcapi/forward"
)

func SendSecGrpRule(ctx context.Context, req *forward.ForwardAgentActionReq) (resp *forward.ForwardAgentActionResp, err error) {
	resp, err = RpcForwardClient.ForwardAgentAction(ctx, req)
	return
}
