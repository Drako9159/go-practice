package repository_adapters

import (
	"GoBaby/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	db_config "GoBaby/cmd/web/domain/repository/config"
	
)

func GetUserByUUID(uuid string) (*models.User, *models.AppError) {
	filter := bson.D{
		{
			Key: "$and", Value: bson.A{
				bson.D{
					{Key: "_id", Value: uuid}
				}.
			},
		},
	}
	var result models.User
	var error *models.AppError
	
	err := db_config.UserCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		error = &models.AppError{
			Message: "Error querying user",
			Err:     err,
			Code:    500,
		}
		// return nil, error
	}

	return result, error
}