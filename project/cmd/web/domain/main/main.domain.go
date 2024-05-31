package mainDomain

import (
	"net/http"
	"GoBaby/internals/utils"
	"GoBaby/internals/models"
	"GoBaby/ui"
)


func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	files := []string{
		"html/base.html",
		"html/pages/main/main.tmpl.html",
	}

	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, files...)
} 