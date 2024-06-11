package routes

import (
	"GoBaby/internals/models"
	logDomain "GoBaby/cmd/web/domain/log"
)	


func LogRender() {

	mux.HandleFunc("GET "+models.RoutesInstance.LOGS, logDomain.LogView)
	
}

