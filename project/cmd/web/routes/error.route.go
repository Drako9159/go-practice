package routes

import (
	"GoBaby/internals/models"
	errorDomain "GoBaby/cmd/web/domain/error"
)

func ErrorRender() {
	mux.HandleFunc("GET "+models.RoutesInstance.CLEAR_ERROR, errorDomain.ClearErrorTemplate)
}

