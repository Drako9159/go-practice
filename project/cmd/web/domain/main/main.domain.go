package mainDomain

import (
	"net/http"
	"fmt"
	"GoBaby/internals/utils"
	"GoBaby/internals/models"
	"html/template"
)


func MainView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.MAIN)

	files := []string{
		"ui/html/base.html",
		"ui/html/pages/main/main.tmpl.html",
	}

	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, files...)
} 