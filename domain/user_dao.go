package domain

import (
	"MongoDB-CRUD-Operation/config/database"
	"MongoDB-CRUD-Operation/config/logger"
	"MongoDB-CRUD-Operation/response"
	"context"
)

var(
	userCollection = database.ConnectDB().Database("user_database").Collection("users")
)

func (user *User) SaveUser() *response.RestBody{
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil{
		logger.Error("There is something error when inserting data into MongoDB Database: ", err)
		return response.NewInternalServerError("There is something error in our database!")
	}
	return response.NewStatusOK(insertResult.InsertedID)
}
