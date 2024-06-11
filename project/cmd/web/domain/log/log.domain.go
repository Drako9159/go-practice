package logDomain

import (
	"GoBaby/internals/models"
	"GoBaby/internals/utils"
	"net/http"
	"GoBaby/ui"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	repository_domain "GoBaby/cmd/web/domain/repository"
	"log/slog"
)

type LogViewModel struct {
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

	users, err := repository_domain.GetUserByUUID(0) 
	if err != nil {
		slog.Log(context.TODO(), slog.LevelError, err.Message)
		return
	}

	utils.ParseTemplateFiles(w, "logTable", users.Logs, utils.EmptyFuncMap, ui.Content, "html/pages/logs/logTable.tmpl.html")
}


func SaveLog(countdown int) *models.AppError {
	uuid := 0
	currentTime := time.Now()
	primitiveDate := primitive.NewDateTimeFromTime(currentTime)
	err := repository_domain.AddLogByUUID(uuid, models.Log{Date: primitiveDate, Duration: countdown})
		
	
	return err
}
