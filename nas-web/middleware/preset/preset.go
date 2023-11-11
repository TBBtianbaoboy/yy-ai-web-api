package preset

import (
	"fmt"
	"log"
	"nas-web/config"
	"net/http"
	"time"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	cover "github.com/kataras/iris/v12/middleware/recover"
	"gopkg.in/natefinch/lumberjack.v2"
)

func PreSetting(app *iris.Application) {
	//-----------------------------------------------------------preset log system
	rlog := log.New(&lumberjack.Logger{
		Filename:   config.IrisConfig.Log.LogPath,
		MaxSize:    config.IrisConfig.Log.MaxSize, // MB
		MaxBackups: 3,
		MaxAge:     config.IrisConfig.Log.MaxAge, // Days
		Compress:   true,                         // compress log
	}, "", 3)
	app.Logger().SetLevel(config.IrisConfig.Other.JwtLogLevel)
	c := logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query:              true,
		MessageContextKeys: []string{"message"},
		MessageHeaderKeys:  []string{"User-Agent"},
	}
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := fmt.Sprintf("%s | %v | %4v | %s | %s | %s | %v", now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, headerMessage)
		rlog.Println(output)
	}

	customLogger := logger.New(c)
	app.Use(
		customLogger,
		cover.New(),
	)
	//delete the last one / in url path,because there has error
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		path := r.URL.Path
		if len(path) > 1 && path[len(path)-1] == '/' && path[len(path)-2] != '/' {
			path = path[:len(path)-1]
			r.RequestURI = path
			r.URL.Path = path
		}
		next(w, r)
	})

	// -------------------------- register swagger for debuging
	conf := &swagger.Config{
		URL: "/swagger/doc.json",
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(conf, swaggerFiles.Handler))
}
