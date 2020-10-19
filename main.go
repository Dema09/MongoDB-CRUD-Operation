package main

import (
	"MongoDB-CRUD-Operation/app"
	"MongoDB-CRUD-Operation/config/database"
)

func main(){
	database.ConnectDB()
	app.StartApplication()
}
