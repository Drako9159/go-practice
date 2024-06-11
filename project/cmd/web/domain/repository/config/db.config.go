package db_config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"GoBaby/internals/models"
	"log/slog"
)

var dataBase = "GoBaby"
var collections = map[string]string{
	"users": "users",
	"logs": "logs",
}



var UserCollection *mongo.Collection

func InitializeUsersCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(dataBase).Collection(collections["users"])
}


func InitializeDb() (*mongo.Client, *models.AppError) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	
	if err != nil {
		return nil, &models.AppError{
			Message: "Error connecting to database",
			Err:     err,
			Code:    500,
		}
	}

	UsersCollection := InitializeUsersCollection(client)

	result := UsersCollection.FindOne(context.TODO(), bson.M{"_id": 0})
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			user := models.User{
				UserName: "test",
				Logs:     make([]models.Log, 0),
				Id:       0,
			}

			_, error := UsersCollection.InsertOne(context.TODO(), user)

			if error != nil {
				return nil, &models.AppError{
					Message: "Error inserting user",
					Err:     err,
					Code:    500,
				}
			}
		} else {
			return nil, &models.AppError{
				Message: "Error querying user",
				Err:     err,
				Code:    500,
			}
		}
	}  else {
		slog.Info("user with _id 0 exists")
	}

	slog.Info("Database initialized")
	return client, nil	
}


