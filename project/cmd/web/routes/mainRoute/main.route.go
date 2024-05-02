package mainRoute

import (
	"GoBaby/cmd/web/routes"
)

func MainRender() {
	routes.GetMuxInstance().HandleFunc("/getUsers", func(w http.ResponseWriter, r *http.Request) {})
}

