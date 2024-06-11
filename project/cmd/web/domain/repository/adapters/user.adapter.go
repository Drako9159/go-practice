package repository_adapters

import (
	"GoBaby/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	db_config "GoBaby/cmd/web/domain/repository/config"
	"context"
)

func GetUserByUUID(uuid int) (models.User, *models.AppError) {
	filter := bson.D{
		{
			Key: "$and", Value: bson.A{
				bson.D{
					{Key: "_id", Value: uuid},
				},
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
	}

	return result, error
}


func SetUser(user *models.User) *models.AppError {
	_, err := db_config.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return &models.AppError{
			Message: "Error inserting user",
			Err:     err,
			Code:    500,
		}
	}
	return nil
}


func AddLogByUUID(uuid int, log models.Log) *models.AppError {
	//_, error := db_config.UserCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: uuid}}, bson.D{{Key: "$push", Value: bson.D{{Key: "logs", Value: log}}}})

	filter := bson.M{"_id": uuid}

	update := bson.M{
		"$push": bson.M{
			"logs": log,
		},
	}

	_, err := db_config.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return &models.AppError{
			Message: "Error inserting log",
			Err:     err,
			Code:    500,
		}
	}

	return nil
}