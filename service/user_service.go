package service

import (
	"MongoDB-CRUD-Operation/config/logger"
	"MongoDB-CRUD-Operation/constant"
	"MongoDB-CRUD-Operation/domain"
	"MongoDB-CRUD-Operation/response"
	"bufio"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	_ "image/jpeg"
	"io"
	"io/ioutil"
	"math"
	"os"
	"time"
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

	userData := &domain.User{}
	userData.FirstName = c.PostForm("first_name")
	userData.LastName = c.PostForm("last_name")
	userData.Address = c.PostForm("address")
	userData.ProfilePicture = convertProfilePicture(c)
	
	editUserDataResponse, editError := userData.EditUserProfile(userId)
	if editError != nil{
		return nil, editError
	}
	return response.NewStatusOK(editUserDataResponse), nil
}

func ShowUserProfileByUserId(userId string)(*response.RestBody, *response.RestBody){

	userData := &domain.User{}
	profileMapping := &domain.ProfileResponse{}

	findUserResponse, err := userData.FindUserByUserId(userId)
	if err != nil {
		return nil, err
	}

	imageFile, imageErr := os.Open(constant.PhotoProfilePath + findUserResponse.ProfilePicture)
	if imageErr != nil{
		logger.Error("Error when opening the picture", imageErr)
		return nil, response.NewInternalServerError("Cannot find your image!")
	}

	reader := bufio.NewReader(imageFile)
	content, _ := ioutil.ReadAll(reader)
	encodedImageInString := base64.StdEncoding.EncodeToString(content)

	profileMapping.FirstName = findUserResponse.FirstName
	profileMapping.LastName = findUserResponse.LastName
	profileMapping.Address = findUserResponse.Address
	profileMapping.ProfilePicture = encodedImageInString
	profileMapping.Age = getUserAge(findUserResponse.DateOfBirth)

	logger.Info("Getting user profile successfully!")
	return response.NewStatusOK(profileMapping), nil

}

func getUserAge(dateOfBirth string) float64 {
	userAge, err := time.Parse(constant.DateFormat, dateOfBirth)
	if err != nil{
		logger.Error("Error when parsing user date of birth", err)
		return 0
	}
	birthDay := time.Date(userAge.Year(), userAge.Month(), userAge.Day(), 0, 0, 0, 0, time.UTC)
	today := time.Now()

	age := math.Floor(today.Sub(birthDay).Hours() / 24 / 365)
	return age
}

func GetAllAdultUser() (*response.RestBody, *response.RestBody){
	var userData domain.User
	adultUserResponse, err := userData.FindAllAdultUser()
	if err != nil{
		return nil, err
	}
	return response.NewStatusOK(adultUserResponse), nil
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

	logger.Info("Successfully convert and save image to file system!")
	return handler.Filename
}
