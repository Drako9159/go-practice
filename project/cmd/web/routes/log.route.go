package routes

import (
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	logDomain "GoBaby/cmd/web/domain/log"
	"net/http"
	"GoBaby/ui"
)


func LogRender() {

	mux.HandlerFunc("GET "+models.RoutesInstance.LOGS, logDomain.LogView)
}

