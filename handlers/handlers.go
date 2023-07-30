package handlers

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

func SetUpS3Uploader() *manager.Uploader {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	return uploader
}

func SaveToS3Bucket(c *gin.Context) {
	uploader := SetUpS3Uploader()

	file, err := c.FormFile("upload")

	if err != nil {
		log.Fatal("Error in getting the file")
	}

	f, err := file.Open()

	if err != nil {
		log.Fatal("Error in Opening the file")

	}
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("golang-testing"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "File Uploaded Successfully",
		"Link":    result.Location,
	})

}
