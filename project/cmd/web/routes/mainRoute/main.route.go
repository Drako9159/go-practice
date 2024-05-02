package mainRoute

import (
	"GoBaby/cmd/web/routes"
	"net/http"
	"GoBaby/internals/models"
)

func MainRender() {
	routes.GetMuxInstance().HandleFunc("GET "+models.RoutesInstance.MAIN, func(w http.ResponseWriter, r *http.Request) {})
}

