// ---------------------------------
// File Name    : resp_api.go
// Author       : aico
// Mail         : 2237616014@qq.com
// Github       : https://github.com/TBBtianbaoboy
// Site         : https://www.lengyangyu520.cn
// Create Time  : 2021-12-14 17:43:04
// Description  :
// ----------------------------------
package support

import (
	"nas-common/mlog"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"go.uber.org/zap"
)

// 创建http响应
func makeResponse(ctx context.Context, data interface{}, msg string, errMsg string, code int) {
	resp := bson.M{
		"code":    code,
		"error":   errMsg,
		"data":    data,
		"message": msg,
	}
	_, err := ctx.JSON(resp)
	if err != nil {
		mlog.Info("make response error ", zap.Error(err))
	}
}

// 构建正确的接口返回
func SendApiResponse(ctx context.Context, data interface{}, msg string) {
	if msg == "" {
		msg = "success"
	}
	makeResponse(ctx, data, msg, "", iris.StatusOK)
}

// 构建错误的接口返回
func SendApiErrorResponse(ctx context.Context, msg string, statusCode int) {
	makeResponse(ctx, nil, msg, msg, statusCode)
}

// 设置登录cookie
func SetAuthCookie(ctx context.Context, token string) {
	if token != "" {
		ctx.SetCookie(&http.Cookie{
			Name:     Auth,
			Value:    token,
			Path:     "/",
			HttpOnly: true,
		})
	}
}
