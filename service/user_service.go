package service

import (
	"MongoDB-CRUD-Operation/domain"
	"MongoDB-CRUD-Operation/response"
	"go.mongodb.org/mongo-driver/bson"
)

func SaveUserData(user domain.User) (interface{},*response.RestBody){
	if err := user.ValidateUser(); err != nil{
		return nil, err
	}
	insertResult, err := user.SaveUser()

	if err != nil{
		return nil, err
	}
	return insertResult, nil
}

func GetAllUserData()([]bson.M, *response.RestBody){
	userData := &domain.User{}
	userDataResponse, err := userData.GetAllUser()
	if err != nil{
		return nil, err
	}
	return userDataResponse, nil
}
