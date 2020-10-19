package service

import (
	"MongoDB-CRUD-Operation/domain"
	"MongoDB-CRUD-Operation/response"
)

func SaveUserData(user domain.User) (*domain.User,*response.RestBody){
	if err := user.ValidateUser(); err != nil{
		return nil, err
	}
	 if err := user.SaveUser(); err != nil{
		return nil, err
	}

	return &user, nil
}
