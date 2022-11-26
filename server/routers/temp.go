package routers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func upload(c *gin.Context) {
	// Single file
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	log.Println(file.Filename)

	// Upload the file to specific dst.
	filename := filepath.Base(file.Filename)
	// dst := fmt.Sprintf("/tmp/filename/%s", filename)
	// if err := c.SaveUploadedFile(file, dst); err != nil {
	// 	c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
	// 	return
	// }
	// c.File

	log.Println(filename)
	log.Println(file.Size)
	fmt.Printf("%v %T", file.Header["Content-Type"], file.Header["Content-Type"])
	fileContent, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer fileContent.Close()
	var buf bytes.Buffer
	io.Copy(&buf, fileContent)
	if err != nil {
		log.Fatal(err)
	}
	content := buf.String()
	log.Println(content)
	buf.Reset()

	// if err := saveFile(); err != nil {
	// 	c.String(http.StatusInternalServerError, "save file err: %s", err.Error())
	// }

	// print(saveFile())

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
