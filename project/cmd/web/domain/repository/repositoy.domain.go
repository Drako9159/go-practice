package repository_domain

import (
	"GoBaby/internals/models"
	db_config "GoBaby/cmd/web/domain/repository/config"
	"go.mongodb.org/mongo-driver/mongo"
	repository_adapters "GoBaby/cmd/web/domain/repository/adapters"
)

func InitializeDb() (*mongo.Client, *models.AppError) {
	return db_config.InitializeDb()
}

func GerUserByUUID(uuid int) (*models.User, *models.AppError) {
	return repository_adapters.GetUserByUUID(uuid)
}

func SetUser(user models.User) *models.AppError {
	return repository_adapters.SetUser(user)
}

func AddLogByUUID(uuid int, log models.Log) *models.AppError {
	return repository_adapters.AddLogByUUID(uuid, log)
}