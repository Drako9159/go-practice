package logDomain

import (
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	"net/http"
)

type LogViewModel static {
	Logs []models.Log
}

func LogView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.LOGS)
	
	files := []string{
		"html/base.html",
		"html/pages/log/log.tmpl.html",
	}

	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, files...)
}