package repository_adapters

import (
	"GoBaby/internals/models"
	db_config "GoBaby/cmd/web/domain/repository/config"
	"context"
)

func AddMonitorLog(monitorLog models.MonitorLog) *models.AppError {
	
	_, err := db_config.MonitorCollection.InsertOne(context.TODO(), monitorLog)
	if err != nil {
		return &models.AppError{
			Message: "Error inserting log",
			Err:     err,
			Code:    500,
		}
	}
   
	return nil
}