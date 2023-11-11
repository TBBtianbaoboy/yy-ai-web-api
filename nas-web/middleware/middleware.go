//---------------------------------
//File Name    : middleware.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-18 14:41:57
//Description  :
//----------------------------------
package middleware

import (
	"nas-web/config"
	"nas-web/middleware/jwt"
	"nas-web/middleware/basic"
	"nas-web/support"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func JwtMiddleware(ctx context.Context) {
	if !basic.CheckURL(ctx.Path(), config.IrisConfig.Other.IgnoreUrls) {
		if ctx.Values().Get(jwt.DefaultContextKey) == nil {
			//验证模式 cookie
			authCookie := ctx.GetCookie(support.Auth)
			if authCookie != "" && ctx.GetHeader("authorization") == "" {
				ctx.Request().Header.Set("authorization", authCookie)
			}
			//下载的 authorization 处理
			authorization := ctx.FormValue("authorization")
			if authorization != "" && ctx.GetHeader("authorization") == "" {
				ctx.Request().Header.Set("authorization", authorization)
				delete(ctx.FormValues(), authorization)
			}
			//jwt token拦截
			if !jwt.Server(ctx) {
				ctx.StatusCode(iris.StatusUnauthorized)
				ctx.StopExecution()
			}
		}
	}
	ctx.Next()
}
