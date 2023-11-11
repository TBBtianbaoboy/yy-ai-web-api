package v1

import (
	"nas-web/controller"
	"nas-web/interal/wrapper"

	"github.com/kataras/iris/v12/core/router"
)

func RegisterImageRouter(party router.Party) {
	party.Handle("GET", "/", wrapper.Handler(controller.AgentController{}.ListAgent))
	party.Handle("GET", "/info/system", wrapper.Handler(controller.AgentController{}.AgentSystemInfo))
	party.Handle("GET", "/info/port", wrapper.Handler(controller.AgentController{}.AgentPortInfo))
	party.Handle("POST", "/download", wrapper.Handler(controller.AgentController{}.DownloadAgent))
	party.Handle("DELETE", "/", wrapper.Handler(controller.AgentController{}.DeleteAgent))
	party.Handle("POST", "/secgrp", wrapper.Handler(controller.AgentController{}.AddAgentSecGrpRule))
	party.Handle("DELETE", "/secgrp", wrapper.Handler(controller.AgentController{}.DeleteAgentSecGrpRule))
	party.Handle("GET", "/secgrp", wrapper.Handler(controller.AgentController{}.ListAgentSecGrp))
	party.Handle("POST", "/baseline", wrapper.Handler(controller.AgentController{}.StartBaselineScan))
	party.Handle("PUT", "/baseline", wrapper.Handler(controller.AgentController{}.UpdateAgentBaseline))
	party.Handle("GET", "/baseline", wrapper.Handler(controller.AgentController{}.ListAgentBaseline))
	party.Handle("GET", "/baseline/info", wrapper.Handler(controller.AgentController{}.GetBaselineInfo))
}
