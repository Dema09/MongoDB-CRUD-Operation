package app

import "MongoDB-CRUD-Operation/controller"

func UrlMapping(){
	router.POST("/v1/createUser", controller.CreateUser)
}
