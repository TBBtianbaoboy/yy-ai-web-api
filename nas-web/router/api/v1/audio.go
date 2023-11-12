package v1

import (
	"nas-web/controller"
	"nas-web/interal/wrapper"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterAudioRouter(party router.Party) {
	party.Handle("POST", "/transcriptions", wrapper.Handler(controller.AudioController{}.Transcriptions))
}
