package app

import "MongoDB-CRUD-Operation/controller"

func UrlMapping(){
	router.POST("/v1/createUser", controller.CreateUser)
	router.GET("/v1/getAllUser", controller.GetAllUser)
	router.PUT("/v1/updateUserData", controller.EditUserData)
	router.DELETE("/v1/deleteUserData", controller.DeleteUserData)
	router.GET("/v1/getUserById", controller.GetUserById)
	router.PUT("/v1/editProfile", controller.EditProfile)
}
