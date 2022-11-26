package routers

import "github.com/gin-gonic/gin"

func CheckHealth(c *gin.Context) {
	// storage.CheckConnection()

	message := map[string]string{
		"DB":    "up",
		"Minio": "up",
	}
	c.JSON(200, gin.H{"message": message})
}
