package routers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bhushan-mdn/the-archives/models"
	"github.com/bhushan-mdn/the-archives/pkg/storage"
	"github.com/gin-gonic/gin"
)

type AddFileInput struct {
	Name     string `json:"name" binding:"required"`
	Location string `json:"location" binding:"required"`
	Size     string `json:"size" binding:"required"`
}

// Create File creates a single file
func CreateFile(c *gin.Context) {
	file := AddFileInput{}
	err := c.BindJSON(&file)
	if err != nil {
		log.Fatal("Error: ", err)
		c.JSON(400, "Wrong format")
	}

	data := map[string]interface{}{
		"name":     file.Name,
		"location": file.Location,
		"size":     file.Size,
	}

	err = models.CreateFile(data)
	if err != nil {
		log.Fatal("Error: ", err)
		c.JSON(500, "File could not be created")
	}
	c.JSON(201, "File created successfully")
}

// GetAllFiles returns all the stored files
func GetAllFiles(c *gin.Context) {
	files, err := models.GetFiles()
	log.Println(files)
	if err != nil {
		log.Fatal(err)
	}
	resp := map[string]interface{}{"files": files}
	c.JSON(200, resp)
}

func GetFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id should be int"})
	}
	file, err := models.GetFile(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": file})
}

type UpdateFileInput struct {
	Name string `json:"name" binding:"required"`
}

func UpdateFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id should be int"})
	}
	file, err := models.GetFile(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	input := UpdateFileInput{}
	err = c.Bind(&input)
	if err != nil {
		log.Fatal("Error: ", err)
		c.JSON(400, "Wrong format")
	}

	file.Name = input.Name

	err = models.EditFile(id, file)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"message": "Error updating file in DB"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated file"})
}

func DeleteFile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "id must be int"})
		return
	}
	log.Println(id)

	exists, err := models.ExistFileByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "Error getting file"})
		return
	}
	if !exists {
		c.JSON(404, gin.H{"message": "File does not exist"})
		return
	}

	file, err := models.GetFile(id)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"message": "Error getting file"})
		return
	}

	err = storage.RemoveObject(file.Location)
	if err != nil {
		log.Println(err)
		c.JSON(404, gin.H{"message": "Error deleting file from storage"})
		return
	}

	err = models.DeleteFile(id)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"message": "File deleted successfully"})
}
