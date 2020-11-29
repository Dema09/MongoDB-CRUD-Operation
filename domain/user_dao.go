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

func (user *User) DeleteUserByUserId(userId string) (string, *response.RestBody){
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil{
		logger.Error(fmt.Sprintf("Cannot Read Id %s Because: ", userId), err)
		return "", response.NewInternalServerError("There something error in our database!")
	}
	deleteResponse, err := userCollection.DeleteOne(context.TODO(),
		bson.M{
		"_id": id,
		})

	if err != nil{
		logger.Error("Cannot Execute Delete Statement: ", err)
		return "", response.NewInternalServerError("There is something error in our database!")
	}

	return fmt.Sprintf("Successfully Delete %d Data With Id: %s", deleteResponse.DeletedCount, userId), nil

}

func (user *User) FindUserDataById(userId string) (*User, *response.RestBody){
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil{
		logger.Error(fmt.Sprintf("Cannot Read Id %s Because: ", userId), err)
		return nil, response.NewInternalServerError("There is something error in our system!")
	}

	userResponse := userCollection.FindOne(context.Background(),
		bson.M{
			"_id": id,
		})

	userDecode := &User{}
	userResponse.Decode(userDecode)
	return userDecode, nil
}

func (user *User) EditUserProfile(userId string) (string,*response.RestBody){
	userResponse, err := userCollection.UpdateOne(context.TODO(),
		bson.M{"_id" : userId},
		bson.D{
		{"$set", bson.D{
			{"firstname", user.FirstName},
			{"lastname", user.LastName},
			{"address", user.Address},
			{"profilepicture", user.ProfilePicture},
			},
		}})

	if err != nil{
		logger.Error("Cannot update the data", err)
		return "",response.NewInternalServerError("There is something error in our database!")
	}
	return fmt.Sprintf("Success Update %d Data!", userResponse.ModifiedCount), nil
}

func (user *User) FindUserByUserId(userId string) (*User, *response.RestBody){
	userIdInObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil{
		logger.Error(fmt.Sprintf("Error when getting an object_id from userid %s", userId), err)
		return nil, response.NewInternalServerError("There is something error in our database!")
	}

	userResponse := userCollection.FindOne(context.Background(),
		bson.M{
			"_id": userIdInObjectId,
		})
	userData := &User{}
	userResponse.Decode(userData)
	return userData, nil
}

func (user *User) FindAllAdultUser() (*User, *response.RestBody){
	userResponse, err := userCollection.Find(context.TODO(),
		bson.M{
			"age": bson.M{"$gt" : 17},
		})
	if err != nil{
		logger.Error("Error when getting adult user's data", err)
		return nil, response.NewInternalServerError("There is something error in our database!")
	}

	userData := &User{}
	userResponse.Decode(userData)
	return userData, nil

}

