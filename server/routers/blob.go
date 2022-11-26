package routers

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/bhushan-mdn/the-archives/models"
	"github.com/bhushan-mdn/the-archives/pkg/storage"
	"github.com/gin-gonic/gin"
)

func CreateBlob(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(400, "Unable to find file in form")
		return
	}
	fileType := file.Header.Get("Content-Type")
	fileName := filepath.Base(file.Filename)
	fileSize := file.Size
	fileLocation := "files/" + fileName
	fmt.Println(fileName)
	fmt.Println(fileType)
	fmt.Println(fileSize, "Bytes")
	fmt.Println(fileLocation)

	// if fileSize > config.App.FileMaxSize {
	// 	log.Println("File Size > 15 MB")
	// 	c.JSON(413, "File Size > 15 MB")
	// 	return
	// }

	exists, err := models.ExistFileByName(fileName)
	log.Println(exists)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Error connecting to DB")
		return
	}

	if exists {
		log.Println("File already exists")
		c.JSON(400, "File already exists")
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fileContent)
	err = storage.PutObject(fileContent, fileLocation, fileType, fileSize)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Error uploading file")
		return
	}

	newFile := map[string]interface{}{
		"name":     fileName,
		"location": fileLocation,
		"type":     fileType,
		"size":     fileSize,
	}

	err = models.CreateFile(newFile)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Error uploading metadata to DB")
		return
	}

	c.JSON(201, "Successfully uploaded new file")
}

func GetBlob(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Invalid ID format")
		c.JSON(400, "Invalid ID format")
		return
	}

	file, err := models.GetFile(id)
	if err != nil {
		log.Println("File not found")
		c.JSON(404, "File not found")
		return
	}

	location := file.Location

	object, err := storage.GetObject(location)
	if err != nil {
		log.Println(err)
		c.JSON(500, "Error connecting to storage")
		return
	}

	stat, err := object.Stat()
	if err != nil {
		log.Println(err)
		c.JSON(500, "Error reading object")
		return
	}

	contentLength := stat.Size
	contentType := stat.ContentType

	c.DataFromReader(200, contentLength, contentType, object, nil)
}
