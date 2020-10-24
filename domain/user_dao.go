package domain

import (
	"MongoDB-CRUD-Operation/config/database"
	"MongoDB-CRUD-Operation/config/logger"
	"MongoDB-CRUD-Operation/response"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var(
	userCollection = database.ConnectDB().Database("user_database").Collection("users")
)

func (user *User) SaveUser() (interface{}, *response.RestBody){
	insertResult, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil{
		logger.Error("There is something error when inserting data into MongoDB Database: ", err)
		return nil,response.NewInternalServerError("There is something error in our database!")
	}
	return insertResult.InsertedID, nil
}

func (user *User) GetAllUser() ([]bson.M,*response.RestBody){
	//primitive.M is an unordered representation of a BSON (Since MongoDB saving a data as BSON)
	userResult, err := userCollection.Find(context.TODO(), bson.M{})

	if err != nil{
		logger.Error("There is something error when getting user data from database: ", err)
		return nil, response.NewInternalServerError("There is an error in our database!")
	}
	var users []bson.M

	if err = userResult.All(context.TODO(), &users); err != nil{
		logger.Error("Cannot fetch user data from database:", err)
		return nil, response.NewInternalServerError("There is an error in our database!")
	}
	return users, nil
}

func (user *User) UpdateUserById(userId string)(string, *response.RestBody){
	id, _ := primitive.ObjectIDFromHex(userId)
	result, err := userCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.D{
		{"$set", bson.D{
			{"firstname", user.FirstName},
			{"lastname", user.LastName},
			{"address", user.Address},
		}},
	})

	if err != nil{
		logger.Error("Error when trying to update the user data: ", err)
		return "", response.NewInternalServerError("There is something error in our database!")
	}
	return fmt.Sprintf("Updated %d Data!", result.ModifiedCount), nil
}
