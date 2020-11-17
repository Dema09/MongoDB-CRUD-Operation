package service

import (
	"MongoDB-CRUD-Operation/config/logger"
	"MongoDB-CRUD-Operation/constant"
	"MongoDB-CRUD-Operation/domain"
	"MongoDB-CRUD-Operation/response"
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "image/jpeg"
	"io"
	"io/ioutil"
	"os"
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

func EditUserDataById(userId string, updateUserData domain.User)(interface{}, *response.RestBody){
	updateUserResponse, err := updateUserData.UpdateUserById(userId)

	if err != nil{
		return nil, err
	}
	return updateUserResponse, nil
}

func DeleteUserDataById(userId string)(string, *response.RestBody){
	userData := &domain.User{}
	deleteUserDataResponse, err := userData.DeleteUserByUserId(userId)

	if err != nil{
		return "", err
	}
	return deleteUserDataResponse, nil
}

func GetUserDataById(userId string) (interface{}, *response.RestBody){
	userData := &domain.User{}
	userByIdResponse, err := userData.FindUserDataById(userId)
	if err != nil{
		return nil, err
	}
	return userByIdResponse, nil
}

func EditCurrentProfile(c *gin.Context, userId string) (*response.RestBody, *response.RestBody){
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil{
		logger.Error("Cannot Parse Id!", err)
		return nil,  response.NewBadRequest(fmt.Sprintf("Cannot Parse id %s", userId))
	}
	
	userData := &domain.User{UserId: id}
	userData.FirstName = c.PostForm("first_name")
	userData.LastName = c.PostForm("last_name")
	userData.Address = c.PostForm("address")
	userData.ProfilePicture = convertProfilePicture(c)
	
	editUserDataResponse, editError := userData.EditUserProfile()
	if editError != nil{
		return nil, editError
	}
	return response.NewStatusOK(editUserDataResponse), nil
}

func ShowUserProfileByUserId(userId string)(*response.RestBody, *response.RestBody){
	userIdFromObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil{
		logger.Error(fmt.Sprintf("Cannot Parse Id with Id: %s", userId), err)
		return nil, response.NewBadRequest("Your Id is Invalid!")
	}

	userData := &domain.User{UserId: userIdFromObjectId}
	findUserResponse := userData.FindUserByUserId()

	imageFile, imageErr := os.Open(constant.PhotoProfilePath + findUserResponse.ProfilePicture)
	if imageErr != nil{
		logger.Error("Error when opening the picture", imageErr)
		return nil, response.NewInternalServerError("Cannot find your image!")
	}

	reader := bufio.NewReader(imageFile)
	content, _ := ioutil.ReadAll(reader)
	encodedImageInString := base64.StdEncoding.EncodeToString(content)

	userData.UserId = findUserResponse.UserId
	userData.FirstName = findUserResponse.FirstName
	userData.LastName = findUserResponse.LastName
	userData.Address = findUserResponse.Address
	userData.ProfilePicture = encodedImageInString

	return response.NewStatusOK(userData), nil

}

func convertProfilePicture(c *gin.Context) string {
	c.Request.ParseMultipartForm(10 << 20)
	file, handler, err := c.Request.FormFile("profile_picture")

	if err != nil{
		logger.Error("Cannot Read Picture", err)
		return ""
	}

	defer file.Close()

	filePath, fileErr := os.OpenFile(constant.PhotoProfilePath + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if fileErr != nil{
		logger.Error("Cannot Open File Path", err)
		return ""
	}

	defer filePath.Close()
	io.Copy(filePath, file)
	return handler.Filename
}
