package v1

import (
	"nas-web/controller"
	"nas-web/interal/wrapper"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterUserRouter(party router.Party) {
	party.Handle("POST", "/", wrapper.Handler(controller.UserController{}.AddUser))
	party.Handle("DELETE", "/", wrapper.Handler(controller.UserController{}.DeleteUser))
	party.Handle("PUT", "/", wrapper.Handler(controller.UserController{}.EditUser))
	party.Handle("GET", "/", wrapper.Handler(controller.UserController{}.ListUser))
	party.Handle("GET", "/info", wrapper.Handler(controller.UserController{}.GetUserInfo))
	party.Handle("POST", "/reset_passwd", wrapper.Handler(controller.UserController{}.ResetPasswd))
	party.Handle("PUT", "/passwd/", wrapper.Handler(controller.UserController{}.UpdateUserPasswd))
	party.Handle("PUT", "/status/", wrapper.Handler(controller.UserController{}.UpdateUserStatus))
}
