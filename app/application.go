package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication(){
	UrlMapping()
	router.Run(":8080")
}
