package routes

import (
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	logDomain "GoBaby/cmd/web/domain/log"
	"net/http"
	"GoBaby/ui"
)


func LogRender() {

	mux.HandlerFunc("GET "+models.RoutesInstance.LOGS, logDomain.LogView)
}

func LogTable(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.LOG_TABLE)
	utils.ParseTemplateFiles(w, "logTable", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, "html/pages/logs/logTable.tmpl.html")
}