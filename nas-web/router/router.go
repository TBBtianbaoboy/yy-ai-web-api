package router

import (
	"nas-web/middleware"
	"nas-web/router/api/common"
	v1 "nas-web/router/api/v1"

	"github.com/kataras/iris/v12"
)

func InitRouters(app *iris.Application) {
	app.Use(middleware.JwtMiddleware) //jwt check middleware

	commonRouter := app.Party("/")
	{
		common.RegisterAuthRouter(commonRouter)
	}
	appRouter := app.Party("/v1/")
	{
		//User Router
		appUserRouter := appRouter.Party("/user")
		{
			v1.RegisterUserRouter(appUserRouter)
		}
		//Chat Router
		appChatRouter := appRouter.Party("/chat")
		{
			v1.RegisterChatRouter(appChatRouter)
		}
		//Image Router
		appImageRouter := appRouter.Party("/image")
		{
			v1.RegisterImageRouter(appImageRouter)
		}
	}
}
