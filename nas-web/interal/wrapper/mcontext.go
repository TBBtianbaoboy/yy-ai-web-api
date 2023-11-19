package wrapper

import (
	"nas-web/middleware/jwt"
	"sync"

	"nas-common/models"

	"github.com/kataras/iris/v12"
)

type Context struct {
	iris.Context
	UserToken *models.UserToken
}

var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

// 获取用户Token相关信息
func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original // set the context to the original one in order to have access to iris's implementation.
	ctx.UserToken = jwt.GetUserToken(original)
	return ctx
}

// 用户Token信息传递
func release(ctx *Context) {
	ctx.UserToken = nil
	contextPool.Put(ctx)
}

// 所有的操作都始于此，终于此
func Handler(handler func(*Context)) iris.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		handler(ctx)
		release(ctx)
	}
}
