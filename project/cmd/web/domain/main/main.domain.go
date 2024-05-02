package mainDomain

import (
	"net/http"
	"fmt"
	"GoBaby/internals/utils"
	"GoBaby/internals/models"
	
)


func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	for := []string{
		"ui/html/base.html",
		"ui/html/pages/main/main.tmpl.html",
	}

	// 54
}