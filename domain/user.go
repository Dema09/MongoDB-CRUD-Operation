package domain

import (
	"MongoDB-CRUD-Operation/response"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	UserId primitive.ObjectID `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string  `json:"address"`
	ProfilePicture string `json:"profile_picture"`
}

func (user *User) ValidateUser() *response.RestBody{
	if user.FirstName == "" || user.LastName == ""{
		return response.NewBadRequest("Your First Name or Last Name is Empty!")
	}
	return nil
}
