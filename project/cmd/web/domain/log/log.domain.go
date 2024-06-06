package logDomain

import (
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	"net/http"
	"GoBaby/ui"
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

func LogTable(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.LOG_TABLE)
	utils.ParseTemplateFiles(w, "logTable", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, "html/pages/logs/logTable.tmpl.html")
}