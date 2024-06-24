package repository_domain

import (
	"GoBaby/internals/models"
	db_config "GoBaby/cmd/web/domain/repository/config"
	repository_adapters "GoBaby/cmd/web/domain/repository/adapters"
	"log/slog"
	
)

func InitializeDb() {
	err := db_config.InitializeDb()
	if err != nil {
		slog.Error(err.Message)
		panic(err.Error) // we want it to panic as it's a critical error
	}
}

func GetUserByUUID(uuid int) (models.User, *models.AppError) {
	return repository_adapters.GetUserByUUID(uuid)
}

func SetUser(user *models.User) *models.AppError {
	return repository_adapters.SetUser(user)
}

func AddLogByUUID(uuid int, log models.Log) *models.AppError {
	return repository_adapters.AddLogByUUID(uuid, log)
}