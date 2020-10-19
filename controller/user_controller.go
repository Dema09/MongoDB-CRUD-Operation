package controller

import (
	"MongoDB-CRUD-Operation/domain"
	"MongoDB-CRUD-Operation/response"
	"MongoDB-CRUD-Operation/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context){
	var user domain.User
	if  err := c.ShouldBindJSON(&user); err != nil{
		restError := response.NewBadRequest("Invalid JSON Body!")
		c.JSON(restError.StatusCode, restError)
		return
	}

	userResponse, err := service.SaveUserData(user)
	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, response.NewStatusCreated("The user data has been inserted to database!", userResponse))
}
