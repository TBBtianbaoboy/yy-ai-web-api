package main

import (
	nasweb "nas-web"
	"nas-web/config"
	_ "nas-web/docs" // this to solve swagger failed to load spec problem
	"nas-web/middleware/preset"
	"nas-web/router"
	"nas-web/support"

	"github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
	app := iris.New()
	// preset log and swagger
	preset.PreSetting(app)
	// init all routers
	router.InitRouters(app)
	return app
}

func main() {
	if result := nasweb.InitService(); !result {
		panic(support.InitServiceError)
	}
	app := newApp()
	err := app.Run(iris.Addr(config.IrisConfig.Web.Host), iris.WithRemoteAddrHeader("X-Forwarded-For"))
	if err != nil {
		panic(support.InitAppError)
	}
}
