package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
	"vespera-server/bucket"
	"vespera-server/database"
	"vespera-server/services"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

// load env

func UploadUserImage(c *gin.Context) {
	userID := c.Param("userID")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image file"})
		return
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	if bucket.S3Bucket == nil {
		log.Printf("Bucket not initialized")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bucket not initialized"})
		return
	}
	_, err = bucket.S3Bucket.Svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket.S3Bucket.Name),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(header.Header.Get("Content-Type")),
	})
	if err != nil {
		log.Printf("Failed to upload file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	imageURL := fmt.Sprintf("http://localhost:9000/%s/%s", bucket.S3Bucket.Name, fileName)

	err = services.AddUserImage(database.DB, userID, imageURL)

	if err != nil {

		log.Printf("Failed to add image to database: %v\n", err)
		// Delete the image from the bucket
		_, err = bucket.S3Bucket.Svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket.S3Bucket.Name), Key: aws.String(fileName)})

		if err != nil {
			log.Printf("Failed to delete image from bucket: %v\n", err)
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add image to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "url": imageURL})
}

func UploadEventoImage(c *gin.Context) {
	eventoID := c.Param("eventoID")

	if eventoID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read image file"})
		return
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	if bucket.S3Bucket == nil {
		log.Printf("Bucket not initialized")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bucket not initialized"})
		return
	}
	_, err = bucket.S3Bucket.Svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket.S3Bucket.Name),
		Key:         aws.String(fileName),
		Body:        bytes.NewReader(buf.Bytes()),
		ContentType: aws.String(header.Header.Get("Content-Type")),
	})
	if err != nil {
		log.Printf("Failed to upload file: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	imageURL := fmt.Sprintf("http://localhost:9000/%s/%s", bucket.S3Bucket.Name, fileName)

	err = services.AddEventoImage(database.DB, eventoID, imageURL)

	if err != nil {

		log.Printf("Failed to add image to database: %v\n", err)
		// Delete the image from the bucket
		_, err = bucket.S3Bucket.Svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket.S3Bucket.Name), Key: aws.String(fileName)})

		if err != nil {
			log.Printf("Failed to delete image from bucket: %v\n", err)
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add image to database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "url": imageURL})
}
