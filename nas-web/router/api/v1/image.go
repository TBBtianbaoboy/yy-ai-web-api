package v1

import (
	"nas-web/controller"
	"nas-web/interal/wrapper"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterImageRouter(party router.Party) {
	party.Handle("POST", "/generate", wrapper.Handler(controller.ImageController{}.GenerateImage))
}
