package mainRoute

import (
	"GoBaby/cmd/web/routes"
	"GoBaby/cmd/web/domain/main"
	"GoBaby/internals/models"
)

func MainRender() {
	routes.GetMuxInstance().HandleFunc("GET "+models.RoutesInstance.MAIN, mainDomain.MainView)
}

