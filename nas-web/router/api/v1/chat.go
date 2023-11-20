package v1

import (
	"nas-web/controller"
	"nas-web/interal/wrapper"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterChatRouter(party router.Party) {
	party.Handle("POST", "/no_context_no_stream", wrapper.Handler(controller.ChatController{}.SendNoContextNoStreamChat))
	party.Handle("POST", "/no_context_stream", wrapper.Handler(controller.ChatController{}.SendNoContextStreamChat))
	party.Handle("POST", "/context_stream", wrapper.Handler(controller.ChatController{}.SendContextStreamChat))
	party.Handle("DELETE", "/delete_context_stream", wrapper.Handler(controller.ChatController{}.DeleteContextStreamChat))
	party.Handle("GET", "/get_all_sessions", wrapper.Handler(controller.ChatController{}.GetAllSessions))
	party.Handle("GET", "/get_session_messages", wrapper.Handler(controller.ChatController{}.GetSessionMessages))
}
