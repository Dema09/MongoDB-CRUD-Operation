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

func GetAllUser(c *gin.Context){
	userResponse, err := service.GetAllUserData()
	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, response.NewStatusOK(userResponse))
}

func EditUserData(c *gin.Context){
	var updateUserBody domain.User
	userId := c.Query("user_id")

	if err := c.ShouldBindJSON(&updateUserBody); err != nil{
		responseError := response.NewBadRequest("Invalid JSON Body!")
		c.JSON(responseError.StatusCode, responseError)
		return
	}
	editProfileResponse, err := service.EditUserDataById(userId, updateUserBody)

	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, response.NewStatusOK(editProfileResponse))
}

func DeleteUserData(c *gin.Context){
	userId := c.Query("user_id")

	deleteUserDataById, err := service.DeleteUserDataById(userId)
	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, deleteUserDataById)
}

func GetUserById(c *gin.Context){
	userId := c.Query("user_id")
	userResponse, err := service.GetUserDataById(userId)

	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, userResponse)
}

func EditProfile(c *gin.Context){
	userId := c.Query("user_id")
	editProfileResponse, err := service.EditCurrentProfile(c, userId)

	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(editProfileResponse.StatusCode, editProfileResponse)
}

func ShowUserProfileByUserId(c *gin.Context){
	userId := c.Query("user_id")
	showImageResponse, err := service.ShowUserProfileByUserId(userId)
	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, showImageResponse)
}

func GetAdultUser(c *gin.Context){
	getAdultUserResponse, err := service.GetAllAdultUser()
	if err != nil{
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(getAdultUserResponse.StatusCode, getAdultUserResponse)
}
