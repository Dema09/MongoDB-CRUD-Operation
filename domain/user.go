package domain

import (
	"MongoDB-CRUD-Operation/response"
)

type User struct{
	UserId string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string  `json:"address"`
}

func (user *User) ValidateUser() *response.RestBody{
	if user.FirstName == "" || user.LastName == ""{
		return response.NewBadRequest("Your First Name or Last Name is Empty!")
	}
	return nil
}
