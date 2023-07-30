package main

import (
	"S3-Bucket/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.GET("/",handlers.Home)
	r.POST("/",handlers.SaveToS3Bucket)

	r.Run()

}
