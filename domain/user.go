package domain

import (
	"MongoDB-CRUD-Operation/response"
)

type User struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string  `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	ProfilePicture string `json:"profile_picture"`
}

func (user *User) ValidateUser() *response.RestBody{
	if user.FirstName == "" || user.LastName == ""{
		return response.NewBadRequest("Your First Name or Last Name is Empty!")
	}
	return nil
}
