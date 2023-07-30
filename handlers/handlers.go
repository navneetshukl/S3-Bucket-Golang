package handlers

import "github.com/gin-gonic/gin"

func Home(c *gin.Context) {
	c.HTML(200, "home.html", gin.H{})
}

func SaveFileToDisk(c *gin.Context) {
	file, err := c.FormFile("upload")

	if err != nil {
		return
	}
	c.SaveUploadedFile(file, "uploads/"+file.Filename)
	c.HTML(200, "home.html", gin.H{})
}
