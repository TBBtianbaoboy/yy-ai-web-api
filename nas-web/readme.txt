// generate debug page in web
go install github.com/swaggo/swag/cmd/swag@latest
swag init -g middleware/preset/preset.go

//iris version
github.com/kataras/iris/v12 v12.1.8

//captcha version
github.com/mojocn/base64Captcha v0.0.0-20190509095025-87c9c59224d8

// 清除 mod 缓存
go clean -modcache


